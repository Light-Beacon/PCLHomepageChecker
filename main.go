package main

import (
	"fmt"
	"io"
	"net/http"
	"strings"
)

var client = &http.Client{}

func getrequest(url string, headers map[string]string) string {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		panic(err)
	}

	for header, value := range headers {
		req.Header.Set(header, value)
	}

	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	if resp.StatusCode >= 400 {
		panic(fmt.Sprintf("返回 HTTP %d", resp.StatusCode))
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	return string(body)
}

func unsafe_check(url string, withua bool, withref bool) {
	headers := map[string]string{}
	if withref {
		headers["Referer"] = "http://9999.pcl2.server/"
	}
	if withua {
		headers["User-Agent"] = "PCL2/9999"
	}
	resp := getrequest(url, headers)
	if strings.HasPrefix(resp, "<!doctype html>") {
		panic("返回 HTML")
	}
}

func check(url string, withua bool, withref bool, successstr string, failedstr string) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println(failedstr, r)
		}
	}()
	unsafe_check(url, withua, withref)
	fmt.Println(successstr)
}

func main() {
	//url := "http://news.bugjump.net"
	url := "http://pcl.mcnews.thestack.top"
	fmt.Println("检查主页是否可用")
	check(url, true, true, "正常", "异常")
	fmt.Println("检查是否设置 UA 过滤")
	check(url, true, false, "未设置", "已设置:")
	fmt.Println("检查是否设置 Referer 过滤")
	check(url, false, true, "未设置", "已设置:")
	fmt.Println("检查是否设置版本号")
	check(url+"/version", true, true, "正常", "异常")
}
