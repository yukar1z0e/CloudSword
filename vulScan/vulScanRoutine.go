package vulScan

import (
	"CloudSword/structure"
	"CloudSword/util"
	"fmt"
	"strings"
	"time"
)

func vulScan(baseUrl string, vulStructure chan structure.VulStruct, result chan structure.VulCheckStruct) {
	vulInfo := <-vulStructure
	vulName := vulInfo.VulName
	reqInfo := vulInfo.VulReq
	addUrl := reqInfo.AddUrl
	method := reqInfo.Method
	payload := reqInfo.Payload
	resInfo := vulInfo.VulRes
	//operation := resInfo.Operation
	key := resInfo.Key
	value := resInfo.Value
	url := baseUrl + addUrl
	var tmp structure.VulCheckStruct
	var respBody structure.ResponseBody
	if method == "GET" {
		respBody, _ = util.DoGet(url, time.Second)
	} else if method == "POST" {
		respBody, _ = util.DoPost(url, payload, time.Second)
	}
	if strings.Contains(strings.ToLower(respBody.RespContent), strings.ToLower(key)) && strings.Contains(strings.ToLower(respBody.RespContent), strings.ToLower(value)) {
		tmp.VulHost = baseUrl
		tmp.VulName = vulName
		tmp.VulDetail = "url: " + url + " method: " + method + " param: " + payload
		fmt.Println("\x1b[33m" + baseUrl + "\x1b[0m" + " \x1b[31m存在 " + tmp.VulName + "漏洞\x1b[0m" + "\x1b[32m vul detail：" + tmp.VulDetail + "\x1b[0m")
		result <- tmp
	} else {
		tmp.VulHost = baseUrl
		tmp.VulName = ""
		tmp.VulDetail = ""
		fmt.Println("\x1b[34m" + baseUrl + " 不存在 " + vulName + "\x1b[0m")
		result <- tmp
	}
}
