package module

type RegitserParam struct {
	ProductId   string `json:"productId"`
	Operation   string `json:"operation"`
	Operator    string `json:"operator"`
	CreateTime  uint64 `json:"createTime"`
	BatchNumber string `json:"batchNumber"` // 入栏批次
	Kind        string `json:"kind"`        // 种类
	Type        string `json:"type"`        // 品种
	MapPosition string `json:"mapPosition"` // 入栏地理位置
}

type FeedParam struct {
	ProductId   string `json:"productId"`
	Operation   string `json:"operation"`
	Operator    string `json:"operator"`
	FeedName    string `json:"feedName"`    //喂养的饲料
	FeedTime    uint64 `json:"feedTime"`    //喂养的时间
	MapPosition string `json:"mapPosition"` // 喂养地理位置
}

// 防疫结构
type VaccineParam struct {
	ProductId     string `json:"productId"`
	Operation     string `json:"operation"`     //防疫类型
	Operator      string `json:"operator"`      //防疫人
	VaccineName   string `json:"vaccineName"`   //防疫的药品名称
	VaccineType   string `json:"vaccineType"`   //防疫项目
	VaccineNumber string `json:"vaccineNumber"` //防疫药品的数量
	VaccineTime   uint64 `json:"vaccineTime"`   //防疫时间
	MapPosition   string `json:"mapPosition"`   // 地理位置
}

// 救治结构
type SaveParam struct {
	ProductId       string `json:"productId"`
	Operation       string `json:"operation"`       //救治类型
	Operator        string `json:"operator"`        //救治人
	SaveName        string `json:"saveName"`        //救治的药品名称
	SaveType        string `json:"saveType"`        //救治项目
	SaveNumber      string `json:"saveNumber"`      //救治药品的数量
	SaveConsequence string `json:"saveConsequence"` //救治结果
	SaveTime        uint64 `json:"saveTime"`        //救治时间
	MapPosition     string `json:"mapPosition"`     // 地理位置
}

// 检疫结构
type ExamParam struct {
	ProductId       string `json:"productId"`
	Operation       string `json:"operation"`       //检疫类型
	Operator        string `json:"operator"`        //检疫人
	ExamTime        uint64 `json:"examTime"`        //防疫时间
	ExamConsequence string `json:"examConsequence"` //防疫结果
	MapPosition     string `json:"mapPosition"`     // 地理位置
}

// 灭尸结构
type DestroyParam struct {
	ProductId   string `json:"productId"`
	LostReason  string `json:"lostReason"`  //灭尸原因
	LostWay     string `json:"lostWay"`     //灭尸方式
	Operation   string `json:"operation"`   //灭尸类型
	Operator    string `json:"operator"`    //灭尸人
	LostTime    uint64 `json:"lostTime"`    //灭尸时间
	MapPosition string `json:"mapPosition"` //地理位置
}

// 出栏结构
type OutputParam struct {
	ProductId   string `json:"productId"`
	Operation   string `json:"operation"`   //出栏类型
	Operator    string `json:"operator"`    //出栏人
	OutputTime  uint64 `json:"outputTime"`  //出栏时间
	MapPosition string `json:"mapPosition"` //地理位置
}

// 待宰结构
type WaitButcherParam struct {
	ProductId       string `json:"productId"`
	Operation       string `json:"operation"`       //出栏类型
	Operator        string `json:"operator"`        //待宰人
	WaitButcherTime uint64 `json:"waitButcherTime"` //待宰时间
	MapPosition     string `json:"mapPosition"`     //地理位置
}

// 屠宰结构
type ButcherParam struct {
	ProductId   string `json:"productId"`
	HookNo      string `json:"hookNo"`      //挂钩号码
	Operation   string `json:"operation"`   //出栏类型
	Operator    string `json:"operator"`    //屠宰人
	ButcherTime uint64 `json:"butcherTime"` //屠宰时间
	MapPosition string `json:"mapPosition"` //地理位置
}

// 资产号查询
type QueryParam struct {
	ProductId string `json:"productId"`
}

// 交易号查询
type QueryTxParam struct {
	TxId string `json:"txId"`
}

// 批次号查询
type BatchParam struct {
	BatchNumber string `json:"batchNumber"`
}

// 分块信息
type BlockParam struct {
	BlockId   string `json:"blockId"`
	BlockName string `json:"blockName"`
	ParentId  string `json:"parentId"`
	BlockTime uint64 `json:"blockTime"`
}
