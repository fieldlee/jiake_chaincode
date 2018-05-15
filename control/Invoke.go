package control

import (
	"encoding/json"
	"jiake_chaincode/log"
	"strings"
	"jiake_chaincode/service"
	"jiake_chaincode/module"
	"github.com/hyperledger/fabric/core/chaincode/shim"
	"github.com/hyperledger/fabric/protos/peer"
)

type ReturnInfo struct {
	Status : bool,
	Info:string
}

func (t *ProductTrace) Invoke(stub shim.ChaincodeStubInterface) peer.Response {
	log.Logger.Info("Invoke")

	funcation, args := stub.GetFunctionAndParameters()
	lowFuncation := strings.ToLower(funcation)

	if lowFuncation == "register" { //入栏
		return t.Register(stub,args)
	} else if lowFuncation == "feed" { //喂养
		return t.Feed(stub,args)
	} else if lowFuncation == "vaccine" { //防疫
		return t.Vaccine(stub,args)
	} else if lowFuncation == "output" { //出栏
		return t.Output(stub,args)
	} else if lowFuncation == "exam" { //检疫
		return t.Exam(stub,args)
	} else if lowFuncation == "lost" { //灭尸
		return t.Lost(stub,args)
	} else if lowFuncation == "butcher" { //屠宰
		return t.Butcher(stub,args)
	} else if lowFuncation == "querybyproduct" { //查询资产ID
		return t.QueryByProduct(stub,args)
	} else if lowFuncation == "queryhistorybyproduct" { //查询批次ID
		return t.QueryHistoryByProduct(stub,args)
	} else if lowFuncation == "" { //查询交易ID

	} else if lowFuncation == "QueryBatchByProduct" { //查询批次ID
		return t.QueryBatchByProduct(stub,args)
	}
	return shim.Error("Invalid invoke function name. " + funcation)
}

/**入栏**/
func (t *ProductTrace) Register(stub shim.ChaincodeStubInterface, args []string) peer.Response {
	log.Logger.Info("##############调用Register接口开始###############")
	returnInfo := ReturnInfo{}
	if len(args)>=1 {
		var paramList []module.RegitserParam
		err := json.Unmarshal([]byte(args),&paramList)
		if err != nil {
			returnInfo.Status = false
			returnInfo.Info = ""
		}else{
			return service.Register(stub,paramList)
		}
	}else{
		returnInfo.Status = false
		returnInfo.Info = ""
	}
	jsonreturn , err := json.Marshal(returnInfo)
	return shim.Success(string(jsonreturn[:]))
}

/**喂养**/
func (t *ProductTrace) Feed(stub shim.ChaincodeStubInterface, args []string) peer.Response {
	log.Logger.Info("##############调用Feed接口开始###############")
	returnInfo := ReturnInfo{}
	if len(args)>=1 {
		var paramList []module.FeedParam
		err := json.Unmarshal([]byte(args),&paramList)
		if err != nil {
			returnInfo.Status = false
			returnInfo.Info = ""
		}else{
			return service.Feed(stub,paramList)
		}
		
	}else{
		returnInfo.Status = false
		returnInfo.Info = ""
	}
	jsonreturn , err := json.Marshal(returnInfo)
	return shim.Success(string(jsonreturn[:]))
}

/**防疫**/
func (t *ProductTrace) Vaccine(stub shim.ChaincodeStubInterface, args []string) peer.Response {
	log.Logger.Info("##############调用Vaccine接口开始###############")
	returnInfo := ReturnInfo{}
	if len(args)>=1 {
		var paramList []module.VaccineParam
		err := json.Unmarshal([]byte(args),&paramList)
		if err != nil {
			returnInfo.Status = false
			returnInfo.Info = ""
		}else{
			return service.Vaccine(stub,paramList)
		}
		
	}else{
		returnInfo.Status = false
		returnInfo.Info = ""
	}
	jsonreturn , err := json.Marshal(returnInfo)
	return shim.Success(string(jsonreturn[:]))
}

/**出栏**/
func (t *ProductTrace) Output(stub shim.ChaincodeStubInterface, args []string) peer.Response {
	log.Logger.Info("##############调用Output接口开始###############")
	returnInfo := ReturnInfo{}
	if len(args)>=1 {
		var paramList []module.OutputParam
		err := json.Unmarshal([]byte(args),&paramList)
		if err != nil {
			returnInfo.Status = false
			returnInfo.Info = ""
		}else{
			return service.Output(stub,paramList)
		}
		
	}else{
		returnInfo.Status = false
		returnInfo.Info = ""
	}
	jsonreturn , err := json.Marshal(returnInfo)
	return shim.Success(string(jsonreturn[:]))
}

