/*
 * @Description: master节点,负责负载均衡客户端节点
 * @Author: Rocky Hoo
 * @Date: 2021-03-28 08:16:59
 * @LastEditTime: 2021-03-28 09:02:48
 * @LastEditors: Please set LastEditors
 * @CopyRight:
 * Copyright (c) 2021 XiaoPeng Studio
 */
package Master

import (
	"log"
	"main/src/Enum/NodeConstants"
	"main/src/Slave/Node"
	"main/src/Utils/DBUtils/RedisUtil"
)


var Idle=make(map[string]*Node.Node)
var Working=make(map[string]*Node.Node)
var Task []interface{}

/**
 * @description:根据节点信息设置空闲队列 
 * @param  {*}
 * @return {*}
 * @param {*Node.Node} node
 */
func SetIdle(node *Node.Node){
	if _,ok :=Working[node.Id];!ok{
		log.Printf("Idle dosen't exist %s\n",node.Id)
		return
	}
	delete(Working,node.Id)
	Idle[node.Id]=node
}

/**
 * @description:根据节点信息设置工作队列
 * @param  {*}
 * @return {*}
 * @param {*Node.Node} node
 */
func SetWorking(node *Node.Node){
	if _,ok :=Working[node.Id];!ok{
		log.Printf("Working dosen't exist %s\n",node.Id)
		return
	}
	delete(Idle,node.Id)
	Working[node.Id]=node	
}

/**
 * @description:从redis中取出一个Node类型，同步其状态
 * @param  {*}
 * @return {*}
 * @param {*} Node
 */
func SetStatus(node *Node.Node){
	if node ==nil{
		return
	}
	if node.Status==NodeConstants.Idle{
		SetIdle(node)
		return
	}
	SetWorking(node)
}

/**
 * @description: 从空闲队列中取一个节点绑定任务
 * @param  {*}
 * @return {*}
 * @param {interface{}} job
 */
func ToWork(job interface{}) {
	if len(Idle) == 0 {
		log.Println("All Nodes are working!")
		return 
	}
	for _,node:=range Idle{
		if node==nil{
			continue
		}
		r,err:=RedisUtil.InitRedis()
		if err!=nil{
			log.Println("Redis init failed!")
			return
		}
		r.RPush("job",job)
		break
	}
}
