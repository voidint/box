package redis

import (
	"context"
	"time"

	"github.com/pkg/errors"
	"github.com/redis/go-redis/v9"
)

// cache implements Redis-based caching solution.
// It wraps go-redis client to provide standard cache interface.
type cache struct {
	rdb *redis.Client
}

// NewCache creates a Redis cache instance.
// client: Configured go-redis client connection
func NewCache(client *redis.Client) *cache {
	return &cache{
		rdb: client,
	}
}

// Del removes multiple keys from cache.
// Returns number of deleted keys and any error encountered.
// Wraps redis DEL command errors with stack trace.
func (c *cache) Del(ctx context.Context, keys ...string) (affected int64, err error) {
	if affected, err = c.rdb.Del(ctx, keys...).Result(); err != nil {
		return affected, errors.WithStack(err)
	}
	return affected, nil
}

// Get retrieves byte slice value for given key.
// Returns redis.Nil error when key does not exist.
// Wraps underlying redis GET command errors.
func (c *cache) Get(ctx context.Context, key string) (val []byte, err error) {
	if val, err = c.rdb.Get(ctx, key).Bytes(); err != nil {
		return nil, errors.WithStack(err)
	}
	return val, nil
}

// Set stores value with TTL expiration.
// expire: Time-to-live duration (<=0 means no expiration)
// Wraps redis SET command errors with stack trace.
func (c *cache) Set(ctx context.Context, key string, val []byte, expire time.Duration) (err error) {
	if err = c.rdb.Set(ctx, key, val, expire).Err(); err != nil {
		return errors.WithStack(err)
	}
	return nil
}

// GetAsUint64 retrieves uint64 value for given key.
// Returns 0 and error if value cannot be converted.
// Wraps underlying redis GET command errors.
func (c *cache) GetAsUint64(ctx context.Context, key string) (val uint64, err error) {
	if val, err = c.rdb.Get(ctx, key).Uint64(); err != nil {
		return 0, errors.WithStack(err)
	}
	return val, nil
}

// SetUint64 stores uint64 value with TTL expiration.
// expire: Time-to-live duration (<=0 means no expiration)
// Wraps redis SET command errors with stack trace.
func (c *cache) SetUint64(ctx context.Context, key string, val uint64, expire time.Duration) (err error) {
	if err = c.rdb.Set(ctx, key, val, expire).Err(); err != nil {
		return errors.WithStack(err)
	}
	return nil
}

// IsKeyNotFound checks if error represents missing key.
// Returns true if error is redis.Nil.
func (c *cache) IsKeyNotFound(err error) bool {
	return errors.Is(err, redis.Nil)
}
