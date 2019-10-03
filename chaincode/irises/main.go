package main

import (
	"fmt"
	"github.com/hyperledger/fabric/core/chaincode/shim"
	"github.com/hyperledger/fabric/protos/peer"
)

type IrisesChaincode struct {

}

func (t *IrisesChaincode) Init (stub shim.ChaincodeStubInterface) peer.Response {
	return shim.Success(nil)
}

func (t *IrisesChaincode) Invoke(stub shim.ChaincodeStubInterface) peer.Response{
	fun,args := stub.GetFunctionAndParameters()

	if fun =="addDataAndTrain"{
		fmt.Println("fuckfuckfuck")
		return t.addTraindataAndTrain(stub,args)
	}else if fun =="addResult"{
		fmt.Println("predict......")
		return t.addResult(stub,args)
	}else if fun =="query"{
		return t.queryFlowerInfoByID(stub,args)
	}

	return shim.Error("wrong func name")
}

func main(){
	err:=shim.Start(new(IrisesChaincode))
	if err!=nil{
		fmt.Printf("wrong when starting IrisesChaincode: %s",err)
	}
}