package control

import (
	"encoding/json"
	"jiakechaincode/log"
	"jiakechaincode/module"
	"jiakechaincode/service"
	"strings"

	"github.com/hyperledger/fabric/core/chaincode/shim"
	"github.com/hyperledger/fabric/protos/peer"
)

func (t *ProductTrace) Invoke(stub shim.ChaincodeStubInterface) peer.Response {
	log.Logger.Info("Invoke")

	funcation, args := stub.GetFunctionAndParameters()
	lowFuncation := strings.ToLower(funcation)

	if lowFuncation == "register" { //入栏
		return t.Register(stub, args)
	} else if lowFuncation == "feed" { //喂养
		return t.Feed(stub, args)
	} else if lowFuncation == "vaccine" { //防疫
		return t.Vaccine(stub, args)
	} else if lowFuncation == "save" { //救治
		return t.Save(stub, args)
	} else if lowFuncation == "output" { //出栏
		return t.Output(stub, args)
	} else if lowFuncation == "exam" { //检疫
		return t.Exam(stub, args)
	} else if lowFuncation == "lost" { //灭尸
		return t.Lost(stub, args)
	} else if lowFuncation == "butcher" { //屠宰
		return t.Butcher(stub, args)
	} else if lowFuncation == "querybyproduct" { //查询资产ID
		return t.QueryByProduct(stub, args)
	} else if lowFuncation == "queryhistorybyproduct" { //查询批次ID
		return t.QueryHistoryByProduct(stub, args)
	} else if lowFuncation == "querybytx" { //查询交易ID
		return t.QueryByTX(stub, args)
	} else if lowFuncation == "querybatchbyproduct" { //查询批次ID
		return t.QueryBatchByProduct(stub, args)
	}
	return shim.Error("Invalid invoke function name. " + funcation)
}

/**入栏**/
func (t *ProductTrace) Register(stub shim.ChaincodeStubInterface, args []string) peer.Response {
	log.Logger.Info("##############调用Register接口开始###############")
	returnInfo := module.ReturnInfo{}
	if len(args) >= 1 {
		var paramList []module.RegitserParam
		err := json.Unmarshal([]byte(args[0]), &paramList)
		if err != nil {
			log.Logger.Error("Register:err" + err.Error())
			returnInfo.Success = false
			returnInfo.Info = err.Error()
		} else {
			return service.Register(stub, paramList)
		}
	} else {
		log.Logger.Error("Register:参数不对，请核实参数信息。")
		returnInfo.Success = false
		returnInfo.Info = "参数不对，请核实参数信息"
	}
	jsonreturn, err := json.Marshal(returnInfo)
	if err != nil {
		return shim.Error("err:" + err.Error())
	}
	return shim.Success(jsonreturn)
}

/**喂养**/
func (t *ProductTrace) Feed(stub shim.ChaincodeStubInterface, args []string) peer.Response {
	log.Logger.Info("##############调用Feed接口开始###############")
	returnInfo := module.ReturnInfo{}
	if len(args) >= 1 {
		var paramList []module.FeedParam
		err := json.Unmarshal([]byte(args[0]), &paramList)
		if err != nil {
			log.Logger.Error("Feed:err" + err.Error())
			returnInfo.Success = false
			returnInfo.Info = err.Error()
		} else {
			return service.Feed(stub, paramList)
		}

	} else {
		log.Logger.Error("Feed:参数不对，请核实参数信息。")
		returnInfo.Success = false
		returnInfo.Info = "参数不对，请核实参数信息"
	}
	jsonreturn, err := json.Marshal(returnInfo)
	if err != nil {
		return shim.Error("err:" + err.Error())
	}
	return shim.Success(jsonreturn)
}

/**防疫**/
func (t *ProductTrace) Vaccine(stub shim.ChaincodeStubInterface, args []string) peer.Response {
	log.Logger.Info("##############调用Vaccine接口开始###############")
	returnInfo := module.ReturnInfo{}
	if len(args) >= 1 {
		var paramList []module.VaccineParam
		err := json.Unmarshal([]byte(args[0]), &paramList)
		if err != nil {
			log.Logger.Error("Vaccine:err" + err.Error())
			returnInfo.Success = false
			returnInfo.Info = err.Error()
		} else {
			return service.Vaccine(stub, paramList)
		}

	} else {
		log.Logger.Error("Vaccine:参数不对，请核实参数信息。")
		returnInfo.Success = false
		returnInfo.Info = "参数不对，请核实参数信息"
	}
	jsonreturn, err := json.Marshal(returnInfo)
	if err != nil {
		return shim.Error("err:" + err.Error())
	}
	return shim.Success(jsonreturn)
}

