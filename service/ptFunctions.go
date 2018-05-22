package service

import (
	"encoding/json"
	"jiakechaincode/common"
	"jiakechaincode/log"
	"jiakechaincode/module"

	"github.com/hyperledger/fabric/core/chaincode/shim"
)

func goRegister(stub shim.ChaincodeStubInterface, param module.RegitserParam, regChan chan ChanInfo) {
	defer wg.Done()
	tChan := ChanInfo{}
	tChan.ProductId = param.ProductId
	log.Logger.Info("goRegister--productid:" + param.ProductId)
	// 	verify product if exist or not
	jsonParam, err := stub.GetState(common.PRODUCT_INFO + common.ULINE + param.ProductId)
	log.Logger.Info("------------------------------------------------------------------")
	log.Logger.Info(string(jsonParam[:]))

	if jsonParam != nil {
		log.Logger.Error("goRegister -- get product by productid -- err: 已经入栏，不能再次入栏" + "	productid:" + param.ProductId)
		tChan.Status = false
		tChan.ErrorCode = common.ERR["HADREGISTER"]
		regChan <- tChan
		return
	} else {
		if err != nil {
			log.Logger.Error("goRegister -- get product by productid -- err:" + err.Error() + "	productid:" + param.ProductId)
			// tChan.Status = false
			// tChan.ErrorCode = common.ERR["CHAINERR"]
			// regChan <- tChan
			// return
		}
	}

	product := module.Product{}
	product.ProductId = param.ProductId
	product.PreOwner = common.SYSTEM
	product.Type = param.Type
	product.Kind = param.Kind
	product.CreateTime = param.CreateTime
	product.BatchNumber = param.BatchNumber
	product.MapPosition = param.MapPosition
	product.Operation = param.Operation
	product.Operator = param.Operator
	product.PreOwner = common.SYSTEM
	product.CurrentOwner = common.GetUserFromCertification(stub)
	// MODIFY STATUS

	product.Status = common.STATUS["INModule"]

	jsonByte, err := json.Marshal(product)
	if err != nil {
		log.Logger.Error("goRegister -- marshal product err:" + err.Error() + "	productid:" + param.ProductId)
		tChan.Status = false
		tChan.ErrorCode = common.ERR["CHAINERR"]
		regChan <- tChan
		return
	}

	err = stub.PutState(common.PRODUCT_INFO+common.ULINE+param.ProductId, jsonByte)
	if err != nil {
		log.Logger.Error("goRegister -- putState:" + err.Error() + "	productid:" + param.ProductId)
		tChan.Status = false
		tChan.ErrorCode = common.ERR["CHAINERR"]
		regChan <- tChan
		return
	}

	// asset transcation
	changeOwner := module.ChangeAssetOwner{}
	changeOwner.PreOwner = common.SYSTEM
	changeOwner.CurrentOwner = common.GetUserFromCertification(stub)
	changeOwner.ProductId = param.ProductId
	changeOwner.Operation = param.Operation
	changeOwner.Operator = param.Operator
	time, err := stub.GetTxTimestamp()
	if err != nil {
		log.Logger.Error("goRegister -- register change owner get time:" + err.Error() + "	productid:" + param.ProductId)
		tChan.Status = false
		tChan.ErrorCode = common.ERR["CHAINERR"]
		regChan <- tChan
		return
	}
	changeOwner.OperateTime = uint64(time.GetSeconds())
	jsonchangeOwnerBytes, err := json.Marshal(changeOwner)
	err = stub.PutState(common.PRODUCT_TRANSFER+common.ULINE+param.ProductId, jsonchangeOwnerBytes)

	if err != nil {
		log.Logger.Error("goRegister -- PutState change owner:" + err.Error() + "	productid:" + param.ProductId)
		tChan.Status = false
		tChan.ErrorCode = common.ERR["CHAINERR"]
		regChan <- tChan
		return
	}

	tChan.Status = true
	tChan.ErrorCode = common.ERR["NONE"]
	regChan <- tChan
	return
}

