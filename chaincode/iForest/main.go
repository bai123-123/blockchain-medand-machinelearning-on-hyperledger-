package main

import (
	"fmt"
	"github.com/hyperledger/fabric/core/chaincode/shim"
	"github.com/hyperledger/fabric/protos/peer"
)

type IforestChaincode struct {

}

func (t *IforestChaincode) Init (stub shim.ChaincodeStubInterface) peer.Response {
	return shim.Success(nil)
}

func (t *IforestChaincode) Invoke(stub shim.ChaincodeStubInterface) peer.Response{
	fun,args := stub.GetFunctionAndParameters()

	if fun =="addDataAndTrain"{
		fmt.Println("train")
		return t.addData(stub,args)
	}else if fun =="addResult"{
		fmt.Println("predict......")
		return t.addResult(stub,args)
	}

	return shim.Error("wrong func name")
}

func main(){
	err:=shim.Start(new(IforestChaincode))
	if err!=nil{
		fmt.Printf("wrong when starting IrisesChaincode: %s",err)
	}
}