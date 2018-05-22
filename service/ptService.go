package service

import (
	"encoding/json"
	"fmt"
	"jiakechaincode/common"
	"jiakechaincode/log"
	"jiakechaincode/module"
	"strconv"
	"sync"

	"github.com/hyperledger/fabric/core/chaincode/shim"
	"github.com/hyperledger/fabric/protos/peer"
)

var wg sync.WaitGroup

type ChanInfo struct {
	ProductId string `json:"productId"`
	Status    bool   `json:"status"`
	ErrorCode string `json:"errorCode"`
}

type ReturnErrorInfo struct {
	Status    bool       `json:"status"`
	ErrorList []ChanInfo `json:errorList`
}

/**产品注册信息**/
func Register(stub shim.ChaincodeStubInterface, paramList []module.RegitserParam) peer.Response {
	wg.Add(len(paramList)) //添加队列数
	registerChan := make(chan ChanInfo, len(paramList))
	for i, v := range paramList {
		log.Logger.Info("Register --- send :" + strconv.Itoa(i))
		log.Logger.Info("Register --- send :" + v.ProductId)
		go goRegister(stub, v, registerChan)
	}
	// 	获得chan 返回的值
	returnError := ReturnErrorInfo{}

	for j, _ := range paramList {
		log.Logger.Info("Register chan --- range :" + strconv.Itoa(j))
		tChan := <-registerChan //get Channel return value
		log.Logger.Info("Register --- RECIVED :" + tChan.ProductId)
		if tChan.Status == false {
			returnError.ErrorList = append(returnError.ErrorList, tChan)
		}
	}
	close(registerChan)
	wg.Wait()

	if len(returnError.ErrorList) >= 1 {
		returnError.Status = false
	} else {
		returnError.Status = true
	}

	jsonReturn, _ := json.Marshal(returnError)
	log.Logger.Info(string(jsonReturn[:]))
	return shim.Success(jsonReturn)
}

/**喂养上链**/
func Feed(stub shim.ChaincodeStubInterface, paramList []module.FeedParam) peer.Response {
	wg.Add(len(paramList)) //添加队列数
	feedChan := make(chan ChanInfo, len(paramList))
	for i, v := range paramList {
		log.Logger.Info("Feed --range:" + strconv.Itoa(i))
		go goFeed(stub, v, feedChan)
	}
	// 	获得chan 返回的值
	returnError := ReturnErrorInfo{}

	for j, _ := range paramList {
		log.Logger.Info("Feed chan --- range :" + strconv.Itoa(j))
		tChan := <-feedChan //get Channel return value
		if tChan.Status == false {
			returnError.ErrorList = append(returnError.ErrorList, tChan)
		}
	}
	close(feedChan)
	wg.Wait()

	if len(returnError.ErrorList) >= 1 {
		returnError.Status = false
	} else {
		returnError.Status = true
	}
	jsonReturn, _ := json.Marshal(returnError)
	return shim.Success(jsonReturn)
}

/**防疫信息上链**/
func Vaccine(stub shim.ChaincodeStubInterface, paramList []module.VaccineParam) peer.Response {
	wg.Add(len(paramList)) //添加队列数
	vaccineChan := make(chan ChanInfo, len(paramList))
	for i, v := range paramList {
		log.Logger.Info("Vaccine  ---range:" + strconv.Itoa(i))
		go goVaccine(stub, v, vaccineChan)
	}
	// 	获得chan 返回的值
	returnError := ReturnErrorInfo{}

	for j, _ := range paramList {
		log.Logger.Info("goVaccine chan --- range :" + strconv.Itoa(j))
		tChan := <-vaccineChan //get Channel return value
		if tChan.Status == false {
			returnError.ErrorList = append(returnError.ErrorList, tChan)
		}
	}
	close(vaccineChan)
	wg.Wait()

	if len(returnError.ErrorList) >= 1 {
		returnError.Status = false
	} else {
		returnError.Status = true
	}
	jsonReturn, _ := json.Marshal(returnError)
	return shim.Success(jsonReturn)
}

