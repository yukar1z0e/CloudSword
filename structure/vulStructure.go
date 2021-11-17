package structure

type VulStruct struct {
	VulName    string
	VulLevel   string
	VulProduct string
	VulReq     vulrequest
	VulRes     vulresponse
}

type vulrequest struct {
	AddUrl  string
	Method  string
	Payload string
}

type vulresponse struct {
	Key       string
	Value     string
	Operation string
}

type VulCheckStruct struct {
	VulHost   string
	VulName   string
	VulDetail string
}

type ResponseBody struct {
	RespCode    int
	RespContent string
}
