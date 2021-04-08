/*
 * @Description: 负责解析页面的主要工作
 * @Author: Rocky Hoo
 * @Date: 2021-03-24 14:51:53
 * @LastEditTime: 2021-04-07 09:41:46
 * @LastEditors: Please set LastEditors
 * @CopyRight:
 * Copyright (c) 2021 XiaoPeng Studio
 */
package Spider

import (
	"encoding/json"
	"log"
	"main/src/Data"
	"main/src/Utils/ChromeDriverUtil"
	"sync"

	"main/src/Utils/DBUtils/RedisUtil"
	"main/src/Utils/LinksUtil"
	"time"
)

type Spider struct {
	redis       *RedisUtil.Redis
	result_chan chan *Data.Result
	Wg          sync.WaitGroup
}

func InitSpider() *Spider {
	redis, err := RedisUtil.InitRedis()
	if err != nil {
		log.Println("Init Redis Failed:", err.Error())
		return nil
	}
	return &Spider{
		redis:       redis,
		result_chan: make(chan *Data.Result),
	}
}

/**
 * @description: 将爬取到对应的url、resp存入mongodb
 * @param  {*}
 * @return {*}
 * @param {*} url
 * @param {string} resp
 */
func (s *Spider) Crawl(url, resp string) {
	items, err := LinksUtil.ExtractItems(resp, "")
	if err != nil {
		s.result_chan <- &Data.Result{
			Items:   nil,
			Suburls: nil,
			Resp: resp,
		}
		return
	}
	subUrls, err := LinksUtil.ExtractUrls(resp)
	if err != nil {
		s.result_chan <- &Data.Result{
			Items:   items,
			Suburls: subUrls,
			Resp: resp,
		}
		return
	}

	s.result_chan <- &Data.Result{
		Items:   items,
		Suburls: subUrls,
		Resp: resp,
	}
}

/**
 * @description:运行爬虫
 * @param  {*}
 * @return {*}
 */
func (s *Spider) RunSpider() {
	defer s.Wg.Done()
	for {
		select {
		case <-time.After(time.Duration(1)):
			job, err := s.redis.LPop("Jobs")
			if err != nil {
				log.Println("Spider-RunSpider:Lpop error!", err.Error())
				continue
			}
			var task *Data.Job
			json.Unmarshal([]byte(job.(string)),&task)
			resp, err := ChromeDriverUtil.StartChrome(task.Url)
			if err != nil {
				log.Println("Spider-RunSpider:ChromeDriver Error!")
				continue
			}
			go s.Crawl(task.Url, resp)
			result := <-s.result_chan
			jsonbyte,err:=json.Marshal(&result)
			if err!=nil{
				log.Println("Spider-RunSpider:",err.Error())
			}
			s.redis.RPush("Result",jsonbyte)
		}
	}
}
