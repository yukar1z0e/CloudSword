package fingerScan

import (
	"CloudSword/structure"
	"CloudSword/util"
	"crypto/md5"
	"fmt"
	"strings"
)

func fingerScan(baseUrl string, fingerStructure chan structure.FingerStructure, result chan structure.FingerStructure) {
	ret := false
	addonStruct := <-fingerStructure
	addonUrl := addonStruct.AddonUrl
	recognitionTypeId := addonStruct.RecognitionTypeId
	recognitionContent := strings.ToLower(addonStruct.RecognitionContent)
	url := baseUrl + addonUrl
	resp, err := util.HttpRequest(url)
	//fmt.Println("check url " + url)
	if resp != nil && resp.StatusCode == 200 && err == nil {
		if recognitionTypeId == "1" {
			res, err := util.GetContent(url)
			if res != nil && err == nil {
				md5Byte := md5.Sum(res)
				md5Value := fmt.Sprintf("%x", md5Byte)
				if md5Value == recognitionContent {
					ret = true
					fmt.Println("check url " + url + " success")
				} else {
					ret = false
				}
			} else {
				ret = false
			}
		} else if recognitionTypeId == "2" {
			res, err := util.GetContent(url)
			if res != nil && err == nil {
				if strings.Contains(strings.ToLower(string(res)), recognitionContent) {
					ret = true
					//fmt.Println("check url " + url + " success")
				} else {
					ret = false
				}
			} else {
				ret = false
			}
		} else if recognitionTypeId == "3" {
			for _, v := range resp.Header {
				if ret == true {
					break
				} else {
					for _, d := range v {
						if strings.Contains(strings.ToLower(d), recognitionContent) {
							ret = true
							fmt.Println("check url " + url + " success")
							break
						}
					}
				}
			}
		} else {
			ret = false
		}
	} else {
		ret = false
	}
	addonStruct.Status = ret
	result <- addonStruct
}
