/*
 * @Description: 负责解析页面的主要工作
 * @Author: Rocky Hoo
 * @Date: 2021-03-24 14:51:53
 * @LastEditTime: 2021-04-06 12:41:16
 * @LastEditors: Please set LastEditors
 * @CopyRight:
 * Copyright (c) 2021 XiaoPeng Studio
 */
package Spider

import (
	"log"
	"main/src/Data"
	"main/src/Utils/ChromeDriverUtil"
	"main/src/Utils/DBUtils/MongoUtil"
	"main/src/Utils/DBUtils/RedisUtil"
	"main/src/Utils/LinksUtil"
	"time"
)

type Spider struct{
	redis *RedisUtil.Redis
}

func InitSpider()(* Spider){
	redis,err:=RedisUtil.InitRedis()
	if err!=nil{
		log.Println("Init Redis Failed:",err.Error())
		return nil
	}
	return &Spider{
		redis: redis,
	} 
}


/**
 * @description: 将爬取到对应的url、resp存入mongodb
 * @param  {*}
 * @return {*}
 * @param {*} url
 * @param {string} resp
 */
func (s *Spider)Crawl(url, resp string) ([]string, []string, error) {
	items, err := LinksUtil.ExtractItems(resp, "")
	if err != nil {
		return nil, nil, err
	}
	subUrls, err := LinksUtil.ExtractUrls(resp)
	if err != nil {
		return items, nil, err
	}
	MongoUtil.InsertResponse(url, resp)
	MongoUtil.InsertUrls(subUrls)
	return items, subUrls, nil
}

/**
 * @description:运行爬虫
 * @param  {*}
 * @return {*}
 */
func (s *Spider)RunSpider(){
	for{
		select {
		case <-time.After(time.Duration(1)):
			job,err:=s.redis.LPop("jobs")
			if err!=nil{
				log.Println("Spider-RunSpider:Lpop error!",err.Error())
				continue
			}
			task:=job.(Data.Job)
			resp,err:=ChromeDriverUtil.StartChrome(task.Url)
			if err!=nil{
				log.Println("Spider-RunSpider:ChromeDriver Error!")
				continue
			}
			go s.Crawl(task.Url,resp)			
		}
	}
}
