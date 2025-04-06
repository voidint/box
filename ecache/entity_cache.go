package ecache

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/pkg/errors"
	"golang.org/x/exp/constraints"
	"golang.org/x/sync/singleflight"
)

var (
	singleFlightGroup singleflight.Group

	// Marshal 默认序列化方案
	Marshal func(v any) ([]byte, error) = json.Marshal
	// Unmarshal 默认反序列化方案
	Unmarshal func(data []byte, v any) error = json.Unmarshal
	// Expiration 缓存默认过期时间
	Expiration time.Duration
	// Disabled 缓存开关
	Disabled bool
	// 缓存键名分隔符
	Delimiter = ':'
)

type Cache interface {
	Del(ctx context.Context, keys ...string) (affected int64, err error)
	Get(ctx context.Context, key string) (val []byte, err error)
	GetAsUint64(ctx context.Context, key string) (val uint64, err error)
	Set(ctx context.Context, key string, val []byte, expire time.Duration) error
	IsKeyNotFound(err error) bool
}

type CacheableEntity[INT constraints.Integer] interface {
	ID() INT
}

// DelCachedEntity 删除指定键的缓存数据
func DelCachedEntity(ctx context.Context, cache Cache, key string) (affected int64, err error) {
	if Disabled {
		return 0, nil
	}

	if affected, err = cache.Del(ctx, key); err != nil {
		return affected, errors.WithStack(err)
	}
	return affected, nil
}

// GetEntityByID 返回主键ID对应的实体对象。
// 查找顺序：主键ID --> 缓存中实体对象 --> 数据库中实体对象。
// 局限：
// 1、缓存 key格式无法自定义且不够内聚
func GetEntityByID[T CacheableEntity[INT], INT constraints.Integer](
	ctx context.Context,
	cache Cache,
	entityKeyPrefix string,
	id INT,
	getEntityByID func(context.Context, INT) (*T, error),
) (*T, error) {
	// 0、若未启用缓存，则直接查询数据库。
	if Disabled {
		one, err := getEntityByID(ctx, id)
		if err != nil {
			return nil, errors.WithStack(err)
		}
		return one, nil
	}

	// 1、优先尝试从缓存中获取对象
	key := fmt.Sprintf("%s%c%d", entityKeyPrefix, Delimiter, id)

	data, err := cache.Get(ctx, key)
	if err != nil && !cache.IsKeyNotFound(err) {
		return nil, errors.WithStack(err)
	}

	if err == nil {
		// 2、若获取到键对应的值，尝试反序列化，并返回对象。
		var one T
		if err = Unmarshal(data, &one); err == nil {
			return &one, nil
		}
		// 若反序列化失败，不返回，允许继续向下执行。
	}

	// 2、从数据库查询指定ID对象
	entity, err, _ := singleFlightGroup.Do(key, func() (any, error) { // 防止缓存击穿
		return getEntityByID(ctx, id)
	})
	if err != nil {
		return nil, errors.WithStack(err)
	}

	// 3、序列化对象并存入Redis
	if data, err = Marshal(entity); err != nil {
		return nil, errors.WithStack(err)
	}
	if err = cache.Set(ctx, key, data, Expiration); err != nil {
		return nil, errors.WithStack(err)
	}

	// 4、返回查询到的对象
	return entity.(*T), nil
}

// GetEntitiesByID 从Redis缓存中读取指定实体列表。若缓存中不存在，则从指定方法中读取并存入Redis缓存。
func GetEntitiesByID[T CacheableEntity[INT], INT constraints.Integer](
	ctx context.Context,
	cache Cache,
	entityKeyPrefix string,
	id INT,
	getEntitiesByID func(context.Context, INT) ([]*T, error),
) ([]*T, error) {
	var err error
	var items []*T

	// 0、若未启用缓存，则直接查询数据库。
	if Disabled {
		if items, err = getEntitiesByID(ctx, id); err != nil {
			return nil, errors.WithStack(err)
		}
		return items, nil
	}

	// 1、优先尝试从Redis中获取对象列表
	key := fmt.Sprintf("%s%citems%c%d", entityKeyPrefix, Delimiter, Delimiter, id)

	data, err := cache.Get(ctx, key)
	if err != nil && !cache.IsKeyNotFound(err) {
		return nil, errors.WithStack(err)
	}

	if err == nil {
		// 2、若获取到键对应的值，尝试反序列化，并返回对象列表。
		if err = Unmarshal(data, &items); err == nil {
			return items, nil
		}
		// 若反序列化失败，不返回，允许继续向下执行。
	}

	// 2、从数据库查询指定ID对象列表
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

	// 3、序列化对象列表并存入Redis
	if data, err = Marshal(&items); err != nil {
		return nil, errors.WithStack(err)
	}
	if err = cache.Set(ctx, key, data, Expiration); err != nil {
		return nil, errors.WithStack(err)
	}

	return items, nil
}

// UniqueKey 唯一索引约束
type UniqueKey interface {
	~string | constraints.Integer
}

// GetEntityByUniqueKey 返回唯一索引值对应的实体对象。
// 查找顺序：唯一索引值 --> Redis中主键ID --> 数据库中主键ID --> Redis中实体对象 --> 数据库中实体对象。
// 局限：
// 1、Redis key格式无法自定义
// 2、不支持联合唯一索引
// 3、多一次查询
func GetEntityByUniqueKey[T CacheableEntity[INT], INT constraints.Integer, UK UniqueKey](
	ctx context.Context,
	cache Cache,
	entityKeyPrefix string,
	ukKeyPrefix string,
	ukVal UK,
	getIDByUK func(context.Context, UK) (INT, error),
	getEntityByID func(context.Context, INT) (*T, error),
) (*T, error) {
	// 0、若未启用缓存，则直接查询数据库。
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

	// 键：唯一索引值
	// 值：主键ID
	ukKey := fmt.Sprintf("%s%c%v", ukKeyPrefix, Delimiter, ukVal)

	// 1、根据数据库表唯一索引key找到数据库表主键ID
	u64ID, err := cache.GetAsUint64(ctx, ukKey)
	if err == nil {
		// 2、若找到数据库表主键ID，则根据主键ID尝试返回实体对象。
		return GetEntityByID(ctx, cache, entityKeyPrefix, INT(u64ID), getEntityByID)
	}
	if !cache.IsKeyNotFound(err) { // 发生了除「Redis key不存在」之外的其他未知错误
		return nil, errors.WithStack(err)
	}
	// 2、若未找到数据库表主键ID，则调用查询函数获取主键ID值。
	id, err := getIDByUK(ctx, ukVal)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	// 3、保存唯一索引值（键）与主键ID（值）关系
	if err = cache.Set(ctx, ukKey, fmt.Appendf(nil, "%d", id), 0); err != nil {
		return nil, errors.WithStack(err)
	}
	// 4、根据主键ID尝试返回实体对象
	return GetEntityByID(ctx, cache, entityKeyPrefix, id, getEntityByID)
}
