package dirScan

import (
	"CloudSword/structure"
	"CloudSword/util"
	"fmt"
	"strconv"
)

func dirScan(baseUrl string, dirStructure chan structure.DirStructure, result chan structure.DirStructure) {
	addonStructure := <-dirStructure
	addonUrl := addonStructure.AddonUrl
	addonStructure.Status = false
	url := baseUrl + addonUrl
	resp, err := util.HttpRequest(url)
	if err == nil && resp != nil && resp.StatusCode != 404 {
		addonStructure.Status = true
		addonStructure.Code = resp.StatusCode
		if addonStructure.Code == 200 {
			fmt.Printf("\x1b[%dm %v \x1b[0m  \x1b[%dm 200 \x1b[0m\n", 34, url, 33)
		} else if addonStructure.Code == 301 || addonStructure.Code == 302 {
			fmt.Printf("\x1b[%dm %v \x1b[0m  \x1b[%dm %v \x1b[0m\n", 34, url, 32, strconv.Itoa(resp.StatusCode))
		} else {
			fmt.Printf("\x1b[%dm %v \x1b[0m  \x1b[%dm %v \x1b[0m\n", 34, url, 34, strconv.Itoa(resp.StatusCode))
		}
	} else {
		addonStructure.Status = false
	}
	result <- addonStructure
}
