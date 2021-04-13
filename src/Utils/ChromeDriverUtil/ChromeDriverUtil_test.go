package ChromeDriverUtil

import (
	"fmt"
	"testing"
)

func TestChromeDriverUtil(*testing.T) {
	targetUrl := "https://cn.bing.com/"
	InitChrome()
	ret, err := ParseUrl(targetUrl)
	if err != nil {
		fmt.Print(err)
	}
	fmt.Print(ret)
}
