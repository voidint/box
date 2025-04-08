// Copyright (c) 2025 voidint <voidint@126.com>. All rights reserved.
//
// This source code is licensed under the license found in the
// LICENSE file in the root directory of this source tree.
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package ecache

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/pkg/errors"
	"github.com/voidint/box/constraints"
	"golang.org/x/sync/singleflight"
)

var (
	singleFlightGroup singleflight.Group

	// Marshal specifies the default serialization function for cache values
	Marshal func(v any) ([]byte, error) = json.Marshal

	// Unmarshal specifies the default deserialization function for cached data
	Unmarshal func(data []byte, v any) error = json.Unmarshal

	// DefaultExpiration defines the default time-to-live for cache entries
	DefaultExpiration time.Duration = 0 // no expiration

	// Disabled globally turns off caching when true
	Disabled bool

	// Delimiter separates components in cache key strings
	Delimiter = ':'
)

// Cache defines the contract for cache implementations.
// Implementations must handle:
// - Serialization/deserialization using Marshal/Unmarshal functions
// - Error handling with proper error wrapping
// - Concurrency control for thread-safe operations
//
// All methods should return errors compatible with IsKeyNotFound check
// for consistent cache miss detection.

type Cache interface {
	// Del removes multiple entries from cache.
	//
	// Args:
	//   ctx: Context for request cancellation/timeout
	//   keys: Cache keys to delete
	//
	// Returns:
	//   affected: Number of successfully deleted entries
	//   err: Storage errors (e.g. connection issues), nil on success
	Del(ctx context.Context, keys ...string) (affected int64, err error)

	// Get retrieves raw byte slice from cache.
	//
	// Args:
	//   ctx: Context for request cancellation/timeout
	//   key: Cache entry identifier
	//
	// Returns:
	//   val: Cached byte slice (nil on cache miss)
	//   err: Storage errors or nil. Use IsKeyNotFound() to detect cache misses
	Get(ctx context.Context, key string) (val []byte, err error)

	// Set stores byte slice in cache with expiration.
	//
	// Args:
	//   ctx: Context for request cancellation/timeout
	//   key: Cache entry identifier
	//   val: Data to store (nil allowed for cache deletion)
	//   expire: TTL duration (<=0 means no expiration)
	//
	// Returns:
	//   error: Storage errors, nil on success
	Set(ctx context.Context, key string, val []byte, expire time.Duration) error
	// GetAsUint64 retrieves and converts cached value to unsigned integer.
	//
	// Performs strict conversion:
	// - Returns error if value is not 8-byte little-endian format
	// - Returns error if value exceeds uint64 range
	//
	// Args:
	//   key: Cache entry identifier
	//
	// Returns:
	//   val: Converted unsigned integer value
	//   err: Conversion errors or storage errors
	GetAsUint64(ctx context.Context, key string) (val uint64, err error)
	// SetUint64 stores unsigned integer in 8-byte little-endian format.
	//
	// Args:
	//   key: Cache entry identifier
	//   val: Unsigned integer to store
	//   expire: TTL duration (<=0 uses default expiration)
	//
	// Returns:
	//   error: Storage errors, nil on success
	SetUint64(ctx context.Context, key string, val uint64, expire time.Duration) error
	// IsKeyNotFound checks if error represents cache miss (key not exists).
	//
	// Implementation should return true for:
	// - Cache storage-specific "not found" errors
	// - Wrapped versions of these errors
	//
	// This enables reliable cache miss detection across implementations
	IsKeyNotFound(err error) bool
}

// CacheableEntity defines the contract for cacheable domain entities
// Used to enforce ID-based caching constraints
// Implementations should:
// - Use unsigned integer types for identifiers
// - Maintain immutable ID properties
type CacheableEntity[INT constraints.Unsigned] interface {
	ID() INT
}

// DelCachedEntity deletes cached data for specified key
func DelCachedEntity(ctx context.Context, cache Cache, key string) (affected int64, err error) {
	if Disabled {
		return 0, nil
	}

	if affected, err = cache.Del(ctx, key); err != nil {
		return affected, errors.WithStack(err)
	}
	return affected, nil
}

// GetEntityByID implements a dual-check mechanism:
// 1. Check cache first with singleflight deduplication
// 2. Fallback to database on cache miss
// 3. Repopulate cache with database result
//
// The cache-aside pattern prevents stale cache returns while
// singleflight prevents cache stampede
func GetEntityByID[T CacheableEntity[INT], INT constraints.Unsigned](
	ctx context.Context,
	cache Cache,
	entityKeyPrefix string,
	id INT,
	getEntityByID func(context.Context, INT) (*T, error),
) (*T, error) {
	// 0. When cache is disabled, directly query database
	if Disabled {
		one, err := getEntityByID(ctx, id)
		if err != nil {
			return nil, errors.WithStack(err)
		}
		return one, nil
	}

	// 1. First attempt to retrieve object from cache
	key := fmt.Sprintf("%s%c%d", entityKeyPrefix, Delimiter, id)

	data, err := cache.Get(ctx, key)
	if err != nil && !cache.IsKeyNotFound(err) {
		return nil, errors.WithStack(err)
	}

	if err == nil {
		// 2. If value exists, attempt deserialization and return object
		var one T
		if err = Unmarshal(data, &one); err == nil {
			return &one, nil
		}
		// If deserialization fails, continue execution flow
	}

	// 2. Query database for target ID
	entity, err, _ := singleFlightGroup.Do(key, func() (any, error) { // Prevent cache breakdown
		return getEntityByID(ctx, id)
	})
	if err != nil {
		return nil, errors.WithStack(err)
	}

	// 3. Serialize object and store in cache
	if data, err = Marshal(entity); err != nil {
		return nil, errors.WithStack(err)
	}
	if err = cache.Set(ctx, key, data, DefaultExpiration); err != nil {
		return nil, errors.WithStack(err)
	}

	// 4. Return fetched object
	return entity.(*T), nil
}

