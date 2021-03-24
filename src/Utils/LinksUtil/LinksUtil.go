/*
 * Filename: f:\work\DistributeScrawer\src\Utils\LinksUtil\LinksUtil.go
 * Path: f:\work\DistributeScrawer\src\Utils
 * Created Date: Wednesday, March 10th 2021, 8:26:20 pm
 * Author: Rocky Hoo
 *
 * Copyright (c) 2021 XiaoPeng Studio
 */

package LinksUtil

import (
	"log"
	"regexp"
	"strings"
)

func ExtractUrls(resp string) ([]string, error) {

	resp = strings.Replace(resp, " ", "", -1)
	resp = strings.Replace(resp, "\n", "", -1)
	//匹配所有链接形式
	reg, err := regexp.Compile("<ahref=\"([^\"]+)\"[^>]*>[^<]+</a>")
	if err != nil {
		log.Printf("reg parse error:%s\n", err.Error())
		return nil, err
	}
	links := reg.FindAllString(resp, -1)
	return links, nil
}

/**
 * @description:根据给定正则表达式爬取对应item元素
 * @param  {*}
 * @return {*}
 * @param {*} resp
 * @param {string} regex
 */
func ExtractItems(resp, regex string) ([]string, error) {

	resp = strings.Replace(resp, " ", "", -1)
	resp = strings.Replace(resp, "\n", "", -1)
	reg, err := regexp.Compile(regex)
	if err != nil {
		log.Printf("reg parse error:%s\n", err.Error())
		return nil, err
	}
	items := reg.FindAllString(resp, -1)
	return items, nil
}
