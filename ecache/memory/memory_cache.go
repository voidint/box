package memory

import (
	"context"
	"strconv"
	"time"

	"github.com/alicebob/miniredis/v2"
	"github.com/pkg/errors"
)

type cache struct {
	rdb *miniredis.Miniredis
}

func NewCache() (*cache, error) {
	rdb, err := miniredis.Run()
	if err != nil {
		return nil, errors.WithStack(err)
	}
	return &cache{
		rdb: rdb,
	}, nil
}

func (c *cache) Del(ctx context.Context, keys ...string) (affected int64, err error) {
	for _, key := range keys {
		if c.rdb.Del(key) {
			affected++
		}
	}
	return affected, nil
}

func (c *cache) Get(ctx context.Context, key string) (val []byte, err error) {
	data, err := c.rdb.Get(key)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	return []byte(data), nil
}

func (c *cache) GetAsUint64(ctx context.Context, key string) (val uint64, err error) {
	data, err := c.rdb.Get(key)
	if err != nil {
		return 0, errors.WithStack(err)
	}
	return strconv.ParseUint(data, 10, 64)
}
func (c *cache) Set(ctx context.Context, key string, val []byte, expire time.Duration) (err error) {
	if err = c.rdb.Set(key, string(val)); err != nil {
		return errors.WithStack(err)
	}
	c.rdb.SetTTL(key, expire)
	return nil
}
func (c *cache) IsKeyNotFound(err error) bool {
	return errors.Is(err, miniredis.ErrKeyNotFound)
}
