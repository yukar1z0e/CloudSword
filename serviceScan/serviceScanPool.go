package serviceScan

import (
	"CloudSword/portScan"
	"CloudSword/structure"
	"fmt"
	"time"
)

func singleHostCheck(host string, portList []int, timeout time.Duration) (ret []structure.HostInfoStruct) {
	retmp := make(chan structure.HostInfoStruct)

	for i := 0; i < len(portList); i++ {
		go assume_telnet(host, portList[i], timeout, retmp)
		go assume_redis(host, portList[i], timeout, retmp)
		go assume_mongodb(host, portList[i], timeout, retmp)
		go assume_mysql(host, portList[i], timeout, retmp)
		go assume_ssh(host, portList[i], timeout, retmp)
		go assume_ftp(host, portList[i], timeout, retmp)
		go assume_http(host, portList[i], timeout, retmp)
		go assume_mssql(host, portList[i], timeout, retmp)
		a, b, c, d, e, f, g, h := <-retmp, <-retmp, <-retmp, <-retmp, <-retmp, <-retmp, <-retmp, <-retmp
		i := structure.HostInfoStruct{Host: host, Port: portList[i], Service: "open"}
		if a.Service != "" {
			ret = append(ret, a)
		} else if b.Service != "" {
			ret = append(ret, b)
		} else if c.Service != "" {
			ret = append(ret, c)
		} else if d.Service != "" {
			ret = append(ret, d)
		} else if e.Service != "" {
			ret = append(ret, e)
		} else if f.Service != "" {
			ret = append(ret, f)
		} else if g.Service != "" {
			ret = append(ret, g)
		} else if h.Service != "" {
			ret = append(ret, h)
		} else {
			ret = append(ret, i)
		}
	}
	return ret
}

func singleHostScan(host chan structure.HostPortStruct, timeout time.Duration, result chan []structure.HostInfoStruct) {
	hostPice := <-host
	ip := hostPice.Host
	portList := hostPice.Ports
	ret := singleHostCheck(ip, portList, timeout)
	result <- ret
}

func networkSegmentScan(hostList []structure.HostPortStruct, timeout time.Duration) (ret []structure.HostInfoStruct) {
	host := make(chan structure.HostPortStruct)
	resultList := make(chan []structure.HostInfoStruct)
	lenth := len(hostList)
	for i := 0; i < lenth; i++ {
		go singleHostScan(host, timeout, resultList)
	}
	for i := 0; i < lenth; i++ {
		host <- hostList[i]
	}
	for i := 0; i < lenth; i++ {
		tmp := <-resultList
		ret = append(ret, tmp...)
	}
	return ret
}

func Test1() {

	ret := portScan.HostCheckAPI("192.168.137.1/24", "")
	fmt.Println(ret)
	ret1 := structure.PortStruct2HostPortStruct(ret)
	//for i := 0; i < len(ret1); i++ {
	//	fmt.Println(ret1[i])
	//}
	fmt.Println(ret1)
	//var hostList = []structure.HostPortStruct{}
	//hostList = append(hostList, structure.HostPortStruct{Host: "192.168.137.1", Ports: []int{3389, 3306, 135}})
	//hostList=append(hostList,structure.HostPortStruct{Host: "47.100.234.207",Ports: []int{27017}})
	ret2 := networkSegmentScan(ret1, time.Second)
	fmt.Println(ret2)
	//for i := 0; i < len(ret2); i++ {
	//	fmt.Println(ret2[i])
	//}
}
