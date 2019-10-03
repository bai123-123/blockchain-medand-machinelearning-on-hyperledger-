package main

import (
	"fmt"
	"github.com/hyperledger/fabric/core/chaincode/shim"
	"github.com/hyperledger/fabric/protos/peer"
)

type TestChaincode struct {

}

func (t *TestChaincode) Init (stub shim.ChaincodeStubInterface) peer.Response {
	return shim.Success(nil)
}

func (t *TestChaincode) Invoke(stub shim.ChaincodeStubInterface) peer.Response{
	fun,args := stub.GetFunctionAndParameters()

	if fun =="add"{
		return t.addFlower(stub,args)
	}else if fun =="query"{
		return t.queryFlowerInfoByID(stub,args)
	}

	return shim.Error("wrong func name")
}

func main(){
	err:=shim.Start(new(TestChaincode))
	if err!=nil{
		fmt.Printf("wrong when starting IrisesChaincode: %s",err)
	}
}