/**检疫**/
func (t *ProductTrace) Exam(stub shim.ChaincodeStubInterface, args []string) peer.Response {
	log.Logger.Info("##############调用Exam接口开始###############")
	returnInfo := ReturnInfo{}
	if len(args)>=1 {
		var paramList []module.ExamParam
		err := json.Unmarshal([]byte(args),&paramList)
		if err != nil {
			returnInfo.Status = false
			returnInfo.Info = ""
		}else{
			return service.Exam(stub,paramList)
		}
		
	}else{
		returnInfo.Status = false
		returnInfo.Info = ""
	}
	jsonreturn , err := json.Marshal(returnInfo)
	return shim.Success(string(jsonreturn[:]))
}

/**屠宰**/
func (t *ProductTrace) Butcher(stub shim.ChaincodeStubInterface, args []string) peer.Response {
	log.Logger.Info("##############调用Butcher接口开始###############")
	returnInfo := ReturnInfo{}
	if len(args)>=1 {
		var paramList []module.ButcherParam
		err := json.Unmarshal([]byte(args),&paramList)
		if err != nil {
			returnInfo.Status = false
			returnInfo.Info = ""
		}else{
			return service.Butcher(stub,paramList)
		}
		
	}else{
		returnInfo.Status = false
		returnInfo.Info = ""
	}
	jsonreturn , err := json.Marshal(returnInfo)
	return shim.Success(string(jsonreturn[:]))
}

/**灭尸**/
func (t *ProductTrace) Lost(stub shim.ChaincodeStubInterface, args []string) peer.Response {
	log.Logger.Info("##############调用Lost接口开始###############")
	returnInfo := ReturnInfo{}
	if len(args)>=1 {
		var paramList []module.LostParam
		err := json.Unmarshal([]byte(args),&paramList)
		if err != nil {
			returnInfo.Status = false
			returnInfo.Info = ""
		}else{
			return service.Lost(stub,paramList)
		}
		
	}else{
		returnInfo.Status = false
		returnInfo.Info = ""
	}
	jsonreturn , err := json.Marshal(returnInfo)
	return shim.Success(string(jsonreturn[:]))
}

/**查询产品**/
func (t *ProductTrace) QueryByProduct(stub shim.ChaincodeStubInterface,args []string) peer.Response{
	log.Logger.Info("##############调用Queryproduct接口开始###############")
	returnInfo := ReturnInfo{}
	if len(args)>=1 {
		var param module.QueryParam
		err := json.Unmarshal([]byte(args[0]),&param)
		if err != nil {
			returnInfo.Status = false
			returnInfo.Info = ""
		}else{
			return service.QueryByProduct(stub,param)
		}
		
	}else{
		returnInfo.Status = false
		returnInfo.Info = ""
	}
	jsonreturn , err := json.Marshal(returnInfo)
	return shim.Success(string(jsonreturn[:]))
} 

/**查询历史产品**/
func (t *ProductTrace) QueryHistoryByProduct(stub shim.ChaincodeStubInterface,args []string) peer.Response{
	log.Logger.Info("##############调用QueryHistoryproduct接口开始###############")
	returnInfo := ReturnInfo{}
	if len(args)>=1 {
		var param module.QueryParam
		err := json.Unmarshal([]byte(args[0]),&param)
		if err != nil {
			returnInfo.Status = false
			returnInfo.Info = ""
		}else{
			return service.QueryHistoryByProduct(stub,param)
		}
		
	}else{
		returnInfo.Status = false
		returnInfo.Info = ""
	}
	jsonreturn , err := json.Marshal(returnInfo)
	return shim.Success(string(jsonreturn[:]))
}

/**查询批次产品**/
func (t *ProductTrace) QueryBatchByProduct(stub shim.ChaincodeStubInterface,args []string) peer.Response{
	log.Logger.Info("##############调用QueryBatchByProduct接口开始###############")
	returnInfo := ReturnInfo{}
	if len(args)>=1 {
		var param module.BatchParam
		err := json.Unmarshal([]byte(args[0]),&param)
		if err != nil {
			returnInfo.Status = false
			returnInfo.Info = ""
		}else{
			return service.QueryBatchByProduct(stub,param)
		}
		
	}else{
		returnInfo.Status = false
		returnInfo.Info = ""
	}
	jsonreturn , err := json.Marshal(returnInfo)
	return shim.Success(string(jsonreturn[:]))
}

/**查询交易产品**/
func (t *ProductTrace) QueryByTX(stub shim.ChaincodeStubInterface,args []string) peer.Response{
	log.Logger.Info("##############QueryByTX###############")
	returnInfo := ReturnInfo{}
	if len(args)>=1 {
		var param module.QueryTxParam
		err := json.Unmarshal([]byte(args[0]),&param)
		if err != nil {
			returnInfo.Status = false
			returnInfo.Info = ""
		}else{
			return service.QueryByTX(stub,param)
		}
		
	}else{
		returnInfo.Status = false
		returnInfo.Info = ""
	}
	jsonreturn , err := json.Marshal(returnInfo)
	return shim.Success(string(jsonreturn[:]))
}