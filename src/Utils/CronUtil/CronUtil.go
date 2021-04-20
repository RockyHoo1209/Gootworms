/*
 * @Description: 定时器工具类
 * @Author: Rocky Hoo
 * @Date: 2021-04-19 23:25:24
 * @LastEditTime: 2021-04-20 09:03:03
 * @LastEditors: Please set LastEditors
 * @CopyRight:
 * Copyright (c) 2021 XiaoPeng Studio
 */
package CronUtil

import (
	"log"

	"github.com/robfig/cron/v3"
)

/**
 * @description: 定时每天运行
 * @param  {*}
 * @return {*}
 */
func RunDaily(task func()) error {
	c := cron.New()
	_, err := c.AddFunc("@daily", task)
	if err != nil {
		log.Println("CronUtil-RunDaily:" + err.Error())
		return err
	}
	c.Start()
	return nil
}

/**
 * @description: 定时每周运行
 * @param  {*}
 * @return {*}
 */
func RunWeekly(task func()) error {
	c := cron.New()
	_, err := c.AddFunc("@weekly", task)
	if err != nil {
		log.Println("CronUtil-RunWeekly:" + err.Error())
		return err
	}
	c.Start()
	return nil
}

/**
 * @description:自定义周期运行
 * @param  {interval取值范围:time.ParseDuration()支持的格式}
 * @return {*}
 */
func RunInterval(task func(), interval string) error {
	c := cron.New()
	_, err := c.AddFunc("@every "+interval, task)
	if err != nil {
		log.Println("CronUtil-RunInterval:" + err.Error())
		return err
	}
	c.Start()
	return nil
}
