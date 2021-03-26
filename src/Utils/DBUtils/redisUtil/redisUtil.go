/*
 * @Description:封装redis操作
 * @Author: Rocky Hoo
 * @Date: 2021-03-22 21:13:41
 * @LastEditTime: 2021-03-26 19:56:12
 * @LastEditors: Please set LastEditors
 * @CopyRight:
 * Copyright (c) 2021 XiaoPeng Studio
 */
package RedisUtil

import (
	"log"
	"runtime/debug"
	"time"

	"github.com/cenkalti/backoff/v4"
	"github.com/gomodule/redigo/redis"
	"github.com/spf13/viper"
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

func (r *Redis) Ping() error{
	conn:=r.pool.Get()
	defer conn.Close()	
	_,err:=conn.Do("ping")
	return err
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

func InitRedis() (*Redis,error) {
	if RedisClient != nil {
		return RedisClient,nil
	}
	RedisClient = NewClient()
	bf := backoff.NewExponentialBackOff()
	err:=backoff.Retry(func() (err error) {
		err=RedisClient.Ping()
		if err!=nil{
			log.Println("waiting for redis activate...")
		}
		return err
	}, bf)
	if err!=nil{
		return nil,err
	}
	return RedisClient,nil
}

func Reconn(){
	//TODO:redis连接重连:捕获到连接断开的报错,并触发相应的重连机制(backoff)
}
