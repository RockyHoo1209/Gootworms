/*
 * @Description: 任务管理处理管理模块
 * @Author: Rocky Hoo
 * @Date: 2021-04-02 11:17:24
 * @LastEditTime: 2021-04-02 12:45:16
 * @LastEditors: Please set LastEditors
 * @CopyRight:
 * Copyright (c) 2021 XiaoPeng Studio
 */
package TaskUtil


type TaskUtil struct{
	//待处理函数队列
	tasks []func(vals ...interface{})
}

/**
 * @description: 构造函数
 * @param  {*}
 * @return {*}
 */
func NewTaskUtil() *TaskUtil{
	return &TaskUtil{
		tasks: make([]func(args ...interface{}), 0),
	}
}

/**
 * @description:将任务添加到任务队列 
 * @param  {*}
 * @return {*}
 * @param {...interface{}} args	
 */
func(task *TaskUtil) Append(function func(args ...interface{})){
	task.tasks=append(task.tasks,function)
}

/**
 * @description:执行完任务后出队 
 * @param  {*}
 * @return {*}
 */
func(task *TaskUtil) Pop(index int){
	// 使用 … 运算符，可以将一个切片的所有元素追加到另一个切片里
	task.tasks=append(task.tasks[:index],task.tasks[index+1:]...)
}
