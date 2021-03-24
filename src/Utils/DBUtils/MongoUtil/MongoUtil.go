/*
 * @Description: MongoDB操作封装
 * @Author: Rocky Hoo
 * @Date: 2021-03-24 17:41:52
 * @LastEditTime: 2021-03-24 21:27:45
 * @LastEditors: Please set LastEditors
 * @CopyRight:
 * Copyright (c) 2021 XiaoPeng Studio
 */
package MongoUtil

import (
	"log"
	"net"
	"time"

	"github.com/cenkalti/backoff/v4"
	"github.com/globalsign/mgo"
)

var Session *mgo.Session
var dialInfo *mgo.DialInfo

//通过sesstion.copy的refresh来刷新socket连接缓存,解决并发sesstion共享问题:https://blog.csdn.net/u010649766/article/details/79818240
func GetSesstion() *mgo.Session {
	return Session.Copy()
}

func GetCollection(collectionName string) (session *mgo.Session, collection *mgo.Collection) {
	session = GetSesstion()
	db := session.DB("mongo.db")
	collection = db.C(collectionName)
	return session, collection
}

func InitMongo() {
	MongoDBHost := "localhost"
	MongoDBPort := "2107"
	MongoDBDataBase := "mongo.db"
	MongoDBUserName := ""
	MongoDBPassword := ""
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

	//设置主键唯一
	keyindx := mgo.Index{
		Key:    []string{"key"},
		Unique: true,
	}
	session, collection := GetCollection("nodes")
	defer session.Close()
	collection.EnsureIndex(keyindx)
}
