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
	"os"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Printf("input helper to see the validate args!\n")
		return
	}
	switch os.Args[1] {
	case "master":
		log.Printf("starting master...\n")
		Master.InitMaster().RunMaster()
	case "worker":
		log.Printf("stsarting worker...\n")
		s:=Spider.InitSpider()
		if s!=nil{
			s.Wg.Add(1)
			go s.RunSpider()
			s.Wg.Wait()
		}
	}
}