/**出栏**/
func (t *ProductTrace) Output(stub shim.ChaincodeStubInterface, args []string) peer.Response {
	log.Logger.Info("##############调用Output接口开始###############")
	returnInfo := module.ReturnInfo{}
	if len(args) >= 1 {
		var paramList []module.OutputParam
		err := json.Unmarshal([]byte(args[0]), &paramList)
		if err != nil {
			log.Logger.Error("Output:err" + err.Error())
			returnInfo.Success = false
			returnInfo.Info = err.Error()
		} else {
			return service.Output(stub, paramList)
		}

	} else {
		log.Logger.Error("Output:参数不对，请核实参数信息。")
		returnInfo.Success = false
		returnInfo.Info = "参数不对，请核实参数信息"
	}
	jsonreturn, err := json.Marshal(returnInfo)
	if err != nil {
		return shim.Error("err:" + err.Error())
	}
	return shim.Success(jsonreturn)
}

/**检疫**/
func (t *ProductTrace) Exam(stub shim.ChaincodeStubInterface, args []string) peer.Response {
	log.Logger.Info("##############调用Exam接口开始###############")
	returnInfo := module.ReturnInfo{}
	if len(args) >= 1 {
		var paramList []module.ExamParam
		err := json.Unmarshal([]byte(args[0]), &paramList)
		if err != nil {
			log.Logger.Error("Exam:" + err.Error())
			returnInfo.Success = false
			returnInfo.Info = err.Error()
		} else {
			return service.Exam(stub, paramList)
		}

	} else {
		log.Logger.Error("Exam:参数不对，请核实参数信息。")
		returnInfo.Success = false
		returnInfo.Info = "参数不对，请核实参数信息"
	}
	jsonreturn, err := json.Marshal(returnInfo)
	if err != nil {
		return shim.Error("err:" + err.Error())
	}
	return shim.Success(jsonreturn)
}

/**救治**/
func (t *ProductTrace) Save(stub shim.ChaincodeStubInterface, args []string) peer.Response {
	log.Logger.Info("##############调用Save接口开始###############")
	returnInfo := module.ReturnInfo{}
	if len(args) >= 1 {
		var paramList []module.SaveParam
		err := json.Unmarshal([]byte(args[0]), &paramList)
		if err != nil {
			log.Logger.Error("Save:" + err.Error())
			returnInfo.Success = false
			returnInfo.Info = err.Error()
		} else {
			return service.Save(stub, paramList)
		}

	} else {
		log.Logger.Error("Exam:参数不对，请核实参数信息。")
		returnInfo.Success = false
		returnInfo.Info = "参数不对，请核实参数信息"
	}
	jsonreturn, err := json.Marshal(returnInfo)
	if err != nil {
		return shim.Error("err:" + err.Error())
	}
	return shim.Success(jsonreturn)
}

/**屠宰**/
func (t *ProductTrace) Butcher(stub shim.ChaincodeStubInterface, args []string) peer.Response {
	log.Logger.Info("##############调用Butcher接口开始###############")
	returnInfo := module.ReturnInfo{}
	if len(args) >= 1 {
		var paramList []module.ButcherParam
		err := json.Unmarshal([]byte(args[0]), &paramList)
		if err != nil {
			log.Logger.Error("Butcher:" + err.Error())
			returnInfo.Success = false
			returnInfo.Info = err.Error()
		} else {
			return service.Butcher(stub, paramList)
		}

	} else {
		log.Logger.Error("Butcher:参数不对，请核实参数信息。")
		returnInfo.Success = false
		returnInfo.Info = "参数不对，请核实参数信息"
	}
	jsonreturn, err := json.Marshal(returnInfo)
	if err != nil {
		return shim.Error("err:" + err.Error())
	}
	return shim.Success(jsonreturn)
}

/**灭尸**/
func (t *ProductTrace) Lost(stub shim.ChaincodeStubInterface, args []string) peer.Response {
	log.Logger.Info("##############调用Lost接口开始###############")
	returnInfo := module.ReturnInfo{}
	if len(args) >= 1 {
		var paramList []module.DestroyParam
		err := json.Unmarshal([]byte(args[0]), &paramList)
		if err != nil {
			log.Logger.Error("Lost:err:" + err.Error())
			returnInfo.Success = false
			returnInfo.Info = err.Error()
		} else {
			return service.Lost(stub, paramList)
		}

	} else {
		log.Logger.Error("Lost:参数不对，请核实参数信息。")
		returnInfo.Success = false
		returnInfo.Info = "参数不对，请核实参数信息"
	}
	jsonreturn, err := json.Marshal(returnInfo)
	if err != nil {
		return shim.Error("err:" + err.Error())
	}
	return shim.Success(jsonreturn)
}