func toRegister(stub shim.ChaincodeStubInterface, param module.RegitserParam) (tChan ChanInfo) {
	tChan.ProductId = param.ProductId
	log.Logger.Info("goRegister--productid:" + param.ProductId)
	// 	verify product if exist or not
	jsonParam, err := stub.GetState(common.PRODUCT_INFO + common.ULINE + param.ProductId)
	log.Logger.Info("------------------------------------------------------------------")
	log.Logger.Info(string(jsonParam[:]))

	if jsonParam != nil {
		log.Logger.Error("goRegister -- get product by productid -- err: 已经入栏，不能再次入栏" + "	productid:" + param.ProductId)
		tChan.Status = false
		tChan.ErrorCode = common.ERR["HADREGISTER"]
		return
	} else {
		if err != nil {
			log.Logger.Error("goRegister -- get product by productid -- err:" + err.Error() + "	productid:" + param.ProductId)
			// tChan.Status = false
			// tChan.ErrorCode = common.ERR["CHAINERR"]
			// regChan <- tChan
			// return
		}
	}

	product := module.Product{}
	product.ProductId = param.ProductId
	product.PreOwner = common.SYSTEM
	product.Type = param.Type
	product.Kind = param.Kind
	product.CreateTime = param.CreateTime
	product.BatchNumber = param.BatchNumber
	product.MapPosition = param.MapPosition
	product.Operation = param.Operation
	product.Operator = param.Operator
	product.PreOwner = common.SYSTEM
	product.CurrentOwner = common.GetUserFromCertification(stub)
	// MODIFY STATUS

	product.Status = common.STATUS["INModule"]

	jsonByte, err := json.Marshal(product)
	if err != nil {
		log.Logger.Error("goRegister -- marshal product err:" + err.Error() + "	productid:" + param.ProductId)
		tChan.Status = false
		tChan.ErrorCode = common.ERR["CHAINERR"]

		return
	}

	err = stub.PutState(common.PRODUCT_INFO+common.ULINE+param.ProductId, jsonByte)
	if err != nil {
		log.Logger.Error("goRegister -- putState:" + err.Error() + "	productid:" + param.ProductId)
		tChan.Status = false
		tChan.ErrorCode = common.ERR["CHAINERR"]

		return
	}

	// asset transcation
	changeOwner := module.ChangeAssetOwner{}
	changeOwner.PreOwner = common.SYSTEM
	changeOwner.CurrentOwner = common.GetUserFromCertification(stub)
	changeOwner.ProductId = param.ProductId
	changeOwner.Operation = param.Operation
	changeOwner.Operator = param.Operator
	time, err := stub.GetTxTimestamp()
	if err != nil {
		log.Logger.Error("goRegister -- register change owner get time:" + err.Error() + "	productid:" + param.ProductId)
		tChan.Status = false
		tChan.ErrorCode = common.ERR["CHAINERR"]

		return
	}
	changeOwner.OperateTime = uint64(time.GetSeconds())
	jsonchangeOwnerBytes, err := json.Marshal(changeOwner)
	err = stub.PutState(common.PRODUCT_TRANSFER+common.ULINE+param.ProductId, jsonchangeOwnerBytes)

	if err != nil {
		log.Logger.Error("goRegister -- PutState change owner:" + err.Error() + "	productid:" + param.ProductId)
		tChan.Status = false
		tChan.ErrorCode = common.ERR["CHAINERR"]

		return
	}

	tChan.Status = true
	tChan.ErrorCode = common.ERR["NONE"]
	return
}

