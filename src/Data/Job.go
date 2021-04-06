/*
 * @Description: 任务结构体
 * @Author: Rocky Hoo
 * @Date: 2021-04-06 12:11:31
 * @LastEditTime: 2021-04-06 12:13:19
 * @LastEditors: Please set LastEditors
 * @CopyRight: 
 * Copyright (c) 2021 XiaoPeng Studio
 */
package Data
type Job struct {
	Type string //标识任务类型
	Url  string //待爬取的目的url
}