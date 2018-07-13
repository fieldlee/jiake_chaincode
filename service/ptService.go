package service

import (
	"encoding/json"
	"fmt"
	"jiakechaincode/common"
	"jiakechaincode/log"
	"jiakechaincode/module"
	"strconv"

	"github.com/hyperledger/fabric/core/chaincode/shim"
	"github.com/hyperledger/fabric/protos/peer"
)

// var wg sync.WaitGroup

/**产品注册信息**/
func Register(stub shim.ChaincodeStubInterface, paramList []module.RegitserParam) peer.Response {

	// 	获得chan 返回的值
	returnError := module.ReturnErrorInfo{}
	for i, v := range paramList {
		log.Logger.Info("Register --- send :" + strconv.Itoa(i))
		log.Logger.Info("Register --- send :" + v.ProductId)
		tChan := toRegister(stub, v)
		if tChan.Status == false {
			returnError.ErrorList = append(returnError.ErrorList, tChan)
		}
	}
	if len(returnError.ErrorList) >= 1 {
		returnError.Success = false
	} else {
		returnError.Success = true
		// 记录tx count number
		recordTxNumber(stub, len(paramList))
	}

	jsonReturn, _ := json.Marshal(returnError)
	log.Logger.Info(string(jsonReturn[:]))
	return shim.Success(jsonReturn)
}

// func Register(stub shim.ChaincodeStubInterface, paramList []module.RegitserParam) peer.Response {
// 	wg.Add(len(paramList)) //添加队列数
// 	registerChan := make(chan ChanInfo, len(paramList))
// 	for i, v := range paramList {
// 		log.Logger.Info("Register --- send :" + strconv.Itoa(i))
// 		log.Logger.Info("Register --- send :" + v.ProductId)
// 		go goRegister(stub, v, registerChan)
// 	}
// 	// 	获得chan 返回的值
// 	returnError := ReturnErrorInfo{}

// 	for j, _ := range paramList {
// 		log.Logger.Info("Register chan --- range :" + strconv.Itoa(j))
// 		tChan := <-registerChan //get Channel return value
// 		log.Logger.Info("Register --- RECIVED :" + tChan.ProductId)
// 		if tChan.Status == false {
// 			returnError.ErrorList = append(returnError.ErrorList, tChan)
// 		}
// 	}
// 	close(registerChan)
// 	wg.Wait()

// 	if len(returnError.ErrorList) >= 1 {
// 		returnError.Success = false
// 	} else {
// 		returnError.Success = true
// 	}

// 	jsonReturn, _ := json.Marshal(returnError)
// 	log.Logger.Info(string(jsonReturn[:]))
// 	return shim.Success(jsonReturn)
// }

/**喂养上链**/
func Feed(stub shim.ChaincodeStubInterface, paramList []module.FeedParam) peer.Response {

	// 	获得chan 返回的值
	returnError := module.ReturnErrorInfo{}
	for i, v := range paramList {
		log.Logger.Info("Feed --range:" + strconv.Itoa(i))
		tChan := toFeed(stub, v)
		if tChan.Status == false {
			returnError.ErrorList = append(returnError.ErrorList, tChan)
		}
	}

	if len(returnError.ErrorList) >= 1 {
		returnError.Success = false
	} else {
		returnError.Success = true
		// 记录tx count number
		recordTxNumber(stub, len(paramList))
	}

	jsonReturn, _ := json.Marshal(returnError)
	return shim.Success(jsonReturn)
}

/**防疫信息上链**/
func Vaccine(stub shim.ChaincodeStubInterface, paramList []module.VaccineParam) peer.Response {

	// 	获得chan 返回的值
	returnError := module.ReturnErrorInfo{}
	for i, v := range paramList {
		log.Logger.Info("Vaccine  ---range:" + strconv.Itoa(i))
		tChan := toVaccine(stub, v)
		if tChan.Status == false {
			returnError.ErrorList = append(returnError.ErrorList, tChan)
		}
	}
	if len(returnError.ErrorList) >= 1 {
		returnError.Success = false
	} else {
		returnError.Success = true
		// 记录tx count number
		recordTxNumber(stub, len(paramList))
	}
	jsonReturn, _ := json.Marshal(returnError)
	return shim.Success(jsonReturn)
}

/**出栏操作**/
func Output(stub shim.ChaincodeStubInterface, paramList []module.OutputParam) peer.Response {

	// 	获得chan 返回的值
	returnError := module.ReturnErrorInfo{}
	for i, v := range paramList {
		log.Logger.Info("Output --range:" + strconv.Itoa(i))
		tChan := toOutput(stub, v)
		if tChan.Status == false {
			returnError.ErrorList = append(returnError.ErrorList, tChan)
		}
	}
	if len(returnError.ErrorList) >= 1 {
		returnError.Success = false
	} else {
		returnError.Success = true
		// 记录tx count number
		recordTxNumber(stub, len(paramList))
	}
	jsonReturn, _ := json.Marshal(returnError)
	return shim.Success(jsonReturn)
}

/**检疫**/
func Exam(stub shim.ChaincodeStubInterface, paramList []module.ExamParam) peer.Response {

	// 	获得chan 返回的值
	returnError := module.ReturnErrorInfo{}
	for i, v := range paramList {
		log.Logger.Info("Exam --range:" + strconv.Itoa(i))
		tChan := toExam(stub, v)
		if tChan.Status == false {
			returnError.ErrorList = append(returnError.ErrorList, tChan)
		}
	}

	if len(returnError.ErrorList) >= 1 {
		returnError.Success = false
	} else {
		returnError.Success = true
		// 记录tx count number
		recordTxNumber(stub, len(paramList))
	}
	jsonReturn, _ := json.Marshal(returnError)
	return shim.Success(jsonReturn)
}

