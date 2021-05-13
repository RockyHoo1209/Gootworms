/*
 * @Description:
 * @Author: Rocky Hoo
 * @Date: 2021-03-22 21:19:34
 * @LastEditTime: 2021-05-13 16:52:32
 * @LastEditors: Please set LastEditors
 * @CopyRight:
 * Copyright (c) 2021 XiaoPeng Studio
 */
package ConfigUtil

import (
	"fmt"
	"testing"
)

func TestInitConfig(*testing.T) {
	InitConfig("config")
	address := GetString("redis.address")
	var port = GetString("redis.port")
	temp:=GetStringMap("client.rule.item_rule")
	url_rules:=GetStringSlice("client.rule.url_rule")
	fmt.Println(temp)
	fmt.Print(address)
	fmt.Println(port)
	for _,rule:=range(url_rules){
		fmt.Println(rule)
	}
}
