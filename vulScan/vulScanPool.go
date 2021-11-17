package vulScan

import "CloudSword/structure"

func vulScanPool(baseUrl string, vulStructureList []structure.VulStruct) (ret []structure.VulCheckStruct) {
	vulStructure := make(chan structure.VulStruct)
	result := make(chan structure.VulCheckStruct)
	for i := 0; i < len(vulStructureList); i++ {
		go vulScan(baseUrl, vulStructure, result)
	}
	for j := 0; j < len(vulStructureList); j++ {
		vulStructure <- vulStructureList[j]
	}
	for a := 0; a < len(vulStructureList); a++ {
		tmp := <-result
		if tmp.VulHost != "" && tmp.VulName != "" {
			ret = append(ret, tmp)
		}
	}
	return ret
}
