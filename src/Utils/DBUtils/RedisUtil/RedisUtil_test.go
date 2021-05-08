/*
 * @Description:
 * @Author: Rocky Hoo
 * @Date: 2021-03-23 08:44:59
 * @LastEditTime: 2021-04-02 13:10:49
 * @LastEditors: Please set LastEditors
 * @CopyRight:
 * Copyright (c) 2021 XiaoPeng Studio
 */
package RedisUtil

import (
	"fmt"
	"testing"
	"main/src/Utils/TaskUtil"
	"log"
)

func TestRedisUtil(*testing.T) {
	r,err := InitRedis()
	taskUtil:=TaskUtil.NewTaskUtil()
	if err!=nil{
		log.Println(err.Error())
	}
	err = r.RPush("hello", 1000)
	if err != nil {
		fmt.Printf(err.Error())
	}
	res, err := r.LPop("hello")
	taskUtil.Append(func(...interface{}){
		
	})
	if err.Error()=="EOF"{
		Reconn()
		res, err = r.LPop("hello")
	}

	if err != nil {
		fmt.Printf(err.Error())
	}
	fmt.Println(res)
}