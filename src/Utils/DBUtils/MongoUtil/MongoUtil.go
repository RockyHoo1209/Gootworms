/*
 * @Description: MongoDB操作封装(单例,复用)
 * @Author: Rocky Hoo
 * @Date: 2021-03-24 17:41:52
 * @LastEditTime: 2021-04-27 17:17:09
 * @LastEditors: Please set LastEditors
 * @CopyRight:
 * Copyright (c) 2021 XiaoPeng Studio
 */
package MongoUtil

import (
	"log"
	"main/src/Utils/ConfigUtil"
	"net"
	"time"

	"github.com/cenkalti/backoff/v4"
	"github.com/globalsign/mgo"
	"github.com/globalsign/mgo/bson"
)

type Url struct {
	Key string        `json:"key"`
}

type Response struct {
	Key     string      `json:"key"`
	Context interface{} `json:"context"`
}

var Session *mgo.Session
var dialInfo *mgo.DialInfo

//通过sesstion.copy的refresh来刷新socket连接缓存,解决并发sesstion共享问题:https://blog.csdn.net/u010649766/article/details/79818240
func GetSesstion() *mgo.Session {
	return Session.Copy()
}

func GetCollection(collectionName string) (session *mgo.Session, collection *mgo.Collection) {
	session = GetSesstion()
	db := session.DB("Gootworms")
	collection = db.C(collectionName)
	return session, collection
}

/**
 * @description: 插入数据库操作，如果没有则插入，存在则更新
 * @param  {*}
 * @return {*}
 * @param {*} collectionName
 * @param {string} key:MongoDb插入主键
 * @param {interface{}} data:将要插入的数据
 */
func Insert(collectionName, key string, data interface{}) error {
	session, collection := GetCollection(collectionName)
	defer session.Close()
	err := collection.Insert(data)
	if err != nil {
		_, err := collection.Upsert(bson.M{"key": key}, data)
		if nil != err {
			log.Printf("MongoUtil-Insert %s Faied:%s\n", collectionName, err.Error())
			return err
		}
	}
	log.Printf("MongoUtil-Insert %s Success!\n", collectionName)
	return nil
}

func EnsureIndex(collectionName string,keyindx *mgo.Index){
	session, collection := GetCollection(collectionName)
	defer session.Close()
	collection.EnsureIndex(*keyindx)
}

func InsertUrls(url string) {
	Insert("Urls", url, &Url{
		Key: url,
	})
}

func InsertResponse(url string, response interface{}) {
	Insert("Response", url, &Response{
		Key:     url,
		Context: response,
	})
}

func InsertResult(result map[string]interface{}) {
	if result == nil || len(result) < 1 {
		return
	}
	key_name:=ConfigUtil.GetString("server.resultKey")
	key_value,ok:=result[key_name];if !ok{
		return
	}
	result["key"]=key_value
	Insert("Result", key_value.(string), result)
}

func InitMongo() {
	MongoDBHost := ConfigUtil.GetString("mongo.address")
	MongoDBPort := ConfigUtil.GetString("mongo.port")
	MongoDBDataBase := ConfigUtil.GetString("mongo.db")
	MongoDBUserName := ConfigUtil.GetString("mongo.username")
	MongoDBPassword := ConfigUtil.GetString("mongo.password")
	//连接池的数量默认4096,开销的压力是比较大的
	MongoDBPoolLimit := 100
	MongoDBSource := ""
	addr := net.JoinHostPort(MongoDBHost, MongoDBPort)
	timeout := time.Second * 10
	if Session == nil {
		dialInfo = &mgo.DialInfo{
			Addrs:        []string{addr},
			Timeout:      timeout,
			Database:     MongoDBDataBase,
			PoolLimit:    MongoDBPoolLimit,
			PoolTimeout:  timeout,
			ReadTimeout:  timeout,
			WriteTimeout: timeout,
			AppName:      "Gootworms",
			MinPoolSize:  10,
		}
		if MongoDBUserName != "" {
			dialInfo.Username = MongoDBUserName
			dialInfo.Password = MongoDBPassword
			dialInfo.Source = MongoDBSource
		}
		//指数退避重试连接
		bp := backoff.NewExponentialBackOff()
		backoff.Retry(func() (err error) {
			Session, err = mgo.DialWithInfo(dialInfo)
			if err != nil {
				log.Println(err.Error())
				return
			}
			return
		}, bp)
	}

	//设置key唯一
	keyindx := &mgo.Index{
		Key:    []string{"key"},
		Unique: true,
	}
	EnsureIndex("Urls",keyindx)
	EnsureIndex("Reponse",keyindx)
	EnsureIndex("nodes",keyindx)
	EnsureIndex("Result",keyindx)
}
