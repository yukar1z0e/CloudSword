package fingerScan

import (
	"CloudSword/lib"
	"CloudSword/structure"
)

func FingerScanAPI(baseUrl string) (ret []structure.FingerStructure) {
	fingerLibMap := lib.GetFingerLib()
	var fingerStructureList []structure.FingerStructure
	for i := 1; i <= 150; i++ {
		var tmpFingerStructure structure.FingerStructure
		v := fingerLibMap[i]
		tmpFingerStructure.FingerId = i
		tmpFingerStructure.FingerName = v["program_name"]
		tmpFingerStructure.Status = false
		tmpFingerStructure.AddonUrl = v["url"]
		tmpFingerStructure.RecognitionTypeId = v["recognitionType_id"]
		tmpFingerStructure.RecognitionContent = v["recognition_content"]
		fingerStructureList = append(fingerStructureList, tmpFingerStructure)
	}
	tmp := fingerScanPool(baseUrl, fingerStructureList)
	ret = append(ret, tmp...)
	fingerStructureList = nil
	for i := 151; i <= 300; i++ {
		var tmpFingerStructure structure.FingerStructure
		v := fingerLibMap[i]
		tmpFingerStructure.FingerId = i
		tmpFingerStructure.FingerName = v["program_name"]
		tmpFingerStructure.Status = false
		tmpFingerStructure.AddonUrl = v["url"]
		tmpFingerStructure.RecognitionTypeId = v["recognitionType_id"]
		tmpFingerStructure.RecognitionContent = v["recognition_content"]
		fingerStructureList = append(fingerStructureList, tmpFingerStructure)
	}
	tmp = fingerScanPool(baseUrl, fingerStructureList)
	if tmp != nil {
		ret = tmp
		return ret
	} else {
		//ret = append(ret, tmp...)
		fingerStructureList = nil
		for i := 301; i <= 450; i++ {
			var tmpFingerStructure structure.FingerStructure
			v := fingerLibMap[i]
			tmpFingerStructure.FingerId = i
			tmpFingerStructure.FingerName = v["program_name"]
			tmpFingerStructure.Status = false
			tmpFingerStructure.AddonUrl = v["url"]
			tmpFingerStructure.RecognitionTypeId = v["recognitionType_id"]
			tmpFingerStructure.RecognitionContent = v["recognition_content"]
			fingerStructureList = append(fingerStructureList, tmpFingerStructure)
		}
		tmp = fingerScanPool(baseUrl, fingerStructureList)
		if tmp != nil {
			ret = tmp
			return ret
		} else {
			//ret = append(ret, tmp...)
			fingerStructureList = nil
			for i := 451; i <= 600; i++ {
				var tmpFingerStructure structure.FingerStructure
				v := fingerLibMap[i]
				tmpFingerStructure.FingerId = i
				tmpFingerStructure.FingerName = v["program_name"]
				tmpFingerStructure.Status = false
				tmpFingerStructure.AddonUrl = v["url"]
				tmpFingerStructure.RecognitionTypeId = v["recognitionType_id"]
				tmpFingerStructure.RecognitionContent = v["recognition_content"]
				fingerStructureList = append(fingerStructureList, tmpFingerStructure)
			}
			tmp = fingerScanPool(baseUrl, fingerStructureList)
			if tmp != nil {
				ret = tmp
				return ret
			} else {
				//ret = append(ret, tmp...)
				fingerStructureList = nil
				for i := 601; i <= 653; i++ {
					var tmpFingerStructure structure.FingerStructure
					v := fingerLibMap[i]
					tmpFingerStructure.FingerId = i
					tmpFingerStructure.FingerName = v["program_name"]
					tmpFingerStructure.Status = false
					tmpFingerStructure.AddonUrl = v["url"]
					tmpFingerStructure.RecognitionTypeId = v["recognitionType_id"]
					tmpFingerStructure.RecognitionContent = v["recognition_content"]
					fingerStructureList = append(fingerStructureList, tmpFingerStructure)
				}
				tmp = fingerScanPool(baseUrl, fingerStructureList)
				if tmp != nil {
					ret = tmp
					return ret
				}
				//ret = append(ret, tmp...)
				fingerStructureList = nil
			}
		}
	}
	return ret
}