func toFeed(stub shim.ChaincodeStubInterface, param module.FeedParam) (fedChan ChanInfo) {
	fedChan.ProductId = param.ProductId
	// 	verify product if exist or not
	jsonParam, err := stub.GetState(common.PRODUCT_INFO + common.ULINE + param.ProductId)
	if err != nil {
		log.Logger.Error("goFeed -- getState by productid:" + err.Error() + "	productid:" + param.ProductId)
		// fedChan.Status = false
		// fedChan.ErrorCode = common.ERR["CHAINERR"]
		// feedChan <- fedChan
		// return
	}
	if jsonParam == nil {
		log.Logger.Error("goFeed -- get product by id , 未入栏，请入栏" + "	productid:" + param.ProductId)
		fedChan.Status = false
		fedChan.ErrorCode = common.ERR["NOREGISTER"]
		return
	}

	product := module.Product{}
	err = json.Unmarshal(jsonParam, &product)
	if err != nil {
		log.Logger.Error("goFeed -- Unmarshal product:" + err.Error() + "	productid:" + param.ProductId)
		fedChan.Status = false
		fedChan.ErrorCode = common.ERR["CHAINERR"]

		return
	}

	if product.Status != common.STATUS["INModule"] { //入栏状态
		log.Logger.Error("goFeed -- 状态不对，目前不是入栏状态" + "	productid:" + param.ProductId)
		fedChan.Status = false
		fedChan.ErrorCode = common.ERR["STATUSERR"]

		return
	}

	// 	create feed obj
	feedObj := module.Feed{}
	feedObj.ProductId = param.ProductId
	feedObj.TxId = stub.GetTxID()
	feedObj.Operation = param.Operation
	feedObj.Operator = param.Operator
	feedObj.FeedName = param.FeedName
	feedObj.FeedTime = param.FeedTime
	feedObj.MapPosition = param.MapPosition

	product.FeedList = append(product.FeedList, feedObj)
	// marshal product into state
	jsonProduct, err := json.Marshal(product)
	if err != nil {
		log.Logger.Error("goFeed -- marshal product:" + err.Error() + "	productid:" + param.ProductId)
		fedChan.Status = false
		fedChan.ErrorCode = common.ERR["CHAINERR"]
		return
	}

	err = stub.PutState(common.PRODUCT_INFO+common.ULINE+param.ProductId, jsonProduct)
	if err != nil {
		log.Logger.Error("goFeed -- putState:" + err.Error() + "	productid:" + param.ProductId)
		fedChan.Status = false
		fedChan.ErrorCode = common.ERR["CHAINERR"]
		return
	}
	fedChan.Status = true
	fedChan.ErrorCode = common.ERR["NONE"]
	return
}

