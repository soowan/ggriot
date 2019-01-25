package cache

import (
	"github.com/go-redis/redis"
)

// RedisConn is the imported Client from the user.
var RedisConn *redis.Client

// Enabled this is checked to see if ggriot should call the postgres db first or skip calling cache.
var Enabled = false

// UseCache will open a connection to a postgres server that will be used as a cache server.
func UseCache(c *redis.Client) (err error) {
	err = c.Ping().Err()
	if err != nil {
		return err
	}
	RedisConn = c

	Enabled = true
	return nil
}
