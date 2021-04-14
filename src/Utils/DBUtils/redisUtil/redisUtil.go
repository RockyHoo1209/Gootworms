/*
 * @Description:封装redis操作(redis实rpc通信)
 * @Author: Rocky Hoo
 * @Date: 2021-03-22 21:13:41
 * @LastEditTime: 2021-04-12 23:00:16
 * @LastEditors: Please set LastEditors
 * @CopyRight:
 * Copyright (c) 2021 XiaoPeng Studio
 */
package RedisUtil

import (
	"log"
	// "runtime/debug"
	"time"

	"github.com/cenkalti/backoff/v4"
	"github.com/gomodule/redigo/redis"
	// "github.com/spf13/viper"
)

type Redis struct {
	pool *redis.Pool
}

var RedisClient *Redis

/**
 * @description: 新建一个redis连接池;生产环境推荐使用连接池，有效减少tcp建立连接的开销
 * @param  {*}
 * @return {*}
 */
func NewRedisPool() *redis.Pool {
	var address = "192.168.0.103" //viper.GetString("redis.address")
	var password = ""         //viper.GetString("redis.password")
	var database = "1"        //viper.GetString("redis.database")
	var port = "6379"         //viper.GetString("redis.port")
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
		// 每次从连接池中借用连接时测试连接可用性
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
		// debug.PrintStack()
		return err
	}
	return nil
}

func (r *Redis) LPush(collection string, value interface{}) error {
	con := r.pool.Get()
	defer con.Close()
	if _, err := con.Do("LPUSH", collection, value); err != nil {
		log.Println(err)
		// debug.PrintStack()
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
		// debug.PrintStack()
		return "", err
	}
	return res, nil
}

func (r *Redis) Ping() error {
	conn := r.pool.Get()
	defer conn.Close()
	_, err := conn.Do("ping")
	return err
}

func (r *Redis) LPop(collection string) (interface{}, error) {
	con := r.pool.Get()
	defer con.Close()
	res, err := redis.String(con.Do("LPOP", collection))
	if err != nil {
		log.Println(err.Error())
		// debug.PrintStack()
		return "", err
	}
	return res, nil
}

//阻塞
func (r *Redis) BLPop(collection string,timeout int) (string, error) {
	if timeout<=0{
		timeout=60
	}
	con := r.pool.Get()
	defer con.Close()
	res, err := redis.Strings(con.Do("BLPOP", collection,timeout))
	if err != nil {
		log.Println(err.Error())
		// debug.PrintStack()
		return "", err
	}
	return res[1], nil
}

func NewClient() *Redis {
	return &Redis{
		pool: NewRedisPool(),
	}
}

/**
 * @description: 初始化redis对外接口
 * @param  {*}
 * @return {*}
 */
func InitRedis() (*Redis, error) {
	if RedisClient == nil {
		RedisClient = NewClient()
	}
	bf := backoff.NewExponentialBackOff()
	err := backoff.Retry(func() (err error) {
		err = RedisClient.Ping()
		if err != nil {
			log.Println(err.Error())
			log.Println("waiting for redis activate...")
		}
		return err
	}, bf)
	if err != nil {
		return nil, err
	}
	return RedisClient, nil
}

/**
 * @description: 重连redis服务端
 * @param  {*}
 * @return {*}
 */
func Reconn() {
	InitRedis()
}