func toVaccine(stub shim.ChaincodeStubInterface, param module.VaccineParam) (vChan ChanInfo) {
	vChan.ProductId = param.ProductId
	// 	verify product if exist or not
	jsonParam, err := stub.GetState(common.PRODUCT_INFO + common.ULINE + param.ProductId)
	if err != nil {
		log.Logger.Error("goVaccine -- getState:" + err.Error() + "	productid:" + param.ProductId)
		// vChan.Status = false
		// vChan.ErrorCode = common.ERR["CHAINERR"]
		// vaccineChan <- vChan
		// return
	}
	if jsonParam == nil {
		log.Logger.Error("goVaccine -- 未入栏")
		vChan.Status = false
		vChan.ErrorCode = common.ERR["NOREGISTER"]
		return
	}

	product := module.Product{}
	err = json.Unmarshal(jsonParam, &product)
	if err != nil {
		log.Logger.Error("goVaccine -- Unmarshal product err:" + err.Error() + "	productid:" + param.ProductId)
		vChan.Status = false
		vChan.ErrorCode = common.ERR["CHAINERR"]
		return
	}
	if product.Status != common.STATUS["INModule"] { //入栏状态
		log.Logger.Error("goVaccine -- 状态不对，目前不是入栏状态" + "	productid:" + param.ProductId)
		vChan.Status = false
		vChan.ErrorCode = common.ERR["STATUSERR"]
		return
	}
	// create vaccine object
	vaccine := module.Vaccine{}
	vaccine.TxId = stub.GetTxID()
	vaccine.ProductId = param.ProductId
	vaccine.Operation = param.Operation
	vaccine.Operator = param.Operator
	// VaccineName   string `json:"vaccineName"`   //防疫的药品名称
	// VaccineType   string `json:"vaccineType"`   //防疫项目
	// VaccineNumber string `json:"vaccineNumber"` //防疫药品的数量
	// VaccineTime   uint64 `json:"vaccineTime"`   //防疫时间
	vaccine.VaccineName = param.VaccineName
	vaccine.VaccineType = param.VaccineType
	vaccine.VaccineNumber = param.VaccineNumber
	vaccine.VaccineTime = param.VaccineTime
	vaccine.MapPosition = param.MapPosition
	product.VaccineList = append(product.VaccineList, vaccine)
	// marshal product into state
	jsonProduct, err := json.Marshal(product)
	if err != nil {
		log.Logger.Error("goVaccine -- Marshal product err :" + err.Error() + "	prodocut:" + param.ProductId)
		vChan.Status = false
		vChan.ErrorCode = common.ERR["CHAINERR"]
		return
	}

	err = stub.PutState(common.PRODUCT_INFO+common.ULINE+param.ProductId, jsonProduct)
	if err != nil {
		log.Logger.Error("goVaccine -- putstate product err :" + err.Error() + "	prodocut:" + param.ProductId)
		vChan.Status = false
		vChan.ErrorCode = common.ERR["CHAINERR"]
		return
	}
	vChan.Status = true
	vChan.ErrorCode = common.ERR["NONE"]
	return
}

func toOutput(stub shim.ChaincodeStubInterface, param module.OutputParam) (vChan ChanInfo) {

	vChan.ProductId = param.ProductId
	// 	verify product if exist or not
	jsonParam, err := stub.GetState(common.PRODUCT_INFO + common.ULINE + param.ProductId)
	if err != nil {
		log.Logger.Error("goOutput -- getstate product err :" + err.Error() + "	prodocut:" + param.ProductId)
		// vChan.Status = false
		// vChan.ErrorCode = common.ERR["CHAINERR"]
		// outputChan <- vChan
		// return
	}
	if jsonParam == nil {
		log.Logger.Error("goOutput -- 未入栏" + "	prodocut:" + param.ProductId)
		vChan.Status = false
		vChan.ErrorCode = common.ERR["NOREGISTER"]

		return
	}

	product := module.Product{}
	err = json.Unmarshal(jsonParam, &product)
	if err != nil {
		log.Logger.Error("goOutput -- Unmarshal product err :" + err.Error() + "	prodocut:" + param.ProductId)
		vChan.Status = false
		vChan.ErrorCode = common.ERR["CHAINERR"]

		return
	}

	if product.Status != common.STATUS["INModule"] {
		log.Logger.Error("goOutput -- 状态不对，目前不是入栏状态" + "	prodocut:" + param.ProductId)
		vChan.Status = false
		vChan.ErrorCode = common.ERR["STATUSERR"]

		return
	}

	if product.CurrentOwner != common.GetUserFromCertification(stub) {
		log.Logger.Error("goOutput -- 操作人不对，不是资产所有人" + "	prodocut:" + param.ProductId)
		vChan.Status = false
		vChan.ErrorCode = common.ERR["CHANGEERR"]

		return
	}

	// create vaccine object

	product.OutputTxId = stub.GetTxID()
	product.OutputMapPosition = param.MapPosition
	product.OutputOperation = param.Operation
	product.OutputTime = param.OutputTime
	product.OutputOperator = param.Operator
	// MODIFY STATUS
	product.Status = common.STATUS["OUTModule"]
	product.PreOwner = product.CurrentOwner
	product.CurrentOwner = common.SYSTEM
	// marshal product into state
	jsonProduct, err := json.Marshal(product)
	if err != nil {
		log.Logger.Error("goOutput -- Marshal product" + err.Error() + "	prodocut:" + param.ProductId)
		vChan.Status = false
		vChan.ErrorCode = common.ERR["CHAINERR"]

		return
	}

	err = stub.PutState(common.PRODUCT_INFO+common.ULINE+param.ProductId, jsonProduct)
	if err != nil {
		log.Logger.Error("goOutput -- PutState product" + err.Error() + "	prodocut:" + param.ProductId)
		vChan.Status = false
		vChan.ErrorCode = common.ERR["CHAINERR"]

		return
	}

	// ASSET CHANGE OWNER === START
	changeOwner := module.ChangeAssetOwner{}
	changeOwner.PreOwner = common.GetUserFromCertification(stub)
	changeOwner.CurrentOwner = common.SYSTEM
	changeOwner.ProductId = param.ProductId
	changeOwner.Operation = param.Operation
	changeOwner.Operator = param.Operator
	time, err := stub.GetTxTimestamp()
	if err != nil {
		log.Logger.Error("goOutput -- goOutput change owner get time:" + err.Error() + "	productid:" + param.ProductId)
		vChan.Status = false
		vChan.ErrorCode = common.ERR["CHAINERR"]

		return
	}
	changeOwner.OperateTime = uint64(time.GetSeconds())
	jsonchangeOwnerBytes, err := json.Marshal(changeOwner)
	err = stub.PutState(common.PRODUCT_TRANSFER+common.ULINE+param.ProductId, jsonchangeOwnerBytes)

	if err != nil {
		log.Logger.Error("goOutput -- PutState change owner:" + err.Error() + "	productid:" + param.ProductId)
		vChan.Status = false
		vChan.ErrorCode = common.ERR["CHAINERR"]

		return
	}
	// ASSET CHANGE OWNER === END
	vChan.Status = true
	vChan.ErrorCode = common.ERR["NONE"]

	return
}

