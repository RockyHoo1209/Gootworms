/*
 * @Description:页面下载类,负责将页面中有的urls和itmes解析出来
 * @Author: Rocky Hoo
 * @Date: 2021-03-08 12:14:33
 * @LastEditTime: 2021-03-26 08:33:25
 * @LastEditors: Please set LastEditors
 * @CopyRight: Copyright (c) 2021 XiaoPeng Studio
 */

package Downloader

import (
	"log"
	"main/src/Utils/ChromeDriverUtil"
	"main/src/Utils/DBUtils/RedisUtil"
)

/**
 * @description:获得请求url后的请求数据,并把数据发送到redis
 * @param  {*}
 * @return {*}
 * @param {string} url
 */
func Download(url string) (string, error) {
	ChromeDriverUtil.InitChrome()
	resp, err := ChromeDriverUtil.ParseUrl(url)
	if err != nil {
		log.Println(err)
		return "", err
	}
	//TODO:改成进程间通信实现方式
	r := RedisUtil.NewClient()
	r.RPush("response", resp)
	return resp, nil
}
