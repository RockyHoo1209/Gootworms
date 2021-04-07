/*
 * @Description:心跳模块,用来向mater节点发送该节点的工作状态
 * @Author: Rocky Hoo
 * @Date: 2021-03-26 08:10:33
 * @LastEditTime: 2021-04-05 17:31:46
 * @LastEditors: Please set LastEditors
 * @CopyRight:
 * Copyright (c) 2021 XiaoPeng Studio
 */
package Node

import (
	"log"
	"main/src/Data/Enum/NodeConstants"
	"main/src/Utils/DBUtils/RedisUtil"
	"net"
)

type Node struct {
	Id     string //用mac地址代替节点id
	Status uint   //标识当前工作状态
	NetInterface net.Interface
}

/**
 * @description: 发送心跳
 * @param  {*}
 * @return {*}
 */
func (node *Node) HeartBeat() error {
	r, err := RedisUtil.InitRedis()
	if err != nil {
		log.Println(err.Error())
		return err
	}
	err = r.RPush("nodes", node)
	if err != nil {
		return err
	}
	return nil
}

func (node *Node) SetStatus(status uint) {
	node.Status = status
}

func InitNode() (*Node, error) {
	netInterfaces, err := net.Interfaces()
	if err != nil {
		return nil, err
	}
	return &Node{
		// Todo:Id 改成mac地址或用uuid实现
		Id:     "1313",
		NetInterface :netInterfaces[0],
		Status :NodeConstants.Idle,
	}, nil
}