func toExam(stub shim.ChaincodeStubInterface, param module.ExamParam) (vChan ChanInfo) {

	vChan.ProductId = param.ProductId
	// 	verify product if exist or not
	jsonParam, err := stub.GetState(common.PRODUCT_INFO + common.ULINE + param.ProductId)
	if err != nil {
		log.Logger.Error("goExam -- getState product" + err.Error() + "	prodocut:" + param.ProductId)
		// vChan.Status = false
		// vChan.ErrorCode = common.ERR["CHAINERR"]
		// examChan <- vChan
		// return
	}
	if jsonParam == nil {
		log.Logger.Error("goExam -- 查找不到资产，未入栏	prodocut:" + param.ProductId)
		vChan.Status = false
		vChan.ErrorCode = common.ERR["NOREGISTER"]

		return
	}

	product := module.Product{}
	err = json.Unmarshal(jsonParam, &product)
	if err != nil {
		log.Logger.Error("goExam -- Unmarshal product ERR:" + err.Error() + "	prodocut:" + param.ProductId)
		vChan.Status = false
		vChan.ErrorCode = common.ERR["CHAINERR"]

		return
	}

	if product.Status != common.STATUS["OUTModule"] {
		log.Logger.Error("goExam -- 状态不对，目前不是已出栏	prodocut:" + param.ProductId)
		vChan.Status = false
		vChan.ErrorCode = common.ERR["STATUSERR"]

		return
	}
	// ExamOperation   string `json:"examOperation"`   //检疫类型
	// ExamOperator    string `json:"examOperator"`    //检疫人
	// ExamTime        string `json:"examTime"`        //防疫时间
	// ExamResult      string `json:"examResult"`      //防疫结果
	// ExamMapPosition string `json:"examMapPosition"` // 地理位置
	// create vaccine object
	product.ExamTxId = stub.GetTxID()
	product.ExamMapPosition = param.MapPosition
	product.ExamOperation = param.Operation
	product.ExamTime = param.ExamTime
	product.ExamConsequence = param.ExamConsequence
	product.ExamOperator = param.Operator
	// modify status 待宰状态
	product.Status = common.STATUS["BUTCHER"]
	// marshal product into state
	jsonProduct, err := json.Marshal(product)
	if err != nil {
		log.Logger.Error("goExam -- marshal product ERR:" + err.Error() + "	prodocut:" + param.ProductId)
		vChan.Status = false
		vChan.ErrorCode = common.ERR["CHAINERR"]

		return
	}
	err = stub.PutState(common.PRODUCT_INFO+common.ULINE+param.ProductId, jsonProduct)
	if err != nil {
		log.Logger.Error("goExam -- PutState product ERR:" + err.Error() + "	prodocut:" + param.ProductId)
		vChan.Status = false
		vChan.ErrorCode = common.ERR["CHAINERR"]

		return
	}
	vChan.Status = true
	vChan.ErrorCode = common.ERR["NONE"]

	return
}

