package module

type ChanInfo struct {
	ProductId string `json:"productId"`
	Status    bool   `json:"status"`
	ErrorCode string `json:"errorCode"`
}

type ReturnErrorInfo struct {
	Status    bool       `json:"status"`
	ErrorList []ChanInfo `json:"errorList"`
}

type ReturnInfo struct {
	Status bool   `json:"status"`
	Info   string `json:"info"`
}
