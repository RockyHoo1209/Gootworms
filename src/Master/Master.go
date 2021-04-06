/*
 * @Description: Master节点的编写
 * @Author: Rocky Hoo
 * @Date: 2021-04-05 00:01:58
 * @LastEditTime: 2021-04-06 12:34:03
 * @LastEditors: Please set LastEditors
 * @CopyRight:
 * Copyright (c) 2021 XiaoPeng Studio
 */
package Master

import (
	"log"
	"main/src/Data"
	"main/src/Slave/Node"
	"main/src/Utils/DBUtils/RedisUtil"
	"time"
)

type Master struct {
	url_chan    chan interface{}
	idle_chan   chan *Node.Node
	result_chan chan interface{}
	redis       *RedisUtil.Redis
}

// var url_chan = make(chan string, 100)
// var idle_chan = make(chan *Node.Node, 100)
// var result_chan = make(chan string, 100)
// var redis RedisUtil.Redis

func InitMaster() (m *Master) {
	redis, err := RedisUtil.InitRedis()
	if err != nil {
		//Todo:宕机
		log.Println("Master-GetUrl:" + err.Error())
		return nil
	}
	m = &Master{
		url_chan:    make(chan interface{}, 100),
		idle_chan:   make(chan *Node.Node, 100),
		result_chan: make(chan interface{}, 100),
		redis:       redis,
	}
	return
}

func (m *Master) GetUrl() {
	url, err := m.redis.LPop("url")
	if err != nil {
		log.Println("Master-GetUrl:" + err.Error())
		return
	}
	m.url_chan <- url
}

/**
 * @description: 将目标url发送给节点机器
 * @param  {*}
 * @return {*}
 */
func (m *Master) CrawlByUrl() {
	for {
		select {
		case <-time.After(time.Duration(1)):
			url := <-m.url_chan
			// Todo:替换成加减器?
			<-m.idle_chan
			m.redis.RPush("Jobs", &Data.Job{
				Type: "Crawl",
				Url:  url.(string),
			})
		}
	}
}

/**
 * @description:从redis中获取对应任务返回的结果
 * @param  {*}
 * @return {*}
 */
func (m *Master) GetResult() {
	for {
		select {
		case <-time.After(time.Duration(1)):
			result, err := m.redis.LPop("result")
			if err != nil {
				log.Println("Master-GetResult:", err.Error())
				return
			}
			m.result_chan <- result
		}
	}
}

/**
 * @description:对收到result的后续操作
 * @param  {*}
 * @return {*}
 */
func (m *Master) HandleResult() {
	for {
		select {
		case <-time.After(time.Duration(1)):
			// Todo:拿到返回的result进行后续操作
			<-m.result_chan
		}
	}
}

/**
 * @description:每两秒获取一次Url
 * @param  {*}
 * @return {*}
 */
func (m *Master) LoadUrlsFromRedis() {
	for {
		select {
		case <-time.After(time.Duration(2)):
			go m.GetUrl()
		}
	}
}

/**
 * @description: 运行redis
 * @param  {*}
 * @return {*}
 */
func (m *Master) RunMaster() {
	go m.LoadUrlsFromRedis()
	go m.CrawlByUrl()
	go m.GetResult()
}
