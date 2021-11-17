package portScan

import (
	"CloudSword/structure"
	"time"
)

func singleHostCheck(host string, portList []int, n time.Duration) (ret []structure.PortStruct) {
	port := make(chan int)
	result := make(chan structure.PortStruct)
	//创建与端口数一样多的线程
	for i := 0; i < len(portList); i++ {
		go singlePortCheck(host, port, result, n)
	}
	//传入端口
	for j := 0; j < len(portList); j++ {
		port <- portList[j]
	}
	close(port)
	//输出结果
	for a := 0; a < len(portList); a++ {
		tmp := <-result
		if tmp.Status != false {
			ret = append(ret, tmp)
		}

	}
	return ret
}

func networkSegmentScan(hostList []string, portList []int, n time.Duration) (ret []structure.PortStruct) {
	host := make(chan string)
	resultList := make(chan []structure.PortStruct)

	for i := 0; i < len(hostList); i++ {
		go singleHostScan(host, portList, resultList, n)
	}
	for j := 0; j < len(hostList); j++ {
		host <- hostList[j]
	}
	for a := 0; a < len(hostList); a++ {
		tmp := <-resultList
		if tmp != nil {
			ret = append(ret, tmp...)
		}
	}
	return ret
}

/*//测试函数
func test()  {
	t1 := time.Now()
	var ports=[]int{}
	for i:=1;i<65535;i++{
		ports=append(ports,i)
	}
	ret:=singleHostCheck("192.168.137.1",ports)
	if ret!=nil{
		for _,r:=range ret {
			fmt.Println(r)
		}
	}
	t2 := time.Now()
	fmt.Println(t2.Sub(t1))
}*/
