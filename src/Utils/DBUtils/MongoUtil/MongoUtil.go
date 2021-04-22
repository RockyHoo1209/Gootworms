/*
 * @Description: MongoDB操作封装(单例,复用)
 * @Author: Rocky Hoo
 * @Date: 2021-03-24 17:41:52
 * @LastEditTime: 2021-04-22 12:41:20
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
	Id  bson.ObjectId `bson:"_id"`
	Url string        `json:"url"`
}

type Response struct {
	Url     string `json:"Url"`
	Context string `json:"context"`
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

func InsertUrls(url string) {
	session, collection := GetCollection("Urls")
	defer session.Close()
	err := collection.Insert(&Url{
		Id:  bson.NewObjectId(),
		Url: url,
	})
	if err != nil {
		log.Printf("Url insert error:%s\n", err.Error())
		return
	}
	log.Printf("Url:%s Insert Success!\n", url)
}

func InsertResponse(url, response string) {
	session, collection := GetCollection("Response")
	defer session.Close()
	err := collection.Insert(&Response{
		Url:     url,
		Context: response,
	})
	if err != nil {
		log.Println("Insert response failed!")
		return
	}
	log.Println("Insert Response Success!")
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

	//设置主键唯一
	keyindx := mgo.Index{
		Key:    []string{"key"},
		Unique: true,
	}
	//获取节点信息
	session, collection := GetCollection("nodes")
	defer session.Close()
	collection.EnsureIndex(keyindx)
}
