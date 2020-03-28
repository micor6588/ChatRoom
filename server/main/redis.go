package main

import (
	"time"

	"github.com/garyburd/redigo/redis"
)

//定义一个全局的pool
var pool *redis.Pool

// initPool 初始化连接池
func initPool(address string, maxIdle, maxActive int, idleTimeout time.Duration) {
	pool = &redis.Pool{
		MaxIdle:     maxIdle,     //最大链接数量
		MaxActive:   maxActive,   //表示和数据库的最大链接数，0表示没有限制
		IdleTimeout: idleTimeout, //最大空闲时间
		Dial: func() (redis.Conn, error) { //初始化代码连接哪个IP的redis
			return redis.Dial("tcp", address)
		},
	}
}
