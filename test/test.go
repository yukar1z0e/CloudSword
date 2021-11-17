package main

import "CloudSword/vulScan"

func main() {
	vulScan.VulScanAPI("http://xxx:8080/example/HelloWorld.action", "struts", 10)
}