/**出栏操作**/
func Output(stub shim.ChaincodeStubInterface, paramList []module.OutputParam) peer.Response {
	wg.Add(len(paramList)) //添加队列数
	outputChan := make(chan ChanInfo, len(paramList))
	for i, v := range paramList {
		log.Logger.Info("Output --range:" + strconv.Itoa(i))
		go goOutput(stub, v, outputChan)
	}
	// 	获得chan 返回的值
	returnError := ReturnErrorInfo{}
	for j, _ := range paramList {
		log.Logger.Info("goOutput chan --- range :" + strconv.Itoa(j))
		tChan := <-outputChan //get Channel return value
		if tChan.Status == false {
			returnError.ErrorList = append(returnError.ErrorList, tChan)
		}
	}
	close(outputChan)
	wg.Wait()

	if len(returnError.ErrorList) >= 1 {
		returnError.Status = false
	} else {
		returnError.Status = true
	}
	jsonReturn, _ := json.Marshal(returnError)
	return shim.Success(jsonReturn)
}

/**检疫**/
func Exam(stub shim.ChaincodeStubInterface, paramList []module.ExamParam) peer.Response {
	wg.Add(len(paramList)) //添加队列数
	examChan := make(chan ChanInfo, len(paramList))
	for i, v := range paramList {
		log.Logger.Info("Exam --range:" + strconv.Itoa(i))
		go goExam(stub, v, examChan)
	}
	// 	获得chan 返回的值
	returnError := ReturnErrorInfo{}
	for j, _ := range paramList {
		log.Logger.Info("goExam chan --- range :" + strconv.Itoa(j))
		tChan := <-examChan //get Channel return value
		if tChan.Status == false {
			returnError.ErrorList = append(returnError.ErrorList, tChan)
		}
	}
	close(examChan)
	wg.Wait()
	if len(returnError.ErrorList) >= 1 {
		returnError.Status = false
	} else {
		returnError.Status = true
	}
	jsonReturn, _ := json.Marshal(returnError)
	return shim.Success(jsonReturn)
}

/**救治**/
func Save(stub shim.ChaincodeStubInterface, paramList []module.SaveParam) peer.Response {
	wg.Add(len(paramList)) //添加队列数
	saveChan := make(chan ChanInfo, len(paramList))
	for i, v := range paramList {
		log.Logger.Info("Save --range:" + strconv.Itoa(i))
		go goSave(stub, v, saveChan)
	}
	// 	获得chan 返回的值
	returnError := ReturnErrorInfo{}
	for j, _ := range paramList {
		log.Logger.Info("goSave chan --- range :" + strconv.Itoa(j))
		tChan := <-saveChan //get Channel return value
		if tChan.Status == false {
			returnError.ErrorList = append(returnError.ErrorList, tChan)
		}
	}
	close(saveChan)
	wg.Wait()
	if len(returnError.ErrorList) >= 1 {
		returnError.Status = false
	} else {
		returnError.Status = true
	}
	jsonReturn, _ := json.Marshal(returnError)
	return shim.Success(jsonReturn)
}

/**屠宰**/
func Butcher(stub shim.ChaincodeStubInterface, paramList []module.ButcherParam) peer.Response {
	wg.Add(len(paramList)) //添加队列数
	butcherChan := make(chan ChanInfo, len(paramList))
	for i, v := range paramList {
		log.Logger.Info("Butcher --range:" + strconv.Itoa(i))
		go goButcher(stub, v, butcherChan)
	}
	// 	获得chan 返回的值
	returnError := ReturnErrorInfo{}
	for j, _ := range paramList {
		log.Logger.Info("goButcher chan --- range :" + strconv.Itoa(j))
		tChan := <-butcherChan //get Channel return value
		if tChan.Status == false {
			returnError.ErrorList = append(returnError.ErrorList, tChan)
		}
	}
	close(butcherChan)
	wg.Wait()
	if len(returnError.ErrorList) >= 1 {
		returnError.Status = false
	} else {
		returnError.Status = true
	}
	jsonReturn, _ := json.Marshal(returnError)
	return shim.Success(jsonReturn)
}

