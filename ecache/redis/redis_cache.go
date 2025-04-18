// Copyright (c) 2025 voidint <voidint@126.com>
//
// Permission is hereby granted, free of charge, to any person obtaining a copy of
// this software and associated documentation files (the "Software"), to deal in
// the Software without restriction, including without limitation the rights to
// use, copy, modify, merge, publish, distribute, sublicense, and/or sell copies of
// the Software, and to permit persons to whom the Software is furnished to do so,
// subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in all
// copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY, FITNESS
// FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE AUTHORS OR
// COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY, WHETHER
// IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM, OUT OF OR IN
// CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.

package redis

import (
	"context"
	"time"

	"github.com/pkg/errors"
	"github.com/redis/go-redis/v9"
	"github.com/voidint/box/ecache"
)

var _ ecache.Cache = (*cache)(nil) // Ensure cache implements ecache.Cache interface.

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
