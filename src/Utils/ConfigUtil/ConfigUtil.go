/*
 * @Description:配置文件工具类
 * @Author: Rocky Hoo
 * @Date: 2021-03-22 20:46:18
 * @LastEditTime: 2021-03-22 21:43:46
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

type Config struct {
	fileName string
}

/**
 * @description:初始化viper
 * @param  {*}
 * @return {*}
 */
func (c *Config) init() error {
	if c.fileName == "" {
		viper.SetConfigName("config")
		viper.AddConfigPath("./conf")
	} else {
		viper.SetConfigName(c.fileName)
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
	return nil
}

/**
 * @description:热加载配置文件
 * @param  {*}
 * @return {*}
 */
func (c *Config) watchConfig() {
	viper.WatchConfig()
	viper.OnConfigChange(func(e fsnotify.Event) {
		log.Printf("Config file:%s changed!\n", e.Name)
	})
}

func InitConfig(fileName string) error {
	c := &Config{
		fileName: fileName,
	}
	if err := c.init(); err != nil {
		return err
	}
	c.watchConfig()
	return nil
}
