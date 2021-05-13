/*
 * @Description:配置文件工具类
 * @Author: Rocky Hoo
 * @Date: 2021-03-22 20:46:18
 * @LastEditTime: 2021-05-13 16:49:02
 * @LastEditors: Please set LastEditors
 * @CopyRight:
 * Copyright (c) 2021 XiaoPeng Studio
 */
package ConfigUtil

import (
	"fmt"
	"log"

	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

// type Config struct {
// 	fileName string
// }

/**
 * @description:初始化viper
 * @param  {*}
 * @return {*}
 */
func InitConfig(fileName string) error {
	if fileName == "" {
		viper.SetConfigName("config")
		viper.AddConfigPath("./conf")
	} else {
		viper.SetConfigName(fileName)
		viper.AddConfigPath("./")
	}
	viper.SetConfigType("yml")
	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			return fmt.Errorf("ConfigFileNotFoundError!")
		} else {
			return fmt.Errorf("Failed to read config file!")
		}
	}
	WatchConfig()
	return nil
}

/**
 * @description:热加载配置文件
 * @param  {*}
 * @return {*}
 */
func WatchConfig() {
	viper.WatchConfig()
	viper.OnConfigChange(func(e fsnotify.Event) {
		log.Printf("Config file:%s changed!\n", e.Name)
	})
}

func GetString(field_name string) string {
	return viper.GetString(field_name)
}

func GetStringMap(field_name string) map[string]interface{}{
	return viper.GetStringMap(field_name)
}

func GetStringSlice(field_name string) []string{
	return viper.GetStringSlice(field_name)
}