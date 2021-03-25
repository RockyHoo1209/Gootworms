/*
 * @Description:
 * @Author: Rocky Hoo
 * @Date: 2021-03-25 09:43:05
 * @LastEditTime: 2021-03-25 10:22:32
 * @LastEditors: Please set LastEditors
 * @CopyRight:
 * Copyright (c) 2021 XiaoPeng Studio
 */
package MongoUtil

import (
	"log"
	"testing"
)

func TestMongoUtil(*testing.T) {
	InitMongo()
	s, c := GetCollection("nodes")
	defer s.Close()
	data := map[string]string{"title": "MongoDB 教程",
		"description": "MongoDB 是一个 Nosql 数据库",
		"by":          "菜鸟教程",
		"url":         "http://www.runoob.com",
		"likes":       "100",
	}
	err := c.Insert(data)
	if err != nil {
		log.Fatal(err.Error())
	}
}
