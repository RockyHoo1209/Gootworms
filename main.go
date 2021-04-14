/*
 * Filename: f:\work\DistributeScrawer\main.go
 * Path: f:\work\DistributeScrawer
 * Created Date: Monday, March 8th 2021, 10:41:44 am
 * Author: Rocky Hoo
 *
 * Copyright (c) 2021 XiaoPeng Studio
 */

package main

import (
	"fmt"
	"log"
	"main/src/Master"
	"main/src/Slave/Spider"
	"main/src/Utils/DBUtils/MongoUtil"
	"os"
	"runtime"
)

var runType string


/**
 * @description: 应用重启
 * @param  {*}
 * @return {*}
 */
func reStart() {
	switch runType {
	case "worker":
		runWorker()
	case "master":
		runMaster()
	default:
		runWorker()
	}
}

/**
 * @description: 程序崩溃后重新运行
 * @param  {*}
 * @return {*}
 */
func protectRun() {
	defer func() {
		if err := recover(); err != nil {
			switch err.(type) {
			case runtime.Error:
				reStart()
				log.Println("runtime error!")
			default:
				reStart()
				log.Println("unknown error:", err)
			}
		}
	}()
}

/**
 * @description: 节点机器启动运行
 * @param  {*}
 * @return {*}
 */
func runWorker() {
	log.Printf("starting worker...\n")
	runType = "worker"
	s := Spider.InitSpider()
	if s != nil {
		s.Wg.Add(1)
		go s.RunSpider()
		s.Wg.Wait()
	}
}

/**
 * @description: 控制节点运行
 * @param  {*}
 * @return {*}
 */
func runMaster() {
	log.Printf("starting master...\n")
	runType = "master"
	MongoUtil.InitMongo()
	Master.InitMaster().RunMaster()
}

func main() {
	go protectRun()
	if len(os.Args) < 2 {
		fmt.Printf("input helper to see the validate args!\n")
		return
	}
	switch os.Args[1] {
	case "master":
		runMaster()
	case "worker":
		runWorker()
	}
}
