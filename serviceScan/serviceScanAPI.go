package serviceScan

import (
	"CloudSword/structure"
	"time"
)

func ServiceScanAPI(portScanRet []structure.PortStruct) (ret []structure.HostInfoStruct) {
	tmp := structure.PortStruct2HostPortStruct(portScanRet)
	ret = networkSegmentScan(tmp, time.Second)
	return ret
}
