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

func unsafe_check(url string, withua bool, withref bool, maxlength int) {
	headers := map[string]string{}
	if withref {
		headers["Referer"] = "http://9999.pcl2.server/"
	}
	if withua {
		headers["User-Agent"] = "PCL2/9999"
	}
	resp := getrequest(url, headers)
	if maxlength > 0 && len(resp) > maxlength {
		panic(fmt.Sprintf("返回内容过大（%d字符）", len(resp)))
	}
	if strings.HasPrefix(resp, "<!doctype html>") {
		panic("返回 HTML")
	}
}

func check(url string, withua bool, withref bool,
	successstr string, failedstr string, maxlenght int) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println(failedstr, r)
			goterror = true
		}
	}()
	unsafe_check(url, withua, withref, maxlenght)
	fmt.Println(successstr)
}

func main() {
	var url string
	if len(os.Args) <= 1 {
		fmt.Println("提供主页 URL:")
		fmt.Scanln(&url)
	} else {
		url = os.Args[1]
	}
	fmt.Println("检查主页是否可用")
	check(url, true, true,
		"[✓] 正常", "[X] 异常：", -1)
	if goterror {
		return
	}
	if strings.HasSuffix(url, ".xaml") {
		fmt.Println("检查是否设置版本号(.ini)")
		check(url+".ini", true, true,
			"[✓] 正常", "[X] 异常：", 128)
	} else {
		fmt.Println("检查是否设置版本号(version)")
		check(url+"/version", true, true,
			"[✓] 正常", "[X] 异常：", 128)
	}
	fmt.Println("检查是否设置 UA 过滤")
	check(url, true, false,
		"[X] 未设置", "[✓] 已设置:", -1)
	fmt.Println("检查是否设置 Referer 过滤")
	check(url, false, true,
		"[X] 未设置", "[✓] 已设置:", -1)
}
