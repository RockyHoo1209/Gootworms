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

	// "main/src/Slave/Node"
	"main/src/Utils/DBUtils/RedisUtil"
	"sync"
	"time"

	// "github.com/spf13/viper"
)

type Master struct {
	url_chan    chan interface{} //读取url同步通信
	// Todo:讨论用哪种方法实现更好?缓冲通道?
	// work_chan   chan *Node.Node  //读取空闲节点同步通信
	work_count   sync.WaitGroup  //记录工作数量同步通信
	result_chan chan interface{} //爬取结果同步通信
	redis       *RedisUtil.Redis //初始化redis
	wg          sync.WaitGroup
}

func InitMaster() (m *Master) {
	redis, err := RedisUtil.InitRedis()
	if err != nil {
		//Todo:宕机
		log.Println("Master-GetUrl:" + err.Error())
		return nil
	}
	m = &Master{
		url_chan:    make(chan interface{}),
		// work_chan:   make(chan *Node.Node),
		result_chan: make(chan interface{}),
		redis:       redis,
	}
	return
}

func (m *Master) GetUrl() {
	url, err := m.redis.LPop("url")
	if err != nil {
		log.Println("Master-GetUrl:" + err.Error())
		url="www.bilibili.com"
		m.url_chan<-url//viper.GetString("server.firstUrl")
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
	defer m.wg.Done()
	for {
		select {
		case <-time.After(time.Duration(1)):
			url := <-m.url_chan
			log.Println("sending to slaves:",url)
			// Todo:替换成加减器?
			m.work_count.Add(1)
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
	defer m.wg.Done()
	for {
		select {
		case <-time.After(time.Duration(1)):
			log.Println("Getting result...")
			result, err := m.redis.LPop("Result")
			if err != nil {
				log.Println("Master-GetResult:", err.Error())
				continue
			}
			m.result_chan <- result
			m.work_count.Done()
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
	defer m.wg.Done()
	for {
		select {
		case <-time.After(time.Duration(2)):
			log.Println("loading urls...")
			m.GetUrl()
			log.Println("got url")
		}
	}
}

/**
 * @description: 运行redis
 * @param  {*}
 * @return {*}
 */
func (m *Master) RunMaster() {
	m.wg.Add(1)
	go m.LoadUrlsFromRedis()
	m.wg.Add(1)
	go m.CrawlByUrl()
	m.wg.Add(1)
	go m.GetResult()
	// 阻塞直到waitgroup计数器减到0
	m.wg.Wait()
}