// GetEntitiesByID retrieves entity list with cache-aside pattern
// Implements:
// 1. Batch cache lookup with singleflight deduplication
// 2. Database fallback on cache miss
// 3. Cache population with serialized results
func GetEntitiesByID[T CacheableEntity[INT], INT constraints.Unsigned](
	ctx context.Context,
	cache Cache,
	entityKeyPrefix string,
	id INT,
	getEntitiesByID func(context.Context, INT) ([]*T, error),
) ([]*T, error) {
	var err error
	var items []*T

	// 0. When cache is disabled, directly query database
	if Disabled {
		if items, err = getEntitiesByID(ctx, id); err != nil {
			return nil, errors.WithStack(err)
		}
		return items, nil
	}

	// 1. First attempt to retrieve object from cache列表
	key := fmt.Sprintf("%s%citems%c%d", entityKeyPrefix, Delimiter, Delimiter, id)

	data, err := cache.Get(ctx, key)
	if err != nil && !cache.IsKeyNotFound(err) {
		return nil, errors.WithStack(err)
	}

	if err == nil {
		// 2. If cached data exists, deserialize and return the entity list
		if err = Unmarshal(data, &items); err == nil {
			return items, nil
		}
		// If deserialization fails, continue execution flow
	}

	// 2. Query database for target ID列表
	entities, err, _ := singleFlightGroup.Do(key, func() (any, error) {
		return getEntitiesByID(ctx, id)
	})
	if err != nil {
		return nil, err
	}

	items = entities.([]*T)
	if len(items) == 0 {
		return items, nil
	}

	// 3、序列化对象列表并存入缓存
	if data, err = Marshal(&items); err != nil {
		return nil, errors.WithStack(err)
	}
	if err = cache.Set(ctx, key, data, DefaultExpiration); err != nil {
		return nil, errors.WithStack(err)
	}

	return items, nil
}

// UniqueKey defines constraints for database unique keys
// Used in GetEntityByUniqueKey to enforce:
// - String or unsigned integer types
// - Single-field uniqueness (composite keys not supported)
type UniqueKey interface {
	~string | constraints.Unsigned
}

// GetEntityByUniqueKey implements two-level cache resolution:
// 1. UK -> PK lookup cache
// 2. PK -> Entity cache
// Features:
// - Prevents cache stampede with singleflight
// - Automatic cache population for both UK-PK and PK-Entity mappings
// Limitations:
// - Composite unique keys not supported
// - Fixed cache key format
func GetEntityByUniqueKey[T CacheableEntity[INT], INT constraints.Unsigned, UK UniqueKey](
	ctx context.Context,
	cache Cache,
	entityKeyPrefix string,
	ukKeyPrefix string,
	ukVal UK,
	getIDByUK func(context.Context, UK) (INT, error),
	getEntityByID func(context.Context, INT) (*T, error),
) (*T, error) {
	// 0. When cache is disabled, directly query database
	if Disabled {
		id, err := getIDByUK(ctx, ukVal)
		if err != nil {
			return nil, errors.WithStack(err)
		}
		one, err := getEntityByID(ctx, id)
		if err != nil {
			return nil, errors.WithStack(err)
		}
		return one, nil
	}

	// Key represents unique index value
	// Value stores corresponding primary key ID
	ukKey := fmt.Sprintf("%s%c%v", ukKeyPrefix, Delimiter, ukVal)

	// 1. Resolve primary key ID using unique index key
	u64ID, err := cache.GetAsUint64(ctx, ukKey)
	if err == nil {
		return GetEntityByID(ctx, cache, entityKeyPrefix, INT(u64ID), getEntityByID)
	}
	if !cache.IsKeyNotFound(err) { // Handle unexpected errors beyond cache miss
		return nil, errors.WithStack(err)
	}
	// 2. If the primary key ID of the database table is not found, call the query function to obtain the primary key ID value.
	id, err := getIDByUK(ctx, ukVal)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	// 3. Persist unique index to primary key mapping
	if err = cache.SetUint64(ctx, ukKey, uint64(id), -1); err != nil {
		return nil, errors.WithStack(err)
	}
	// 4. Retrieve entity using resolved primary key
	return GetEntityByID(ctx, cache, entityKeyPrefix, id, getEntityByID)
}
