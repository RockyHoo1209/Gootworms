/*
 * @Description:封装redis操作
 * @Author: Rocky Hoo
 * @Date: 2021-03-22 21:13:41
 * @LastEditTime: 2021-03-23 01:07:35
 * @LastEditors: Please set LastEditors
 * @CopyRight:
 * Copyright (c) 2021 XiaoPeng Studio
 */
package redisUtil

import (
	"log"
	"runtime/debug"
	"time"

	"github.com/gomodule/redigo/redis"
	"github.com/spf13/viper"
)

type Redis struct {
	pool *redis.Pool
}

func NewRedisPool() *redis.Pool {
	var address = viper.GetString("redis.address")
	var password = viper.GetString("redis.password")
	var database = viper.GetString("redis.database")
	var port = viper.GetString("redis.port")
	var redisUrl = ""
	if password == "" {
		redisUrl = "redis://" + address + ":" + port + "/" + database
	} else {
		redisUrl = "redis://x:" + password + "@" + address + ":" + port + "/" + database
	}
	return &redis.Pool{
		Dial: func() (redis.Conn, error) {
			return redis.DialURL(redisUrl, redis.DialConnectTimeout(time.Second*10),
				redis.DialReadTimeout(time.Second*10), redis.DialWriteTimeout(time.Second*60))
		},
		TestOnBorrow: func(c redis.Conn, t time.Time) error {
			if time.Since(t) < time.Minute {
				return nil
			}
			_, err := c.Do("PING")
			return err
		},
		MaxIdle:         10,
		MaxActive:       0,
		IdleTimeout:     300 * time.Second,
		Wait:            false,
		MaxConnLifetime: 0,
	}
}

func (r *Redis) RPush(collection string, value interface{}) error {
	con := r.pool.Get()
	if _, err := con.Do("RPUSH", collection, value); err != nil {
		log.Println(err)
		debug.PrintStack()
		return err
	}
	return nil
}

func (r *Redis) LPush(collection string, value interface{}) error {
	con := r.pool.Get()
	if _, err := con.Do("LPUSH", collection, value); err != nil {
		log.Println(err)
		debug.PrintStack()
		return err
	}
	return nil
}

func (r *Redis) RPop(collection string, value interface{}) error {
	con := r.pool.Get()
	if _, err := con.Do("RPOP", collection, value); err != nil {
		log.Println(err)
		debug.PrintStack()
		return err
	}
	return nil
}

func (r *Redis) LPop(collection string, value interface{}) error {
	con := r.pool.Get()
	if _, err := con.Do("LPOP", collection, value); err != nil {
		log.Println(err)
		debug.PrintStack()
		return err
	}
	return nil
}

func NewClient() *Redis {
	return &Redis{
		pool: NewRedisPool(),
	}
}
