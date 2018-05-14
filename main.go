package main

import (
	"fmt"

	"github.com/hyperledger/fabric/core/chaincode/shim"
)

type ProductTrace struct {
}

func main() {
	err := shim.Start(new(ProductTrace))
	if err != nil {
		fmt.Printf("Error starting ProductTrace: %s", err)
	}
}
