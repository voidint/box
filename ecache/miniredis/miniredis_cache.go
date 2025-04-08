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

package miniredis

import (
	"context"
	"strconv"
	"time"

	"github.com/alicebob/miniredis/v2"
	"github.com/pkg/errors"
)

// cache implements key-value storage with TTL management using miniredis.
// It provides thread-safe operations for cache item manipulation.
type cache struct {
	rdb *miniredis.Miniredis
}

// NewCache creates a new miniredis-based cache instance.
// Returns error if failed to start the embedded redis server.
func NewCache() (*cache, error) {
	rdb, err := miniredis.Run()
	if err != nil {
		return nil, errors.WithStack(err)
	}
	return &cache{
		rdb: rdb,
	}, nil
}

// Del removes multiple keys from cache.
// keys: variadic parameter specifying keys to delete
// Returns number of successfully deleted keys and possible error
func (c *cache) Del(ctx context.Context, keys ...string) (affected int64, err error) {
	for _, key := range keys {
		if c.rdb.Del(key) {
			affected++
		}
	}
	return affected, nil
}

// Get retrieves value by key from cache.
// key: cache item identifier
// Returns byte slice and error (ErrKeyNotFound when missing)
func (c *cache) Get(ctx context.Context, key string) (val []byte, err error) {
	data, err := c.rdb.Get(key)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	return []byte(data), nil
}

// GetAsUint64 retrieves value and converts to uint64.
// Performs type conversion using strconv.ParseUint
// Returns conversion error for non-numeric values
func (c *cache) GetAsUint64(ctx context.Context, key string) (val uint64, err error) {
	data, err := c.rdb.Get(key)
	if err != nil {
		return 0, errors.WithStack(err)
	}
	return strconv.ParseUint(data, 10, 64)
}

// Set stores key-value pair with expiration time.
// expire: time-to-live duration (<=0 means no expiration)
// Returns error if failed to set the value
func (c *cache) Set(ctx context.Context, key string, val []byte, expire time.Duration) (err error) {
	if err = c.rdb.Set(key, string(val)); err != nil {
		return errors.WithStack(err)
	}
	c.rdb.SetTTL(key, expire)
	return nil
}

// IsKeyNotFound checks if error indicates missing key
// Wraps miniredis.ErrKeyNotFound detection
func (c *cache) IsKeyNotFound(err error) bool {
	return errors.Is(err, miniredis.ErrKeyNotFound)
}