/**救治**/
func Save(stub shim.ChaincodeStubInterface, paramList []module.SaveParam) peer.Response {

	// 	获得chan 返回的值
	returnError := module.ReturnErrorInfo{}
	for i, v := range paramList {
		log.Logger.Info("Save --range:" + strconv.Itoa(i))
		tChan := toSave(stub, v)
		if tChan.Status == false {
			returnError.ErrorList = append(returnError.ErrorList, tChan)
		}
	}
	if len(returnError.ErrorList) >= 1 {
		returnError.Success = false
	} else {
		returnError.Success = true
		// 记录tx count number
		recordTxNumber(stub, len(paramList))
	}
	jsonReturn, _ := json.Marshal(returnError)
	return shim.Success(jsonReturn)
}

/**待宰**/
func WaitButcher(stub shim.ChaincodeStubInterface, paramList []module.WaitButcherParam) peer.Response {

	// 	获得chan 返回的值
	returnError := module.ReturnErrorInfo{}
	for i, v := range paramList {
		log.Logger.Info("WaitButcher --range:" + strconv.Itoa(i))
		tChan := toWaitButcher(stub, v)
		if tChan.Status == false {
			returnError.ErrorList = append(returnError.ErrorList, tChan)
		}
	}

	if len(returnError.ErrorList) >= 1 {
		returnError.Success = false
	} else {
		returnError.Success = true
		// 记录tx count number
		recordTxNumber(stub, len(paramList))
	}
	jsonReturn, _ := json.Marshal(returnError)
	return shim.Success(jsonReturn)
}

/**屠宰**/
func Butcher(stub shim.ChaincodeStubInterface, paramList []module.ButcherParam) peer.Response {

	// 	获得chan 返回的值
	returnError := module.ReturnErrorInfo{}
	for i, v := range paramList {
		log.Logger.Info("Butcher --range:" + strconv.Itoa(i))
		tChan := toButcher(stub, v)
		if tChan.Status == false {
			returnError.ErrorList = append(returnError.ErrorList, tChan)
		}
	}

	if len(returnError.ErrorList) >= 1 {
		returnError.Success = false
	} else {
		returnError.Success = true
		// 记录tx count number
		recordTxNumber(stub, len(paramList))
	}
	jsonReturn, _ := json.Marshal(returnError)
	return shim.Success(jsonReturn)
}

/**灭尸**/
func Lost(stub shim.ChaincodeStubInterface, paramList []module.DestroyParam) peer.Response {

	// 	获得chan 返回的值
	returnError := module.ReturnErrorInfo{}
	for i, v := range paramList {
		log.Logger.Info("Lost --range:" + strconv.Itoa(i))
		tChan := toLost(stub, v)
		if tChan.Status == false {
			returnError.ErrorList = append(returnError.ErrorList, tChan)
		}
	}

	if len(returnError.ErrorList) >= 1 {
		returnError.Success = false
	} else {
		returnError.Success = true
		// 记录tx count number
		recordTxNumber(stub, len(paramList))
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

/**查询交易历史**/
func QueryTransferHistoryByProduct(stub shim.ChaincodeStubInterface, param module.QueryParam) peer.Response {
	historys, err := stub.GetHistoryForKey(common.PRODUCT_TRANSFER + common.ULINE + param.ProductId)
	if err != nil {
		log.Logger.Error("QueryHistoryByProduct err:" + err.Error())
		return shim.Error(err.Error())
	}
	defer historys.Close()
	results := make([]module.ChangeAssetOwner, 0)
	for historys.HasNext() {
		result, err := historys.Next()
		if err != nil {
			return shim.Error(err.Error())
		}
		changeAsset := module.ChangeAssetOwner{}

		err = json.Unmarshal(result.Value, &changeAsset)
		if err != nil {
			return shim.Error(err.Error())
		} else {
			results = append(results, changeAsset)
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

func recordTxNumber(stub shim.ChaincodeStubInterface, num int) {
	fmt.Println("recordTxNumber:")
	// 记录tx count number
	countByts, err := stub.GetState(common.TX_COUNT + common.ULINE + common.TX_NUMBER)
	if err != nil {
		var txCount module.TxCount
		txCount.Count = uint64(num)
		txBytes, _ := json.Marshal(txCount)
		fmt.Println(txCount)
		_ = stub.PutState(common.TX_COUNT+common.ULINE+common.TX_NUMBER, txBytes)
	} else {
		txCount := module.TxCount{Count: 0}
		err = json.Unmarshal(countByts, &txCount)
		if err != nil {
			var txInfo module.TxCount
			txInfo.Count = uint64(num)
			txInfoBytes, _ := json.Marshal(txInfo)
			fmt.Println(txInfo)
			_ = stub.PutState(common.TX_COUNT+common.ULINE+common.TX_NUMBER, txInfoBytes)
		} else {
			txCount.Count = txCount.Count + uint64(num)
			txBytes, _ := json.Marshal(txCount)
			fmt.Println(txCount)
			_ = stub.PutState(common.TX_COUNT+common.ULINE+common.TX_NUMBER, txBytes)
		}
	}
}
