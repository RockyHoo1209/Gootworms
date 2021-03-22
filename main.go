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
	"main/src/Utils/LinksUtil"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Printf("input helper to see the validate args!\n")
		return
	}
	LinksUtil.Extract("https://news.qq.com/")
	switch os.Args[1] {
	case "master":
		log.Printf("starting master...\n")
	case "worker":
		log.Printf("stsarting worker...\n")
	}
}
