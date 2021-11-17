package structure

import (
	"strconv"
)

type PortStruct struct {
	Host   string
	Port   string
	Status bool
}

type HostInfoStruct struct {
	Host    string
	Port    int
	Service string
}

type HostPortStruct struct {
	Host  string
	Ports []int
}

func PortStruct2HostPortStruct(hostInfoList []PortStruct) (ret []HostPortStruct) {
	hostInfoListLen := len(hostInfoList)
	newHost := hostInfoList[0].Host
	newPort, _ := strconv.Atoi(hostInfoList[0].Port)
	tmp := HostPortStruct{newHost, []int{newPort}}
	ret = append(ret, tmp)
	for i := 1; i < hostInfoListLen; i++ {
		retLen := len(ret)
		newHost := hostInfoList[i].Host
		newPort, _ := strconv.Atoi(hostInfoList[i].Port)
		//fmt.Println(newHost,newPort)
		for j := 0; j < retLen; j++ {
			if newHost == ret[j].Host {
				//fmt.Println(ret[j].Ports)
				ret[j].Ports = append(ret[j].Ports, newPort)
				//fmt.Println(ret[j].Ports)
			} else {
				tmp := HostPortStruct{newHost, []int{newPort}}
				ret = append(ret, tmp)
			}
		}
	}

	return ret
}
