/*
 * @Description:
 * @Author: Rocky Hoo
 * @Date: 2021-03-22 21:19:34
 * @LastEditTime: 2021-03-22 21:45:25
 * @LastEditors: Please set LastEditors
 * @CopyRight:
 * Copyright (c) 2021 XiaoPeng Studio
 */
package ConfigUtil

import (
	"fmt"
	"testing"

	"github.com/spf13/viper"
)

func TestInitConfig(*testing.T) {
	InitConfig("config")
	address := viper.GetString("redis.address")
	var port = viper.GetString("redis.port")
	temp:=viper.GetStringMap("client.rule.item_rule")
	fmt.Println(temp)
	fmt.Print(address)
	fmt.Print(port)
}
