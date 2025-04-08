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
