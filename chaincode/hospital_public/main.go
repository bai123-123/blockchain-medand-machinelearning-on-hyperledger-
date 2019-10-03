package main

import (
	"fmt"
	"github.com/hyperledger/fabric/core/chaincode/shim"
	"github.com/hyperledger/fabric/protos/peer"
)

type HospitalComChaincode struct {

}

func (t *HospitalComChaincode) Init (stub shim.ChaincodeStubInterface) peer.Response {
	return shim.Success(nil)
}

func (t *HospitalComChaincode) Invoke(stub shim.ChaincodeStubInterface) peer.Response{
	fun,args := stub.GetFunctionAndParameters()

	if fun =="addCommonEMR"{
		return t.addCommonEMR(stub,args)
	}else if fun =="queryPubEmrByEmrNo"{
		return t.queryCommonEmrInfoByEmrID(stub,args)
	}else if fun=="updateMedicineQuality"{
		return t.updateInfo(stub,args)
	}else if fun=="createPushCode"{
		return t.createPushCodeByEMrID(stub,args)
	}

	return shim.Error("wrong func name")
}

func main(){
	err:=shim.Start(new(HospitalComChaincode))
	if err!=nil{
		fmt.Printf("wrong when starting IrisesChaincode: %s",err)
	}
}
