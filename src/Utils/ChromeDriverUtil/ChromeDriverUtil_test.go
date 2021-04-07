package ChromeDriverUtil

import (
	"fmt"
	"testing"
)

func TestChromeDriverUtil(*testing.T) {
	targetUrl := "https://www.bilibili.com/"
	ret, err := StartChrome(targetUrl)
	if err != nil {
		fmt.Print(err)
	}
	fmt.Print(ret)
}
