package dirScan

import (
	"CloudSword/lib"
	"CloudSword/structure"
)

func DirScanAPI(baseUrl string, module string, thread int) {
	dirLib := lib.GetDirLib(module)
	//取余
	y := len(dirLib) % thread
	//循环次数
	x := len(dirLib) / thread
	//fmt.Println(y, x, len(dirLib))
	for i := 0; i < x; i++ {
		dirStructureList := dirLib[thread*i : thread*(i+1)-1]
		//fmt.Println(dirStructureList)
		dirScanPool(baseUrl, dirStructureList)
	}
	var dirStructureTmpList []structure.DirStructure
	for i := 0; i < y; i++ {
		dirStructureTmpList = append(dirStructureTmpList, dirLib[thread*x+i])
	}
	//fmt.Println(dirStructureTmpList)
	dirScanPool(baseUrl, dirStructureTmpList)
}