/**灭尸**/
func Lost(stub shim.ChaincodeStubInterface, paramList []module.DestroyParam) peer.Response {
	wg.Add(len(paramList)) //添加队列数
	lostChan := make(chan ChanInfo, len(paramList))
	for i, v := range paramList {
		log.Logger.Info("Lost --range:" + strconv.Itoa(i))
		go goLost(stub, v, lostChan)
	}
	// 	获得chan 返回的值
	returnError := ReturnErrorInfo{}
	for j, _ := range paramList {
		log.Logger.Info("goLost chan --- range :" + strconv.Itoa(j))
		tChan := <-lostChan //get Channel return value
		if tChan.Status == false {
			returnError.ErrorList = append(returnError.ErrorList, tChan)
		}
	}
	close(lostChan)
	wg.Wait()
	if len(returnError.ErrorList) >= 1 {
		returnError.Status = false
	} else {
		returnError.Status = true
	}
	jsonReturn, _ := json.Marshal(returnError)
	return shim.Success(jsonReturn)
}

/**查询**/
func QueryByProduct(stub shim.ChaincodeStubInterface, param module.QueryParam) peer.Response {

	jsonBytes, err := stub.GetState(common.PRODUCT_INFO + common.ULINE + param.ProductId)
	if err != nil {
		log.Logger.Error("QueryByProduct err:" + err.Error())
		return shim.Error(err.Error())
	}
	return shim.Success(jsonBytes)
}

/**查询历史**/
func QueryHistoryByProduct(stub shim.ChaincodeStubInterface, param module.QueryParam) peer.Response {
	historys, err := stub.GetHistoryForKey(common.PRODUCT_INFO + common.ULINE + param.ProductId)
	if err != nil {
		log.Logger.Error("QueryHistoryByProduct err:" + err.Error())
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
	return shim.Success(resultsJSON)
}

/**查询批次**/
func QueryBatchByProduct(stub shim.ChaincodeStubInterface, param module.BatchParam) peer.Response {
	queryString := fmt.Sprintf("{\"selector\": {\"_id\": {\"$regex\": \"%s\"},\"batchNumber\":\"%s\"},\"limit\":\"%d\"}", common.PRODUCT_INFO, param.BatchNumber, 5000)
	queryResults, err := stub.GetQueryResult(queryString)
	if err != nil {
		return shim.Error(err.Error())
	}
	defer queryResults.Close()
	results := make([]module.Product, 0)
	for queryResults.HasNext() {
		result, err := queryResults.Next()
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
	return shim.Success(resultsJSON)
}

/**查询交易**/
func QueryByTX(stub shim.ChaincodeStubInterface, param module.QueryTxParam) peer.Response {
	queryString := fmt.Sprintf("{\"selector\": {\"_id\": {\"$regex\": \"%s\"},\"$or\":[{\"txId\":\"%s\"},{\"outputTxId\":\"%s\"},{\"lostTxId\":\"%s\"},{\"examTxId\":\"%s\"},{\"butcherTxId\":\"%s\"},{\"feedList.txId\":\"%s\"},{\"vaccineList.txId\":\"%s\"},{\"saveList.txId\":\"%s\"}]},\"limit\":\"%d\"}", common.PRODUCT_INFO, param.TxId, param.TxId, param.TxId, param.TxId, param.TxId, param.TxId, param.TxId, param.TxId, 5000)

	queryResults, err := stub.GetQueryResult(queryString)

	if err != nil {
		return shim.Error(err.Error())
	}
	defer queryResults.Close()
	results := make([]module.Product, 0)
	for queryResults.HasNext() {
		result, err := queryResults.Next()
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
	return shim.Success(resultsJSON)
}
