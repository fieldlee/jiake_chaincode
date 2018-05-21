package main

import (
	"fmt"
	"jiakechaincode/control"

	"github.com/hyperledger/fabric/core/chaincode/shim"
)

func main() {
	err := shim.Start(new(control.ProductTrace))
	if err != nil {
		fmt.Printf("Error starting ProductTrace: %s", err)
	}
}
