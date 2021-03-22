/*
 * @Description:页面下载类,负责将页面中有的urls和itmes解析出来
 * @Author: Rocky Hoo
 * @Date: 2021-03-08 12:14:33
 * @LastEditTime: 2021-03-11 08:54:44
 * @LastEditors: Please set LastEditors
 * @CopyRight: Copyright (c) 2021 XiaoPeng Studio
 */

package Downloader

import (
	"fmt"
	"main/src/Utils/LinksUtil"
	"regexp"
)

/**
 * @description:返回页面爬取到的链接列表并根据正则表达式提取相应的item数据
 * @param  {*}
 * @return {*}
 * @param {string} url
 */
func Crawl(url string, regex string) ([]string, []string, error) {
	links, html, err := LinksUtil.Extract(url)
	if err != nil {
		return nil, nil, fmt.Errorf(err.Error())
	}
	if regex == "" {
		return links, nil, nil
	}
	reg, err := regexp.Compile(regex)
	if err != nil {
		return links, nil, fmt.Errorf(err.Error())
	}
	items := reg.FindAllString(html, -1)
	return links, items, nil
}
