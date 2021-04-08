/*
 * @Description:
 * @Author: Rocky Hoo
 * @Date: 2021-03-25 10:39:57
 * @LastEditTime: 2021-04-07 09:42:35
 * @LastEditors: Please set LastEditors
 * @CopyRight:
 * Copyright (c) 2021 XiaoPeng Studio
 */
package Spider

import (
	// "main/src/Utils/ChromeDriverUtil"
	"main/src/Utils/DBUtils/MongoUtil"
	"testing"
)

func TestSpider(*testing.T) {
	MongoUtil.InitMongo()
	s:=InitSpider()
	s.Wg.Add(1)
	go s.RunSpider()
	s.Wg.Wait()
}