func toSave(stub shim.ChaincodeStubInterface, param module.SaveParam) (vChan ChanInfo) {

	vChan.ProductId = param.ProductId
	// 	verify product if exist or not
	jsonParam, err := stub.GetState(common.PRODUCT_INFO + common.ULINE + param.ProductId)
	if err != nil {
		log.Logger.Error("goSave -- getState:" + err.Error() + "	productid:" + param.ProductId)
		// vChan.Status = false
		// vChan.ErrorCode = common.ERR["CHAINERR"]
		// saveChan <- vChan
		// return
	}
	if jsonParam == nil {
		log.Logger.Error("goSave -- 未入栏")
		vChan.Status = false
		vChan.ErrorCode = common.ERR["NOREGISTER"]

		return
	}

	product := module.Product{}
	err = json.Unmarshal(jsonParam, &product)
	if err != nil {
		log.Logger.Error("goSave -- Unmarshal product err:" + err.Error() + "	productid:" + param.ProductId)
		vChan.Status = false
		vChan.ErrorCode = common.ERR["CHAINERR"]

		return
	}
	if product.Status != common.STATUS["INModule"] { //入栏状态
		log.Logger.Error("goSave -- 状态不对，目前不是入栏状态" + "	productid:" + param.ProductId)
		vChan.Status = false
		vChan.ErrorCode = common.ERR["STATUSERR"]

		return
	}
	// create save object

	// SaveName    string `json:"saveName"`    //救治的药品名称
	// SaveType    string `json:"saveType"`    //救治项目
	// SaveNumber  string `json:"saveNumber"`  //救治药品的数量
	// SaveTime    uint64 `json:"saveTime"`    //救治时间

	save := module.Save{}
	save.TxId = stub.GetTxID()
	save.ProductId = param.ProductId
	save.Operation = param.Operation
	save.Operator = param.Operator
	save.SaveName = param.SaveName
	save.SaveType = param.SaveType
	save.SaveNumber = param.SaveNumber
	save.SaveTime = param.SaveTime
	save.MapPosition = param.MapPosition
	product.SaveList = append(product.SaveList, save)
	// marshal product into state
	jsonProduct, err := json.Marshal(product)
	if err != nil {
		log.Logger.Error("goSave -- Marshal product err :" + err.Error() + "	prodocut:" + param.ProductId)
		vChan.Status = false
		vChan.ErrorCode = common.ERR["CHAINERR"]

		return
	}

	err = stub.PutState(common.PRODUCT_INFO+common.ULINE+param.ProductId, jsonProduct)
	if err != nil {
		log.Logger.Error("goSave -- putstate product err :" + err.Error() + "	prodocut:" + param.ProductId)
		vChan.Status = false
		vChan.ErrorCode = common.ERR["CHAINERR"]

		return
	}
	vChan.Status = true
	vChan.ErrorCode = common.ERR["NONE"]

	return
}

