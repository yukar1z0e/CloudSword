package portScan

import (
	"CloudSword/structure"
	"strconv"
	"strings"
	"time"
)

var defaultPorts = []int{21, 22, 23, 25, 53, 80, 81, 110, 111, 123, 135, 137, 139, 161, 389, 443, 445, 465, 500, 515, 520, 523, 548, 623, 636, 873, 902, 1080, 1099, 1433, 1521, 1604, 1645, 1701, 1883, 1900, 2049, 2181, 2375, 2379, 2425, 3128, 3306, 3389, 4730, 5060, 5222, 5351, 5353, 5432, 5555, 5601, 5672, 5683, 5900, 5938, 5984, 6000, 6379, 7001, 7077, 8080, 8081, 8443, 8545, 8686, 9000, 9001, 9042, 9092, 9100, 9200, 9418, 9999, 11211, 27017, 37777, 50000, 50070, 61616}

func HostCheckAPI(host string, port string) (ret []structure.PortStruct) {
	var portList []int
	var hostList []string
	hostTmp := strings.Split(host, "/")
	if len(hostTmp) == 2 {
		portList = defaultPorts
		hostC := strings.Split(hostTmp[0], ".")
		for i := 1; i < 255; i++ {
			hostMain := hostC[0] + "." + hostC[1] + "." + hostC[2] + "."
			tmp := hostMain + strconv.Itoa(i)
			hostList = append(hostList, tmp)
		}
		//扫IP段 一次800线程
		for i := 0; i <= 11; i++ {
			tmp := networkSegmentScan(hostList[20*i:20*i+19], portList, time.Millisecond*400)
			ret = append(ret, tmp...)
		}
		ret = append(ret, networkSegmentScan(hostList[240:253], portList, time.Millisecond*400)...)
	} else if len(hostTmp) == 1 {
		portTmp := strings.Split(port, "-")
		portTmp2 := strings.Split(port, ",")
		if len(portTmp) == 2 {
			beginPort, _ := strconv.Atoi(portTmp[0])
			endPort, _ := strconv.Atoi(portTmp[1])
			for i := beginPort; i <= endPort; i++ {
				portList = append(portList, i)
			}
			ret = singleHostCheck(host, portList, time.Second)
		} else {
			if len(portTmp2) != 1 {
				for _, ps := range portTmp2 {
					p, _ := strconv.Atoi(ps)
					portList = append(portList, p)
				}
				ret = singleHostCheck(host, portList, time.Second)
			} else {
				if port == "all" {
					//扫全端口 一次100个线程
					for i := 1; i < 65536; i++ {
						portList = append(portList, i)
					}
					for i := 0; i <= 654; i++ {
						tmp := singleHostCheck(host, portList[100*i:100*i+99], time.Millisecond*400)
						ret = append(ret, tmp...)
					}
					ret = append(ret, singleHostCheck(host, portList[65500:65534], time.Millisecond*400)...)
				} else {
					if port != "" {
						p, _ := strconv.Atoi(port)
						portList = append(portList, p)
						ret = singleHostCheck(host, portList, time.Second)
					} else {
						ret = singleHostCheck(host, defaultPorts, time.Second)
					}
				}
			}
		}
	}
	return ret
}
