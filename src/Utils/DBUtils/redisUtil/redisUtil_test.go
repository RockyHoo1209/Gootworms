/*
 * @Description:
 * @Author: Rocky Hoo
 * @Date: 2021-03-23 08:44:59
 * @LastEditTime: 2021-03-23 22:32:13
 * @LastEditors: Please set LastEditors
 * @CopyRight:
 * Copyright (c) 2021 XiaoPeng Studio
 */
package redisUtil

import (
	"fmt"
	"testing"
)

func TestRedisUtil(*testing.T) {
	r := NewClient()
	err := r.RPush("hello", 1000)
	if err != nil {
		fmt.Printf(err.Error())
	}
	res, err := r.LPop("hello")
	if err != nil {
		fmt.Printf(err.Error())
	}
	fmt.Println(res)
}
