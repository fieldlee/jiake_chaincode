package common

var ERR = map[string]string{
	"NONE":        "000",
	"NOREGISTER":  "001",
	"HADREGISTER": "002",
	"NOOUTPUT":    "003",
	"HADOUTPUT":   "004",
	"NOID":        "005",
	"STATUSERR":   "006",
	"CHANGEERR":   "007",
	"CHAINERR":    "099",
}

var STATUS = map[string]string{
	"INMODULE":    "inModule",
	"OUTMODULE":   "outModule",
	"EXAM":        "exam",
	"WAITBUTCHER": "waitButcher",
	"BUTCHER":     "butcher",
	"FREEZE":      "freeze",
	"BLOCKED":     "blocked",
	"LOST":        "lost",
}

const (
	//下划线
	ULINE = "_"
	//未确认
	UNCOMFIRM = "UNCONFIRM"

	//tx 总数
	TX_COUNT  = "TX_NUMBER_AMOUNT"
	TX_NUMBER = "1234567890"
	//产品注册所属
	SYSTEM = "None"
	//产品所属KEY
	PRODUCT_OWNER = "PRODUCT_OWNER"
	//产品信息KEY
	PRODUCT_INFO = "PRODUCT_INFO"
	// 产品交易信息
	PRODUCT_TRANSFER = "PRODUCT_TRANSFER"
	// 产品登记信息
	PRODUCT_ACTION = "PRODUCT_ACTION"
)
