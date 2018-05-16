package service

import (
	"encoding/json"
	"jiakechaincode/module"
	"jiakechaincode/log"
	"jiakechaincode/common"
	"sync"

	"github.com/hyperledger/fabric/core/chaincode/shim"
	"github.com/hyperledger/fabric/protos/peer"
)

var wg sync.WaitGroup

type ChanInfo struct{
	ProductId:string,
	Status:bool,
	ErrorCode:string
}

type ReturnErrorInfo struct {
	Status : bool,
	ErrorList :[]ChanInfo
}

/**产品注册信息**/
func Register(stub shim.ChaincodeStubInterface, paramList []module.RegitserParam) peer.Response {
	wg.Add(len(paramList)) //添加队列数
	registerChan := make(chan ChanInfo, len(paramList))
	for _, v := range paramList {
		go goRegister(stub, v,registerChan)
	}
// 	获得chan 返回的值
	returnError := ReturnErrorInfo{}

	for i,_ := range paramList {
		tChan := ChanInfo{}
		tChan <- registerChan //get Channel return value
		if tChan.Status == false {
			returnError.ErrorList = append(returnError.ErrorList,tChan) 
		}
	}
	close(registerChan)
	wg.Wait()

	if len(returnError.ErrorList)>=1 {
		returnError.Status = false
	}else{
		returnError.Status = true
	}

	jsonReturn,err := json.Marshal(returnError)
	return shim.Success(string(jsonReturn[:]))
}


/**喂养上链**/
func Feed(stub shim.ChaincodeStubInterface, paramList []module.FeedParam) peer.Response {
	wg.Add(len(paramList)) //添加队列数
	feedChan := make(chan ChanInfo, len(paramList))
	for _, v := range paramList {
		go goFeed(stub, v,feedChan)
	}
// 	获得chan 返回的值
	returnError := ReturnErrorInfo{}

	for i,_ := range paramList {
		tChan := ChanInfo{}
		tChan <- feedChan //get Channel return value
		if tChan.Status == false {
			returnError.ErrorList = append(returnError.ErrorList,tChan) 
		}
	}
	close(feedChan)
	wg.Wait()

	if len(returnError.ErrorList)>=1 {
		returnError.Status = false
	}else{
		returnError.Status = true
	}
	jsonReturn,err := json.Marshal(returnError)
	return shim.Success(string(jsonReturn[:]))
}

/**防疫信息上链**/
func Vaccine(stub shim.ChaincodeStubInterface, paramList []module.VaccineParam) peer.Response {
	wg.Add(len(paramList)) //添加队列数
	vaccineChan := make(chan ChanInfo, len(paramList))
	for _, v := range paramList {
		go goVaccine(stub, v,vaccineChan)
	}
// 	获得chan 返回的值
	returnError := ReturnErrorInfo{}

	for i,_ := range paramList {
		tChan := ChanInfo{}
		tChan <- vaccineChan //get Channel return value
		if tChan.Status == false {
			returnError.ErrorList = append(returnError.ErrorList,tChan) 
		}
	}
	close(vaccineChan)
	wg.Wait()

	if len(returnError.ErrorList)>=1 {
		returnError.Status = false
	}else{
		returnError.Status = true
	}
	jsonReturn,err := json.Marshal(returnError)
	return shim.Success(string(jsonReturn[:]))
}

/**出栏操作**/
func Output(stub shim.ChaincodeStubInterface,paramList []module.OutputParam) peer.Response{
	wg.Add(len(paramList)) //添加队列数
	outputChan := make(chan ChanInfo, len(paramList))
	for _, v := range paramList {
		go goOutput(stub, v,outputChan)
	}
// 	获得chan 返回的值
	returnError := ReturnErrorInfo{}
	for i,_ := range paramList {
		tChan := ChanInfo{}
		tChan <- outputChan //get Channel return value
		if tChan.Status == false {
			returnError.ErrorList = append(returnError.ErrorList,tChan) 
		}
	}
	close(outputChan)
	wg.Wait()

	if len(returnError.ErrorList)>=1 {
		returnError.Status = false
	}else{
		returnError.Status = true
	}
	jsonReturn,err := json.Marshal(returnError)
	return shim.Success(string(jsonReturn[:]))
}

