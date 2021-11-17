package portScan

import (
	"CloudSword/structure"
	"fmt"
	"net"
	"strconv"
	"time"
)

//扫描单个IP的Routine
func singlePortCheck(host string, port chan int, result chan structure.PortStruct, n time.Duration) {
	p := strconv.Itoa(<-port)
	var ret structure.PortStruct
	ret.Port = p
	ret.Host = host
	conn, err := net.DialTimeout("tcp", host+":"+p, n)
	if err != nil {
		ret.Status = false
		//fmt.Println(host, p, "Close")
	} else {
		fmt.Println(host, p, "Open")
		conn.Close()
		ret.Status = true
	}
	result <- ret
}

//扫描多个IP的Routine
func singleHostScan(host chan string, portList []int, resultList chan []structure.PortStruct, n time.Duration) {
	ip := <-host
	//fmt.Println("start scan: "+ip)
	retList := singleHostCheck(ip, portList, n)
	resultList <- retList
}
