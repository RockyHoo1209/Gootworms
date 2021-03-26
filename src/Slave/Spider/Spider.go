/*
 * @Description: 负责解析页面的主要工作
 * @Author: Rocky Hoo
 * @Date: 2021-03-24 14:51:53
 * @LastEditTime: 2021-03-25 12:01:09
 * @LastEditors: Please set LastEditors
 * @CopyRight:
 * Copyright (c) 2021 XiaoPeng Studio
 */
package Spider

import (
	"main/src/Utils/DBUtils/MongoUtil"
	"main/src/Utils/LinksUtil"
)

func Spider(url, resp string) ([]string, []string, error) {
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
