package main

import (
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
	"time"
)

var client = &http.Client{
	Timeout: time.Second * 10}
var goterror = false

func GetRequest(url string, headers map[string]string, fail func(string)) (string, error) {
	var err error
	err = nil
	defer func() {
		if pan := recover(); pan != nil {
			errmsg := "在执行请求时出了意外错误: " + fmt.Sprint(pan)
			fail(errmsg)
			err = errors.New(errmsg)
		}
	}()
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		fail(err.Error())
		return "", err
	}
	for header, value := range headers {
		req.Header.Set(header, value)
	}
	resp, err := client.Do(req)
	defer resp.Body.Close()
	if err != nil {
		fail(err.Error())
		return "", err
	}
	if resp.StatusCode >= 400 {
		fail(fmt.Sprintf("返回 HTTP %d", resp.StatusCode))
		return "", fmt.Errorf("返回 HTTP %d", resp.StatusCode)
	}
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fail(err.Error())
		return "", err
	}
	return string(body), err
}

func Check(url string, withua bool, ref string, maxlength int,
	success func(string), fail func(string)) {
	headers := map[string]string{}
	headers["Referer"] = ref
	if withua {
		headers["User-Agent"] = "PCL2/2.99.99.99 Mozilla/5.0 AppleWebKit/537.36 Chrome/63.0.3239.132 Safari/537.36"
	}
	resp, err := GetRequest(url, headers, fail)
	if err != nil {
		return
	}
	if maxlength > 0 && len(resp) > maxlength {
		fail(fmt.Sprintf("返回内容过大（%d字符）", len(resp)))
		return
	}
	if strings.HasPrefix(strings.ToLower(
		strings.ReplaceAll(resp, "\n", "")),
		"<!doctype html>") {
		fail("返回 HTML")
		return
	}
	success("")
}

func ClearConsoleStyle() {
	fmt.Print("\033[0m")
}

func CreateHandler(status_text string) func(string) {
	return func(message string) {
		print_message := status_text
		if len(message) > 0 {
			print_message += ": " + message
		}
		//print_message += "\033[30m"
		fmt.Println(print_message)
		ClearConsoleStyle()
	}
}

const default_referer = "http://999.pcl2.server/"

func main() {
	var url string
	if len(os.Args) <= 1 {
		fmt.Println("提供主页 URL:")
		fmt.Scanln(&url)
	} else {
		url = os.Args[1]
	}
	success_handler := CreateHandler("\033[32m[✓] 正常")
	fail_handler := CreateHandler("\033[31m[X] 异常")
	fatalFailHandler := func(msg string) {
		fail_handler(msg)
		os.Exit(1)
	}
	fmt.Println("检查主页是否可用")
	Check(url, true, default_referer, 0, success_handler, fatalFailHandler)
	time.Sleep(time.Millisecond * 500)
	fmt.Println("检查开源版本 PCL 是否可用 [#5523前]")
	Check(url, true, "http://999.pcl2.open.server/", 0, success_handler, fail_handler)
	time.Sleep(time.Millisecond * 500)
	fmt.Println("检查开源版本 PCL 是否可用 [#5523后]")
	Check(url, true, "http://999.open.pcl2.server/", 0, success_handler, fail_handler)
	time.Sleep(time.Millisecond * 500)
	if strings.HasSuffix(url, ".xaml") {
		fmt.Println("检查是否设置版本号(.ini)")
		Check(url+".ini", true, default_referer, 128, success_handler, fail_handler)
	} else {
		fmt.Println("检查是否设置版本号(version)")
		Check(url+"/version", true, default_referer, 128, success_handler, fail_handler)
	}
	time.Sleep(time.Millisecond * 500)
	success_handler = CreateHandler("\033[32m[✓] 已设置")
	fail_handler = CreateHandler("\033[31m[X] 未设置")
	fmt.Println("检查是否设置 UA 过滤")
	Check(url, false, default_referer, 0, fail_handler, success_handler)
	time.Sleep(time.Millisecond * 500)
	fmt.Println("检查是否设置 Referer 过滤")
	Check(url, true, "", 0, fail_handler, success_handler)
	fmt.Println("\n按回车键退出")
	fmt.Scanln()
}
