/*
 * @Description:
 * @Author: Rocky Hoo
 * @Date: 2021-03-11 11:24:31
 * @LastEditTime: 2021-03-11 12:04:06
 * @LastEditors: Please set LastEditors
 * @CopyRight:
 * Copyright (c) 2021 XiaoPeng Studio
 */

package DoRequest

import (
	"fmt"
	"log"
	"testing"
)

func TestDoRequest(*testing.T) {
	cookies := "ptui_loginuin=1284219022@qq.com; RK=h5qZC+4rUu; ptcz=8540c130ba12db31a79f87664f9b1aabf3d2278a0d2b123e6491f78877a02be2; pgv_info=ssid=s8660672755; pgv_pvid=6260282816; pac_uid=0_038c89ad2f2a1; o_minduid=; appuser=30975E1F9E19CD5E; cm_cookie=V1,110279&x5zJl06d7M10&AQEB64RrNRvXhCfT_Pa7HnmsmIstxez8QeeV&210310&210310; _qpsvr_localtk=0.07819160708297379; lv_play_index=71"
	var headers = make(map[string]string)
	headers["Accept"] = "text/html,application/xhtml+xml,application/xml;q=0.9,image/webp,image/apng,*/*;q=0.8"
	headers["Accept-Language"] = "zh-CN,zh;q=0.8,en-US;q=0.5,en;q=0.3"
	headers["Connection"] = "keep-alive"
	headers["User-Agent"] = "Mozilla/5.0 (Windows NT 6.1; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.106 Safari/537.36"

	respBody, respCookies, err := DoRequest("https://news.qq.com/", "Get", "", "", cookies, headers)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("body:", respBody)
	fmt.Println("cookies:", respCookies)
}
