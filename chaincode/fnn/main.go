package main

import (
	"fmt"
	"github.com/hyperledger/fabric/core/chaincode/shim"
	"github.com/hyperledger/fabric/protos/peer"
)

type FNNChaincode struct {

}

func (t *FNNChaincode) Init (stub shim.ChaincodeStubInterface) peer.Response {
	return shim.Success(nil)
}

func (t *FNNChaincode) Invoke(stub shim.ChaincodeStubInterface) peer.Response{
	fun,args := stub.GetFunctionAndParameters()
	if fun =="addInputData"{
		fmt.Println("loadInputData")
		return t.addInputData(stub,args)
	}else if fun =="addOutputData"{
		fmt.Println("loadOutputData")
		return t.addOutputData(stub,args)
	}else if fun =="classifyData" {
		fmt.Println("classify")
		return t.classifyData(stub, args)
	}
	return shim.Error("wrong func name")
}

func main(){
	err:=shim.Start(new(FNNChaincode))
	if err!=nil{
		fmt.Printf("wrong when starting IrisesChaincode: %s",err)
	}
}