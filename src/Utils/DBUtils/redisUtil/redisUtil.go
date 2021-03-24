/*
 * @Description:封装redis操作
 * @Author: Rocky Hoo
 * @Date: 2021-03-22 21:13:41
 * @LastEditTime: 2021-03-24 18:20:17
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
	var port = "6379"
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
	defer con.Close()
	if _, err := con.Do("RPUSH", collection, value); err != nil {
		log.Println(err.Error())
		debug.PrintStack()
		return err
	}
	return nil
}

func (r *Redis) LPush(collection string, value interface{}) error {
	con := r.pool.Get()
	defer con.Close()
	if _, err := con.Do("LPUSH", collection, value); err != nil {
		log.Println(err)
		debug.PrintStack()
		return err
	}
	return nil
}

func (r *Redis) RPop(collection string) (string, error) {
	con := r.pool.Get()
	defer con.Close()
	res, err := redis.String(con.Do("RPOP", collection))
	if err != nil {
		log.Println(err)
		debug.PrintStack()
		return "", err
	}
	return res, nil
}

func (r *Redis) LPop(collection string) (string, error) {
	con := r.pool.Get()
	defer con.Close()
	res, err := redis.String(con.Do("LPOP", collection))
	if err != nil {
		log.Println(err.Error())
		debug.PrintStack()
		return "", err
	}
	return res, nil
}

func NewClient() *Redis {
	return &Redis{
		pool: NewRedisPool(),
	}
}
