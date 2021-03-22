/*
 * @Description:http请求工具类
 * @Author: Rocky Hoo
 * @Date: 2021-03-11 08:15:34
 * @LastEditTime: 2021-03-11 12:09:48
 * @LastEditors: Please set LastEditors
 * @CopyRight: Copyright (c) 2021 XiaoPeng Studio
 */

package DoRequest

import (
	"crypto/tls"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"net/http/cookiejar"
	"net/url"
	"regexp"
	"strings"
)

/**
 * @description: 解析处理cookies
 * @param  {*}
 * @return {*}
 * @param {map[string]string} cookies
 * @param {*} path
 * @param {string} domain
 */
func CookiesHandler(cookies string, path, domain string) (http_cookies []*http.Cookie, err error) {
	if cookies == "" {
		return nil, fmt.Errorf("cookies is nil!")
	}
	if path == "" {
		path = "/"
	}

	var mapCookie = make(map[string]string)
	reg := regexp.MustCompile(`([^=]+)=([^;]*);?`)
	arrCookie := reg.FindAllStringSubmatch(cookies, -1)

	for _, cookie := range arrCookie {
		mapCookie[cookie[1]] = cookie[2]
	}

	for k, v := range mapCookie {
		cookie := &http.Cookie{
			Name:   k,
			Value:  v,
			Path:   path,
			Domain: domain,
		}
		http_cookies = append(http_cookies, cookie)
	}
	return
}

func AddHeaders(headers map[string]string, req *http.Request) {
	if headers == nil || req == nil {
		return
	}
	for key, val := range headers {
		req.Header.Add(key, val)
	}
}

/**
 * @description:将cookies转化成cookiejar
 * @param  {*}
 * @return {*}
 * @param {map[string]string} cookies
 * @param {*} domain
 * @param {string} requrl
 */
func AddCookies(cookies string, domain, requrl string) (*cookiejar.Jar, error) {

	http_cookies, err := CookiesHandler(cookies, "", domain)
	if err != nil {
		return nil, err
	}
	gCookieJar, _ := cookiejar.New(nil)
	cookieUrl, err := url.Parse(requrl)
	if err != nil {
		return nil, err
	}
	gCookieJar.SetCookies(cookieUrl, http_cookies)
	return gCookieJar, nil
}

/**
 * @description:模拟浏览器发起请求:1) complete request 2) prepare header if has 3) prepare cookie if has 4) prepare transport 5）construct req client
 * @param  {*}
 * @return {*}
 * @param {string} url
 */
func DoRequest(requrl, method, param, domain, cookies string, headers map[string]string) (respBody string, respCookies map[string]string, err error) {
	method = strings.ToUpper(method)
	//定义一个空接口
	var paramReader io.Reader = nil
	if method == "POST" {
		paramReader = strings.NewReader(param)
	}
	request, err := http.NewRequest(method, requrl, paramReader)
	if err != nil {
		return
	}
	AddHeaders(headers, request)
	gCookieJar, err := AddCookies(cookies, domain, requrl)
	if err != nil {
		log.Println(err)
		return
	}

	transport := &http.Transport{
		TLSClientConfig: &tls.Config{
			InsecureSkipVerify: true,
			MinVersion:         tls.VersionTLS10,
			MaxVersion:         tls.VersionTLS12,
		},
		//集合的压缩功能关闭
		DisableCompression: true,
	}
	client := &http.Client{
		Transport:     transport,
		Jar:           gCookieJar,
		CheckRedirect: nil,
	}
	response, err := client.Do(request)
	if err != nil {
		return
	}
	defer response.Body.Close()

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return
	}
	respBody = string(body)
	t_respCookies := response.Cookies()
	for _, cookie := range t_respCookies {
		respCookies[cookie.Name] = cookie.Value
	}
	return
}
