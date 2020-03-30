package redis

import (
	"github.com/gomodule/redigo/redis"
	"time"
)

var (
	pool *redis.Pool
)

func init() {
	addr := "150.158.104.248:6379"
	password := "qwerty"

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