func toButcher(stub shim.ChaincodeStubInterface, param module.ButcherParam) (vChan ChanInfo) {

	vChan.ProductId = param.ProductId
	// 	verify product if exist or not
	jsonParam, err := stub.GetState(common.PRODUCT_INFO + common.ULINE + param.ProductId)
	if err != nil {
		log.Logger.Error("goButcher -- GetState product ERR:" + err.Error() + "	prodocut:" + param.ProductId)
		// vChan.Status = false
		// vChan.ErrorCode = common.ERR["CHAINERR"]

		// return
	}
	if jsonParam == nil {
		log.Logger.Error("goButcher -- 查找不到资产，未入栏 	prodocut:" + param.ProductId)
		vChan.Status = false
		vChan.ErrorCode = common.ERR["NOREGISTER"]

		return
	}

	product := module.Product{}
	err = json.Unmarshal(jsonParam, &product)
	if err != nil {
		log.Logger.Error("goButcher -- Unmarshal product err:" + err.Error() + " 	prodocut:" + param.ProductId)
		vChan.Status = false
		vChan.ErrorCode = common.ERR["CHAINERR"]

		return
	}

	if product.Status != common.STATUS["BUTCHER"] {
		log.Logger.Error("goButcher -- 状态不对 ，目前不是待屠宰 	prodocut:" + param.ProductId)
		vChan.Status = false
		vChan.ErrorCode = common.ERR["STATUSERR"]

		return
	}

	// create vaccine object
	product.ButcherTxId = stub.GetTxID()
	product.ButcherMapPosition = param.MapPosition
	product.ButcherOperation = param.Operation
	product.ButcherTime = param.ButcherTime
	product.ButcherOperator = param.Operator
	product.HookNo = param.HookNo
	product.CurrentOwner = common.GetUserFromCertification(stub)
	// modify status 除酸
	product.Status = common.STATUS["FREEZE"]
	// marshal product into state
	jsonProduct, err := json.Marshal(product)
	if err != nil {
		log.Logger.Error("goButcher -- marshal product err:" + err.Error() + " 	prodocut:" + param.ProductId)
		vChan.Status = false
		vChan.ErrorCode = common.ERR["CHAINERR"]

		return
	}
	err = stub.PutState(common.PRODUCT_INFO+common.ULINE+param.ProductId, jsonProduct)
	if err != nil {
		log.Logger.Error("goButcher -- PutState product err:" + err.Error() + " 	prodocut:" + param.ProductId)
		vChan.Status = false
		vChan.ErrorCode = common.ERR["CHAINERR"]

		return
	}

	// ASSET CHANGE OWNER === START
	changeOwner := module.ChangeAssetOwner{}
	changeOwner.PreOwner = product.PreOwner
	changeOwner.CurrentOwner = product.CurrentOwner
	changeOwner.ProductId = param.ProductId
	changeOwner.Operation = param.Operation
	changeOwner.Operator = param.Operator
	time, err := stub.GetTxTimestamp()
	if err != nil {
		log.Logger.Error("goButcher -- goButcher change owner get time:" + err.Error() + "	productid:" + param.ProductId)
		vChan.Status = false
		vChan.ErrorCode = common.ERR["CHAINERR"]

		return
	}
	changeOwner.OperateTime = uint64(time.GetSeconds())
	jsonchangeOwnerBytes, err := json.Marshal(changeOwner)
	err = stub.PutState(common.PRODUCT_TRANSFER+common.ULINE+param.ProductId, jsonchangeOwnerBytes)

	if err != nil {
		log.Logger.Error("goButcher -- PutState change owner:" + err.Error() + "	productid:" + param.ProductId)
		vChan.Status = false
		vChan.ErrorCode = common.ERR["CHAINERR"]

		return
	}
	// ASSET CHANGE OWNER === END
	vChan.Status = true
	vChan.ErrorCode = common.ERR["NONE"]

	return
}

