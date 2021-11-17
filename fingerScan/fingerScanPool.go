package fingerScan

import (
	"CloudSword/structure"
)

func fingerScanPool(baseUrl string, fingerStructureList []structure.FingerStructure) (ret []structure.FingerStructure) {
	fingerStructure := make(chan structure.FingerStructure)
	result := make(chan structure.FingerStructure)
	for i := 0; i < len(fingerStructureList); i++ {
		go fingerScan(baseUrl, fingerStructure, result)
	}
	for j := 0; j < len(fingerStructureList); j++ {
		fingerStructure <- fingerStructureList[j]
	}
	close(fingerStructure)
	for a := 0; a < len(fingerStructureList); a++ {
		tmp := <-result
		if tmp.Status != false {
			ret = append(ret, tmp)
		}
	}
	return ret
}
