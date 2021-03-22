/*
 * @Description:simulate chrome browser to request a web
 * @Author: Rocky Hoo
 * @Date: 2021-03-21 17:38:25
 * @LastEditTime: 2021-03-21 17:39:13
 * @LastEditors: Please set LastEditors
 * @CopyRight:
 * Copyright (c) 2021 XiaoPeng Studio
 */
package ChromeDriverUtil

import (
	"fmt"

	"log"

	"github.com/tebeka/selenium"

	"github.com/tebeka/selenium/chrome"
)

// StartChrome 启动谷歌浏览器headless模式
func StartChrome(targetUrl string) (string, error) {

	opts := []selenium.ServiceOption{}

	caps := selenium.Capabilities{

		"browserName": "chrome",
	}

	// 禁止加载图片，加快渲染速度

	imagCaps := map[string]interface{}{

		"profile.managed_default_content_settings.images": 2,
	}

	chromeCaps := chrome.Capabilities{

		Prefs: imagCaps,

		Path: "",

		Args: []string{

			"--headless", // 设置Chrome无头模式

			"--no-sandbox",

			"--user-agent=Mozilla/5.0 (Macintosh; Intel Mac OS X 10_13_2) AppleWebKit/604.4.7 (KHTML, like Gecko) Version/11.0.2 Safari/604.4.7", // 模拟user-agent，防反爬

		},
	}

	caps.AddChrome(chromeCaps)

	// 启动chromedriver，端口号可自定义

	service, err := selenium.NewChromeDriverService("/bin/chromedriver", 9516, opts...)

	if err != nil {

		log.Printf("Error starting the ChromeDriver server: %v", err)
		return "", err

	}

	// 调起chrome浏览器
	//fmt.Sprintf("http://localhost:%d/wd/hub", 9516)
	webDriver, err := selenium.NewRemote(caps, fmt.Sprintf("http://localhost:%d/wd/hub", 9516))
	if err != nil {

		// panic(err)
		log.Fatal(err)
		return "", err

	}

	// 导航到目标网站
	fmt.Println("before get!")
	err = webDriver.Get(targetUrl)
	fmt.Println("after get!")

	if err != nil {

		log.Println(fmt.Sprintf("Failed to load page: %s\n", err))
		return "", err

	}
	src, err := webDriver.PageSource()
	if err != nil {
		log.Println(err)
		return "", err
	}
	fmt.Print("src:", src)
	log.Println(webDriver.GetCookies())

	defer service.Stop() // 停止chromedriver

	defer webDriver.Quit() // 关闭浏览器
	return src, nil
}
