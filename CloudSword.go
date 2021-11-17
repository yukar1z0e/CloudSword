package main

import (
	"CloudSword/dirScan"
	"CloudSword/fingerScan"
	"CloudSword/portScan"
	"CloudSword/structure"
	"CloudSword/vulScan"
	"flag"
	"fmt"
	"github.com/fatih/color"
	"os"
	"strings"
)

var ver = "4.0"

func help() {
	color.Green("\nHelp:")
	fmt.Println("./CloudSword --config [类型] -h [主机]/24 -p [端口]")
	fmt.Println("-p all 全端口")
}

func main() {
	var host string
	flag.StringVar(&host, "h", "", "主机IP")
	var url string
	flag.StringVar(&url, "u", "", "url")
	var port string
	flag.StringVar(&port, "p", "", "端口")
	var config string
	flag.StringVar(&config, "config", "", "扫描类型")
	var module string
	flag.StringVar(&module, "m", "", "字典类型")
	color.Yellow("云刃 " + ver + " by s1by3")
	ParLen := len(os.Args)
	flag.Parse()
	if ParLen == 1 {
		help()
		os.Exit(0)
	} else if strings.ToLower(config) == "portscan" {
		fmt.Println("start portScan host is: " + host + " port " + port)
		if host != "" {
			ret := portScan.HostCheckAPI(host, port)
			fmt.Println(ret)
		}
	} else if strings.ToLower(config) == "fingerscan" {
		fmt.Println("start fingerScan url is: " + url)
		var ret []structure.FingerStructure
		ret = fingerScan.FingerScanAPI(url)
		if ret != nil {
			for _, v := range ret {
				fmt.Println("finger: " + v.FingerName + " matched")
			}
		} else {
			fmt.Println("no finger matched")
		}
	} else if strings.ToLower(config) == "dirscan" {
		fmt.Println("start dirScan url is: " + url)
		fmt.Println("start dirScan dirlib is: " + module)
		dirScan.DirScanAPI(url, module, 200)
	} else if strings.ToLower(config) == "vulscan" {
		fmt.Println("start vulScan url is: " + url)
		fmt.Println("vulLib is: " + module)
		vulScan.VulScanAPI(url, module, 200)
	}
}
