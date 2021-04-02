/*
 * @Description:
 * @Author: Rocky Hoo
 * @Date: 2021-04-02 11:49:46
 * @LastEditTime: 2021-04-02 12:46:05
 * @LastEditors: Please set LastEditors
 * @CopyRight:
 * Copyright (c) 2021 XiaoPeng Studio
 */
package TaskUtil

import (
	"fmt"
	"testing")



func TestTaskUtil(*testing.T){
	taskUtil:=NewTaskUtil()
	f:=func(test ...interface{}){
		fmt.Println(test)
	}
	taskUtil.Append(f)
	fmt.Println(len(taskUtil.tasks))
	for _,task :=range(taskUtil.tasks){
		task("hello")
		taskUtil.Pop(len(taskUtil.tasks)-1)
	}
	fmt.Println(len(taskUtil.tasks))
}