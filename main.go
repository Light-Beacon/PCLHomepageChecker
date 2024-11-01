package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
)

var client = &http.Client{}
var goterror = false

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

func unsafe_check(url string, withua bool, ref string, maxlength int) {
	headers := map[string]string{}
	headers["Referer"] = ref
	if withua {
		headers["User-Agent"] = "PCL2/9999"
	}
	resp := getrequest(url, headers)
	if maxlength > 0 && len(resp) > maxlength {
		panic(fmt.Sprintf("返回内容过大（%d字符）", len(resp)))
	}
	if strings.HasPrefix(strings.ToLower(
		strings.ReplaceAll(resp, "\n", "")),
		"<!doctype html>") {
		panic("返回 HTML")
	}
}

func check(url string, withua bool, ref string,
	successstr string, failedstr string, maxlenght int) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println(failedstr, r, "\033[0m")
			goterror = true
		}
	}()
	unsafe_check(url, withua, ref, maxlenght)
	fmt.Println(successstr, "\033[0m")
}

func main() {
	var url string
	PCL_REF := "http://9999.pcl2.server/"
	if len(os.Args) <= 1 {
		fmt.Println("提供主页 URL:")
		fmt.Scanln(&url)
	} else {
		url = os.Args[1]
	}
	fmt.Println("检查主页是否可用")
	check(url, true, PCL_REF,
		"\033[32m[✓] 正常", "\033[31m[X] 异常：", -1)
	if goterror {
		return
	}
	fmt.Println("检查开源版本 PCL 是否可用")
	check(url, true, "http://9999.pcl2.open.server/",
		"\033[32m[✓] 正常", "\033[31m[X] 异常：", -1)
	if strings.HasSuffix(url, ".xaml") {
		fmt.Println("检查是否设置版本号(.ini)")
		check(url+".ini", true, PCL_REF,
			"\033[32m[✓] 正常", "\033[31m[X] 异常：", 128)
	} else {
		fmt.Println("检查是否设置版本号(version)")
		check(url+"/version", true, PCL_REF,
			"\033[32m[✓] 正常", "\033[31m[X] 异常：", 128)
	}
	fmt.Println("检查是否设置 UA 过滤")
	check(url, true, "",
		"\033[31m[X] 未设置", "\033[32m[✓] 已设置:", -1)
	fmt.Println("检查是否设置 Referer 过滤")
	check(url, false, PCL_REF,
		"\033[31m[X] 未设置", "\033[32m[✓] 已设置:", -1)
}