/**检疫**/
func Exam(stub shim.ChaincodeStubInterface,paramList []module.ExamParam)peer.Response  {
	wg.Add(len(paramList)) //添加队列数
	examChan := make(chan ChanInfo, len(paramList))
	for _, v := range paramList {
		go goExam(stub, v,examChan)
	}
// 	获得chan 返回的值
	returnError := ReturnErrorInfo{}
	for i,_ := range paramList {
		tChan := ChanInfo{}
		tChan <- examChan //get Channel return value
		if tChan.Status == false {
			returnError.ErrorList = append(returnError.ErrorList,tChan) 
		}
	}
	close(examChan)
	wg.Wait()
	if len(returnError.ErrorList)>=1 {
		returnError.Status = false
	}else{
		returnError.Status = true
	}
	jsonReturn,err := json.Marshal(returnError)
	return shim.Success(string(jsonReturn[:]))
}

/**检疫**/
func Exam(stub shim.ChaincodeStubInterface,paramList []module.SaveParam)peer.Response  {
	wg.Add(len(paramList)) //添加队列数
	saveChan := make(chan ChanInfo, len(paramList))
	for _, v := range paramList {
		go goSave(stub, v,saveChan)
	}
// 	获得chan 返回的值
	returnError := ReturnErrorInfo{}
	for i,_ := range paramList {
		tChan := ChanInfo{}
		tChan <- saveChan //get Channel return value
		if tChan.Status == false {
			returnError.ErrorList = append(returnError.ErrorList,tChan) 
		}
	}
	close(saveChan)
	wg.Wait()
	if len(returnError.ErrorList)>=1 {
		returnError.Status = false
	}else{
		returnError.Status = true
	}
	jsonReturn,err := json.Marshal(returnError)
	return shim.Success(string(jsonReturn[:]))
}

/**屠宰**/
func Butcher(stub shim.ChaincodeStubInterface,paramList []module.ButcherParam)peer.Response{
	wg.Add(len(paramList)) //添加队列数
	butcherChan := make(chan ChanInfo, len(paramList))
	for _, v := range paramList {
		go goButcher(stub, v,butcherChan)
	}
// 	获得chan 返回的值
	returnError := ReturnErrorInfo{}
	for i,_ := range paramList {
		tChan := ChanInfo{}
		tChan <- butcherChan //get Channel return value
		if tChan.Status == false {
			returnError.ErrorList = append(returnError.ErrorList,tChan) 
		}
	}
	close(butcherChan)
	wg.Wait()
	if len(returnError.ErrorList)>=1 {
		returnError.Status = false
	}else{
		returnError.Status = true
	}
	jsonReturn,err := json.Marshal(returnError)
	return shim.Success(string(jsonReturn[:]))
}

/**灭尸**/
func Lost(stub shim.ChaincodeStubInterface,paramList []module.LostParam)peer.Response{
	wg.Add(len(paramList)) //添加队列数
	lostChan := make(chan ChanInfo, len(paramList))
	for _, v := range paramList {
		go goLost(stub, v,lostChan)
	}
	// 	获得chan 返回的值
	returnError := ReturnErrorInfo{}
	for i,_ := range paramList {
		tChan := ChanInfo{}
		tChan <- lostChan //get Channel return value
		if tChan.Status == false {
			returnError.ErrorList = append(returnError.ErrorList,tChan) 
		}
	}
	close(lostChan)
	wg.Wait()
	if len(returnError.ErrorList)>=1 {
		returnError.Status = false
	}else{
		returnError.Status = true
	}
	jsonReturn,err := json.Marshal(returnError)
	return shim.Success(string(jsonReturn[:]))
}