/**查询产品**/
func (t *ProductTrace) QueryByProduct(stub shim.ChaincodeStubInterface, args []string) peer.Response {
	log.Logger.Info("##############调用Queryproduct接口开始###############")
	returnInfo := module.ReturnInfo{}
	if len(args) >= 1 {
		var param module.QueryParam
		err := json.Unmarshal([]byte(args[0]), &param)
		if err != nil {
			log.Logger.Error("QueryByProduct:err。" + err.Error())
			returnInfo.Success = false
			returnInfo.Info = err.Error()
		} else {
			return service.QueryByProduct(stub, param)
		}

	} else {
		log.Logger.Error("QueryByProduct:参数不对，请核实参数信息。")
		returnInfo.Success = false
		returnInfo.Info = "参数不对，请核实参数信息"
	}
	jsonreturn, err := json.Marshal(returnInfo)
	if err != nil {
		return shim.Error("err:" + err.Error())
	}
	return shim.Success(jsonreturn)
}

/**查询历史产品**/
func (t *ProductTrace) QueryHistoryByProduct(stub shim.ChaincodeStubInterface, args []string) peer.Response {
	log.Logger.Info("##############调用QueryHistoryproduct接口开始###############")
	returnInfo := module.ReturnInfo{}
	if len(args) >= 1 {
		var param module.QueryParam
		err := json.Unmarshal([]byte(args[0]), &param)
		if err != nil {
			log.Logger.Error("QueryHistoryByProduct:err:" + err.Error())
			returnInfo.Success = false
			returnInfo.Info = err.Error()
		} else {
			return service.QueryHistoryByProduct(stub, param)
		}

	} else {
		log.Logger.Error("QueryByProduct:参数不对，请核实参数信息。")
		returnInfo.Success = false
		returnInfo.Info = "参数不对，请核实参数信息"
	}
	jsonreturn, err := json.Marshal(returnInfo)
	if err != nil {
		return shim.Error("err:" + err.Error())
	}
	return shim.Success(jsonreturn)
}

/**查询批次产品**/
func (t *ProductTrace) QueryBatchByProduct(stub shim.ChaincodeStubInterface, args []string) peer.Response {
	log.Logger.Info("##############调用QueryBatchByProduct接口开始###############")
	returnInfo := module.ReturnInfo{}
	if len(args) >= 1 {
		var param module.BatchParam
		err := json.Unmarshal([]byte(args[0]), &param)
		if err != nil {
			log.Logger.Error("QueryBatchByProduct:err:" + err.Error())
			returnInfo.Success = false
			returnInfo.Info = err.Error()
		} else {
			return service.QueryBatchByProduct(stub, param)
		}

	} else {
		log.Logger.Error("QueryBatchByProduct:参数不对，请核实参数信息。")
		returnInfo.Success = false
		returnInfo.Info = "参数不对，请核实参数信息"
	}
	jsonreturn, err := json.Marshal(returnInfo)
	if err != nil {
		return shim.Error("err:" + err.Error())
	}
	return shim.Success(jsonreturn)
}

/**查询交易产品**/
func (t *ProductTrace) QueryByTX(stub shim.ChaincodeStubInterface, args []string) peer.Response {
	log.Logger.Info("##############QueryByTX###############")
	returnInfo := module.ReturnInfo{}
	if len(args) >= 1 {
		var param module.QueryTxParam
		err := json.Unmarshal([]byte(args[0]), &param)
		if err != nil {
			log.Logger.Error("QueryByTX:err:" + err.Error())
			returnInfo.Success = false
			returnInfo.Info = err.Error()
		} else {
			return service.QueryByTX(stub, param)
		}

	} else {
		log.Logger.Error("QueryByTX:参数不对，请核实参数信息。")
		returnInfo.Success = false
		returnInfo.Info = "参数不对，请核实参数信息"
	}
	jsonreturn, err := json.Marshal(returnInfo)
	if err != nil {
		return shim.Error("err:" + err.Error())
	}
	return shim.Success(jsonreturn)
}
