package cache

import (
	"errors"
	"fmt"
	"github.com/cespare/xxhash"
	"github.com/go-redis/redis"
	"github.com/json-iterator/go"
)

var (
	// ErrNoCache is the error that is used when requested call doesn't get cached.
	ErrNoCache = errors.New("call will not be cached")

	// ErrNoData is the error that is used when data that doesn't expire isn't in database
	ErrNoData = errors.New("call is not cached")
)

// ReadCache will return the cached call from redis.
func ReadCache(s interface{}, cp *CachedParams) error {
	if cp.Cached == true && Enabled == true {
		val, err := RedisConn.Get(fmt.Sprint(xxhash.Sum64String(cp.CallKey+"_"+cp.CallType))).Bytes()
		switch err {
		case redis.Nil:
			return ErrNoData
		case nil:
		default:
			return err
		}

		if err := jsoniter.Unmarshal(val, &s); err != nil {
			return err
		}
		return nil
	}
	return ErrNoCache
}

// StoreCache will store the call into the redis cache.
func StoreCache(cp *CachedParams, resp []byte) (err error) {
	if cp.Cached == true && Enabled == true {
		switch cp.Expire {
		case true:
			if err := RedisConn.Set(fmt.Sprint(xxhash.Sum64String(cp.CallKey+"_"+cp.CallType)), resp, cp.Expiration).Err(); err != nil {
				return err
			}
		case false:
			if err := RedisConn.Set(fmt.Sprint(xxhash.Sum64String(cp.CallKey+"_"+cp.CallType)), resp, 0).Err(); err != nil {
				return err
			}
		}
	}
	return nil
}
