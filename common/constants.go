package common

var ERR = map[string]interface{}{
	"NORegister":"001",
	"HadRegister":"002",
	"NOOutput":"003",
	"HadOutput":"004",
	"NoId":"005"
}

const (
	//下划线
	ULINE = "_"
	//未确认
	UNCOMFIRM = "UNCONFIRM"
	//产品注册所属
	SYSTEM = "None"
	//产品所属KEY
	PRODUCT_OWNER = "PRODUCT_OWNER"
	//交易类型（产品注册）
	REGISTER = "01"
	//交易类型(权属变更)
	CHANGE_OWNER = "02"
	//交易类型(确认权属变更)
	CONFIRM_CHANGE_OWNER = "03"
	//交易类型(产品信息变更)
	CHANGE_PRODUCT = "04"
	//交易类型(产品销毁)
	DESTROY = "99"
	//产品信息KEY
	PRODUCT_INFO = "PRODUCT_INFO"
	//产品ID
	PRODUCT_ID = "productId"
)
