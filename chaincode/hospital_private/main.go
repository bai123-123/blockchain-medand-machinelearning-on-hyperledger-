package main

import (
	"fmt"
	"github.com/hyperledger/fabric/core/chaincode/shim"
	"github.com/hyperledger/fabric/protos/peer"
)

type HospitalPriChaincode struct {

}

func (t *HospitalPriChaincode) Init (stub shim.ChaincodeStubInterface) peer.Response {
	return shim.Success(nil)
}

func (t *HospitalPriChaincode) Invoke(stub shim.ChaincodeStubInterface) peer.Response{
	fun,args := stub.GetFunctionAndParameters()

	if fun =="addPriEMR"{
		return t.addEMR(stub,args)
	}else if fun =="queryPriEmrByEmrNo"{
		return t.queryEmrInfoByEmrID(stub,args)
	}

	return shim.Error("wrong func name")
}

func main(){
	err:=shim.Start(new(HospitalPriChaincode))
	if err!=nil{
		fmt.Printf("wrong when starting IrisesChaincode: %s",err)
	}
}
