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
	"encoding/json"
	"log"
	"regexp"
	"strings"
)

/**
 * @description:
 * @param  {*}
 * @return {*}
 * @param {*} resp
 * @param {string} regexp_url:url正则匹配模式
 */
func ExtractUrls(resp, regexp_url string) ([]string, error) {
	resp = strings.Replace(resp, " ", "", -1)
	resp = strings.Replace(resp, "\n", "", -1)
	resp = strings.Replace(resp, "\t", "", -1)
	//匹配所有链接形式
	// reg, err := regexp.Compile("<ahref=\"([^\"]+)\"[^>]*>[^<]+</a>")
	// reg, err := regexp.Compile("((http[s]{0,1})://)(([a-zA-Z0-9\\._-]+\\.[a-zA-Z]{2,6})|([0-9]{1,3}\\.[0-9]{1,3}\\.[0-9]{1,3}\\.[0-9]{1,3}))(:[0-9]{1,4})*(/[a-zA-Z0-9\\&%_\\./-~-]*)?")
	reg, err := regexp.Compile(regexp_url)
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
	resp = strings.Replace(resp, "\t", "", -1)
	reg, err := regexp.Compile(regex)
	if err != nil {
		log.Printf("reg parse error:%s\n", err.Error())
		return nil, err
	}
	items := reg.FindAllString(resp, -1)
	return items, nil
}

/**
 * @description:根据给定正则表达式爬取对应item元素
 * @param  {*}
 * @return {*}
 * @param {*} resp
 * @param {string} regex
 */
func ExtractItems2Map(url,resp string, regexes map[string]interface{}) (map[string]interface{}, error) {
	resp = strings.Replace(resp, " ", "", -1)
	resp = strings.Replace(resp, "\n", "", -1)
	resp = strings.Replace(resp, "\t", "", -1)
	mapResult:=map[string]interface{}{}
	for k, v := range regexes {
		reg, err := regexp.Compile(v.(string))
		if err != nil {
			log.Printf("reg parse error:%s\n", err.Error())
			return nil, err
		}
		items := reg.FindAllStringSubmatch(resp, -1)
		if items==nil{
			continue
		}
		mapResult["url"]=url
		if k=="base" {
			if err := json.Unmarshal([]byte(items[0][1]), &mapResult); err != nil {
				log.Fatal(err)
			}
			continue
		}
		field_content:=""
		for _,item:=range(items){
			for i,c:=range(item){
				if i==0{
					continue
				}
				field_content+=c
			}
		}
		mapResult[k]=field_content
	}
	return mapResult, nil
}
