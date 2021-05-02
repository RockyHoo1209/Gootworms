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
	"fmt"
	"log"
	"main/src/Utils/ConfigUtil"
	"testing"
)

func TestMongoUtil(*testing.T) {
	ConfigUtil.InitConfig("config")
	InitMongo()
	for i:=1;i<=6;i++{
	data := map[string]interface{}{
		"key":"1",
		// "_id":bson.NewObjectId(),
		"title": "MongoDB 程"+fmt.Sprint(i),
		"description": "MongoDB 是一个 Nosql 数据库",
		"by":          "菜鸟教程",
		"url":         "https://www.runoob.com",
		"likes":       "100",
	}
	err := Insert("nodes","1",data);if err!=nil{
		log.Println(err.Error())	
	}
	}
}
