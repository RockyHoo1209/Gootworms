/*
 * Filename: f:\work\DistributeScrawer\src\Utils\LinksUtil\LinksUtil_test.go
 * Path: f:\work\DistributeScrawer\src\Utils
 * Created Date: Wednesday, March 10th 2021, 9:19:48 pm
 * Author: Rocky Hoo
 *
 * Copyright (c) 2021 XiaoPeng Studio
 */

package LinksUtil

import (
	"fmt"
	"log"
	"testing"
)

func TestExtract(t *testing.T) {
	links, html, err := Extract("https://www.runoob.com/redis/redis-commands.html")
	if err != nil {
		log.Fatalln(err)
	}
	for _, link := range links {
		fmt.Println("link:", link)
	}
	fmt.Println("html:" + html)
	fmt.Println("linksLen", len(links))
	fmt.Println("------------------------ok--------------------------")
}