/**查询**/
func QueryByProduct(stub shim.ChaincodeStubInterface,param module.QueryParam)peer.Response{

	jsonBytes , err := stub.GetState(common.PRODUCT_INFO+common.ULINE+param.ProductId)
	if err != nil {
		log.Logger.Error("QueryByProduct err:"+err.Error())
		return shim.Error(err.Error())
	}
	return shim.Success(string(jsonBytes[:]))
}

/**查询历史**/
func QueryHistoryByProduct(stub shim.ChaincodeStubInterface,param module.QueryParam)peer.Response{
	historys, err := stub.GetHistoryForKey(common.PRODUCT_INFO+common.ULINE+param.ProductId)
	if err != nil {
		log.Logger.Error("QueryHistoryByProduct err:"+err.Error())
		return shim.Error(err.Error())
	}
	defer historys.Close()
	results := make([]module.Product, 0)
	for historys.HasNext() {
		result, err := historys.Next()
		if err != nil {
			return shim.Error(err.Error())
		}
		product := module.Product{}

		err = json.Unmarshal(result.Value, &product)
		if err != nil {
			return shim.Error(err.Error())
		} else {
			results = append(results, product)
		}
	}
	resultsJSON, err := json.Marshal(results)
	if err != nil {
		return shim.Error(err.Error())
	}
	return shim.Success(string(resultsJSON[:]))
}

/**查询批次**/
func QueryBatchByProduct(stub shim.ChaincodeStubInterface,param module.BatchParam)peer.Response{
	queryString := fmt.Sprintf("{\"selector\": {\"_id\": {\"$regex\": \"%s\"},\"batchNumber\":\"%s\"},\"limit\":\"%d\"}", common.PRODUCT_INFO,param.BatchNumber,5000)
	queryResults , err := stub.GetQueryResult(queryString)
	if err != nil {
		return shim.Error(err.Error())
	}
	defer queryResults.Close()
	results := make([]module.Product, 0)
	for queryResults.HasNext(){
		result , err := queryResults.Next()
		if err != nil {
			return shim.Error(err.Error())
		}
		product := module.Product{}
		err = json.Unmarshal(result.Value, &product)
		if err != nil {
			return shim.Error(err.Error())
		} else {
			results = append(results, product)
		}
	}
	resultsJSON, err := json.Marshal(results)
	if err != nil {
		return shim.Error(err.Error())
	}
	return shim.Success(string(resultsJSON[:]))
}

/**查询交易**/
func QueryByTX(stub shim.ChaincodeStubInterface,param module.QueryTxParam)peer.Response{
	queryString := fmt.Sprintf("{\"selector\": {\"_id\": {\"$regex\": \"%s\"},\"$or\":[{\"txId\":\"%s\"},{\"outputTxId\":\"%s\"},{\"lostTxId\":\"%s\"},{\"examTxId\":\"%s\"},{\"butcherTxId\":\"%s\"},{\"feedList.txId\":\"%s\"},{\"vaccineList.txId\":\"%s\"},{\"saveList.txId\":\"%s\"}]},\"limit\":\"%d\"}", common.PRODUCT_INFO,param.TxId,param.TxId,param.TxId,param.TxId,param.TxId,param.TxId,param.TxId,param.TxId,5000)

	queryResults , err := stub.GetQueryResult(queryString)
	
	if err != nil {
		return shim.Error(err.Error())
	}
	defer queryResults.Close()
	results := make([]module.Product, 0)
	for queryResults.HasNext(){
		result , err := queryResults.Next()
		if err != nil {
			return shim.Error(err.Error())
		}
		product := module.Product{}
		err = json.Unmarshal(result.Value, &product)
		if err != nil {
			return shim.Error(err.Error())
		} else {
			results = append(results, product)
		}
	}
	resultsJSON, err := json.Marshal(results)
	if err != nil {
		return shim.Error(err.Error())
	}
	return shim.Success(string(resultsJSON[:]))
}