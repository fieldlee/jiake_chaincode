package control

import (
	"jiake_chaincode/log"

	"github.com/hyperledger/fabric/core/chaincode/shim"
	"github.com/hyperledger/fabric/protos/peer"
)

func (t *ProductTrace) Invoke(stub shim.ChaincodeStubInterface) peer.Response {
	log.Logger.Info("Invoke")

	funcation, args := stub.GetFunctionAndParameters()
	if funcation == "Register" { //入栏

	} else if funcation == "Feed" { //喂养

	} else if funcation == "" { //防疫

	} else if funcation == "" { //出栏

	} else if funcation == "" { //检疫

	} else if funcation == "" { //灭尸

	} else if funcation == "" { //分块

	} else if funcation == "" { //查询资产ID

	} else if funcation == "" { //查询批次ID

	} else if funcation == "" { //查询批次ID

	} else if funcation == "" { //查询批次ID

	}
	return shim.Error("Invalid invoke function name. " + funcation)
}
