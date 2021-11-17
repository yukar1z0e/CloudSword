package dirScan

import "CloudSword/structure"

func dirScanPool(baseUrl string, dirStructureList []structure.DirStructure) (ret []structure.DirStructure) {
	dirStructure := make(chan structure.DirStructure)
	result := make(chan structure.DirStructure)
	for i := 0; i < len(dirStructureList); i++ {
		go dirScan(baseUrl, dirStructure, result)
	}
	for j := 0; j < len(dirStructureList); j++ {
		dirStructure <- dirStructureList[j]
	}
	close(dirStructure)
	for a := 0; a < len(dirStructureList); a++ {
		tmp := <-result
		if tmp.Status {
			ret = append(ret, tmp)
		}
	}
	return ret
}
