package vulScan

import (
	"CloudSword/structure"
	"CloudSword/vulScan/payload"
)

func VulScanAPI(baseUrl string, module string, thread int) {
	vulLib := payload.GetPayloadLib(module)
	//取余
	y := len(vulLib) % thread
	//循环次数
	x := len(vulLib) / thread
	for i := 0; i < x; i++ {
		vulStructureList := vulLib[thread*i : thread*(i+1)-1]
		vulScanPool(baseUrl, vulStructureList)
	}
	var tmpVulStructureList []structure.VulStruct
	for i := 0; i < y; i++ {
		tmpVulStructureList = append(tmpVulStructureList, vulLib[thread*x+i])
	}
	vulScanPool(baseUrl, tmpVulStructureList)
}
