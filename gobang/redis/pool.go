package redis

import (
	"github.com/gomodule/redigo/redis"
	"gobang/config"
	"time"
)

var (
	pool *redis.Pool
)

func init() {
	addr := config.Config.Get("redis.addr").(string)
	password := config.Config.Get("redis.pwd").(string)

	pool = &redis.Pool{
		Dial: func() (conn redis.Conn, err error) {
			c, err := redis.Dial("tcp", addr)
			if err != nil {
				return nil, err
			}
			if _, err := c.Do("AUTH", password); err != nil {
				c.Close()
				return nil, err
			}
			return c, nil
		},
		MaxIdle:     3,
		IdleTimeout: 10 * time.Second,
	}
}
