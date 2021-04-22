/*
 * @Description: 负责解析页面的主要工作
 * @Author: Rocky Hoo
 * @Date: 2021-03-24 14:51:53
 * @LastEditTime: 2021-04-22 17:08:37
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
	"main/src/Utils/ConfigUtil"
	"main/src/Utils/CronUtil"
	"sync"

	"main/src/Utils/DBUtils/RedisUtil"
	"main/src/Utils/LinksUtil"
	// "time"
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
	ChromeDriverUtil.InitChrome()
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
func (s *Spider) Crawl(url, resp, regex string) {
	items, err := LinksUtil.ExtractItems(resp, regex)
	if err != nil {
		s.result_chan <- &Data.Result{
			Items:   nil,
			Suburls: nil,
			Resp:    resp,
		}
		return
	}
	regexp_url:=ConfigUtil.GetString("client.rule.url_rule")
	subUrls, err := LinksUtil.ExtractUrls(resp,regexp_url)
	if err != nil {
		s.result_chan <- &Data.Result{
			Items:   items,
			Suburls: subUrls,
			Resp:    resp,
		}
		return
	}

	s.result_chan <- &Data.Result{
		Items:   items,
		Suburls: subUrls,
		Resp:    resp,
	}
}

/**
 * @description:单次运行爬虫
 * @param  {*}
 * @return {*}
 */
func (s *Spider) RunSpider() {
	regex := ConfigUtil.GetString("client.rule.item_rule")
	job, err := s.redis.BLPop("Jobs", 0)
	if err != nil {
		log.Println("Spider-RunSpider:Lpop error!", err.Error())
		return
	}
	var task *Data.Job

	// json-string转结构体
	json.Unmarshal([]byte(job), &task)

	resp, err := ChromeDriverUtil.ParseUrl(task.Url)
	if err != nil {
		log.Println("Spider-RunSpider:ChromeDriver Error!")
		return
	}
	go s.Crawl(task.Url, resp, regex)
	result := <-s.result_chan
	jsonbyte, err := json.Marshal(&result)
	if err != nil {
		log.Println("Spider-RunSpider:", err.Error())
	}
	s.redis.RPush("Result", jsonbyte)
}

/**
 * @description: 爬虫周期性运行
 * @param  {*}
 * @return {*}
 */
func (s *Spider) RunInterval(Interval string) error {
	defer s.Wg.Done()
	if Interval=="daily"{
		err:=CronUtil.RunDaily(s.RunSpider)
		return err
	}
	if Interval=="weekly"{
		err:=CronUtil.RunWeekly(s.RunSpider)
		return err
	}
	err:=CronUtil.RunInterval(s.RunSpider,Interval)
	return err
}