func toLost(stub shim.ChaincodeStubInterface, param module.DestroyParam) (vChan ChanInfo) {

	vChan.ProductId = param.ProductId
	// 	verify product if exist or not
	jsonParam, err := stub.GetState(common.PRODUCT_INFO + common.ULINE + param.ProductId)
	if err != nil {
		log.Logger.Error("goLost -- GetState product ERR:" + err.Error() + "	prodocut:" + param.ProductId)
		// vChan.Status = false
		// vChan.ErrorCode = common.ERR["CHAINERR"]
		// lostChan <- vChan
		// return
	}
	if jsonParam == nil {
		log.Logger.Error("goLost -- 查找不到资产，未入栏 	prodocut:" + param.ProductId)
		vChan.Status = false
		vChan.ErrorCode = common.ERR["NOREGISTER"]

		return
	}

	product := module.Product{}
	err = json.Unmarshal(jsonParam, &product)
	if err != nil {
		log.Logger.Error("goLost -- Unmarshal product err:" + err.Error() + " 	prodocut:" + param.ProductId)
		vChan.Status = false
		vChan.ErrorCode = common.ERR["CHAINERR"]

		return
	}

	if product.Status != common.STATUS["INModule"] {
		log.Logger.Error("goLost -- 状态不对 ，目前不是入栏 	prodocut:" + param.ProductId)
		vChan.Status = false
		vChan.ErrorCode = common.ERR["STATUSERR"]

		return
	}

	// create lost object
	product.LostTxId = stub.GetTxID()
	product.LostMapPosition = param.MapPosition
	product.LostOperation = param.Operation
	product.LostTime = param.LostTime
	product.LostOperator = param.Operator
	product.LostReason = param.LostReason
	product.LostWay = param.LostWay
	product.PreOwner = product.CurrentOwner
	product.CurrentOwner = common.SYSTEM
	// modify status 灭尸
	product.Status = common.STATUS["LOST"]
	// marshal product into state
	jsonProduct, err := json.Marshal(product)
	if err != nil {
		log.Logger.Error("goLost -- marshal product err:" + err.Error() + " 	prodocut:" + param.ProductId)
		vChan.Status = false
		vChan.ErrorCode = common.ERR["CHAINERR"]

		return
	}
	err = stub.PutState(common.PRODUCT_INFO+common.ULINE+param.ProductId, jsonProduct)
	if err != nil {
		log.Logger.Error("goLost -- PutState product err:" + err.Error() + " 	prodocut:" + param.ProductId)
		vChan.Status = false
		vChan.ErrorCode = common.ERR["CHAINERR"]

		return
	}

	// ASSET CHANGE OWNER === START
	changeOwner := module.ChangeAssetOwner{}
	changeOwner.PreOwner = product.PreOwner
	changeOwner.CurrentOwner = product.CurrentOwner
	changeOwner.ProductId = param.ProductId
	changeOwner.Operation = param.Operation
	changeOwner.Operator = param.Operator
	time, err := stub.GetTxTimestamp()
	if err != nil {
		log.Logger.Error("goLost -- goButcher change owner get time:" + err.Error() + "	productid:" + param.ProductId)
		vChan.Status = false
		vChan.ErrorCode = common.ERR["CHAINERR"]

		return
	}
	changeOwner.OperateTime = uint64(time.GetSeconds())
	jsonchangeOwnerBytes, err := json.Marshal(changeOwner)
	err = stub.PutState(common.PRODUCT_TRANSFER+common.ULINE+param.ProductId, jsonchangeOwnerBytes)

	if err != nil {
		log.Logger.Error("goLost -- PutState change owner:" + err.Error() + "	productid:" + param.ProductId)
		vChan.Status = false
		vChan.ErrorCode = common.ERR["CHAINERR"]

		return
	}
	// ASSET CHANGE OWNER === END
	vChan.Status = true
	vChan.ErrorCode = common.ERR["NONE"]

	return
}
