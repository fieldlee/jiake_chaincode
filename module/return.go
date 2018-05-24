package module

type ChanInfo struct {
	ProductId string `json:"productId"`
	Status    bool   `json:"status"`
	ErrorCode string `json:"errorCode"`
}

type ReturnErrorInfo struct {
	Success   bool       `json:"success"`
	ErrorList []ChanInfo `json:"errorList"`
}

type ReturnInfo struct {
	Success bool   `json:"success"`
	Info    string `json:"info"`
}
