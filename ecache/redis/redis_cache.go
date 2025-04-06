package redis

import (
	"context"
	"time"

	"github.com/pkg/errors"
	"github.com/redis/go-redis/v9"
)

type cache struct {
	rdb *redis.Client
}

func NewCache(client *redis.Client) *cache {
	return &cache{
		rdb: client,
	}
}

func (c *cache) Del(ctx context.Context, keys ...string) (affected int64, err error) {
	if affected, err = c.rdb.Del(ctx, keys...).Result(); err != nil {
		return affected, errors.WithStack(err)
	}
	return affected, nil
}

func (c *cache) Get(ctx context.Context, key string) (val []byte, err error) {
	if val, err = c.rdb.Get(ctx, key).Bytes(); err != nil {
		return nil, errors.WithStack(err)
	}
	return val, nil
}

func (c *cache) GetAsUint64(ctx context.Context, key string) (val uint64, err error) {
	if val, err = c.rdb.Get(ctx, key).Uint64(); err != nil {
		return 0, errors.WithStack(err)
	}
	return val, nil
}

func (c *cache) Set(ctx context.Context, key string, val []byte, expire time.Duration) (err error) {
	if err = c.rdb.Set(ctx, key, val, expire).Err(); err != nil {
		return errors.WithStack(err)
	}
	return nil
}

func (c *cache) IsKeyNotFound(err error) bool {
	return errors.Is(err, redis.Nil)
}
