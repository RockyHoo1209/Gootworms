/*
 * @Description:simulate chrome browser to request a web
 * @Author: Rocky Hoo
 * @Date: 2021-03-21 17:38:25
 * @LastEditTime: 2021-05-13 19:19:16
 * @LastEditors: Please set LastEditors
 * @CopyRight:
 * Copyright (c) 2021 XiaoPeng Studio
 */
package ChromeDriverUtil

import (
	"fmt"
	"main/src/Utils/ConfigUtil"

	"log"

	"github.com/tebeka/selenium"

	"github.com/tebeka/selenium/chrome"
)

type Service struct {
	opts       []selenium.ServiceOption
	caps       selenium.Capabilities
	imagCaps   map[string]interface{}
	chromeCaps chrome.Capabilities
}

var chrome_service *Service
// InitChrome 初始化谷歌浏览器
func InitChrome(){
	if chrome_service!=nil{
		log.Println("ChromeDriverUtil-InitChrome:Chrome is already inited!")
		return
	}
	chrome_service = &Service{
		opts: []selenium.ServiceOption{},
		caps: selenium.Capabilities{
			"browserName": "chrome",
		},
		imagCaps: map[string]interface{}{
			"profile.managed_default_content_settings.images": 2,
		},
		chromeCaps: chrome.Capabilities{
			Prefs: map[string]interface{}{
				"profile.managed_default_content_settings.images": 2,
			},
			Path: "",

			Args: []string{

				"--headless", // 设置Chrome无头模式

				"--no-sandbox", //root模式下也可以运行

				"--user-agent=Mozilla/5.0 (Macintosh; Intel Mac OS X 10_13_2) AppleWebKit/604.4.7 (KHTML, like Gecko) Version/11.0.2 Safari/604.4.7", // 模拟user-agent，防反爬

				"--disable-dev-shm-usage",
			},
		},
	}
}


//解析url
func ParseUrl(targetUrl string)(string,error){
	chrome_service.caps.AddChrome(chrome_service.chromeCaps)

	// 启动chromedriver，端口号可自定义

	service, err := selenium.NewChromeDriverService(ConfigUtil.GetString("chromedriver.path"), 9516, chrome_service.opts...)

	if err != nil {

		log.Printf("Error starting the ChromeDriver server: %v", err)
		return "", err

	}

	// 调起chrome浏览器
	//fmt.Sprintf("http://localhost:%d/wd/hub", 9516)
	webDriver, err := selenium.NewRemote(chrome_service.caps, fmt.Sprintf("http://localhost:%d/wd/hub", 9516))
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
