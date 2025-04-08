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

package memory

import (
	"context"
	"time"

	gocache "github.com/patrickmn/go-cache"
	"github.com/pkg/errors"
)

var (
	// ErrKeyNotFound indicates the requested key does not exist in the cache
	ErrKeyNotFound = errors.New("key not found")

	// ErrUnexpectedType occurs when attempting to retrieve data with unexpected type
	ErrUnexpectedType = errors.New("unknown data type")
)

type cache struct {
	db *gocache.Cache
}

// NewCache creates an in-memory cache instance with configurable expiration.
// defaultExpiration: default TTL for cache entries (use time.Duration(0) for no expiration)
// cleanupInterval: interval for automatic removal of expired entries (use time.Duration(0) to disable)
func NewCache(defaultExpiration, cleanupInterval time.Duration) *cache {
	return &cache{
		db: gocache.New(defaultExpiration, cleanupInterval),
	}
}

// Del removes multiple entries from the cache.
// Returns number of deleted items and any potential error (currently always nil)
func (c *cache) Del(ctx context.Context, keys ...string) (affected int64, err error) {
	for _, key := range keys {
		c.db.Delete(key)
	}
	return int64(len(keys)), nil
}

// Get retrieves a value as byte slice from the cache.
// Returns ErrKeyNotFound if the key doesn't exist
// Returns ErrUnexpectedType if value cannot be cast to []byte
func (c *cache) Get(ctx context.Context, key string) (val []byte, err error) {
	v, ok := c.db.Get(key)
	if !ok {
		return nil, errors.WithStack(ErrKeyNotFound)
	}
	item := v.(gocache.Item)
	if val, ok = item.Object.([]byte); !ok {
		return nil, errors.WithStack(ErrUnexpectedType)
	}
	return val, nil
}

// GetAsUint64 retrieves a numeric value from the cache as uint64.
// Returns ErrKeyNotFound if the key doesn't exist
// Returns ErrUnexpectedType if value cannot be cast to uint64
func (c *cache) GetAsUint64(ctx context.Context, key string) (val uint64, err error) {
	v, ok := c.db.Get(key)
	if !ok {
		return 0, errors.WithStack(ErrKeyNotFound)
	}
	item := v.(gocache.Item)
	if val, ok = item.Object.(uint64); !ok {
		return 0, errors.WithStack(ErrUnexpectedType)
	}
	return val, nil
}

// Set stores a byte slice in the cache with expiration.
// expire: 0 uses default expiration, <0 means no expiration
func (c *cache) Set(ctx context.Context, key string, val []byte, expire time.Duration) error {
	c.db.Set(key, val, expire)
	return nil
}

// SetUint64 stores a numeric value in the cache with expiration.
// expire: 0 uses default expiration, <0 means no expiration
func (c *cache) SetUint64(ctx context.Context, key string, val uint64, expire time.Duration) error {
	c.db.Set(key, val, expire)
	return nil
}

// IsKeyNotFound checks if an error indicates missing key
// Helps determine error type without direct dependency on package errors
func (c *cache) IsKeyNotFound(err error) bool {
	return errors.Is(err, ErrKeyNotFound)
}
