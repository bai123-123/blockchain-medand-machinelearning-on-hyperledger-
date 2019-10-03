package main

import (
	"encoding/json"
	"github.com/hyperledger/fabric/core/chaincode/shim"
	"github.com/hyperledger/fabric/protos/peer"
)

type Irises struct {
	FlowerID string
	Calyx_length string
	Calyx_width string
	Petal_length string
	Petal_width string

	//**!!

}

func PutFlower(stub shim.ChaincodeStubInterface, flower Irises) ([]byte, bool) {


	b, err := json.Marshal(flower)
	if err != nil {
		return nil, false
	}


	err = stub.PutState(flower.FlowerID, b)
	if err != nil {
		return nil, false
	}

	return b, true
}

func (t *TestChaincode) addFlower(stub shim.ChaincodeStubInterface, args []string) peer.Response {

	if len(args) != 2{
		return shim.Error("wrong number of amount about parameters")
	}
	var flower Irises
	err := json.Unmarshal([]byte(args[0]), &flower)
	if err != nil {
		return shim.Error("wrong occour when deserialization")
	}
	_, bl := PutFlower(stub, flower)
	if !bl {
		return shim.Error("wrong when saving EHR")
	}

	err = stub.SetEvent(args[1], []byte{})
	if err != nil {
		return shim.Error(err.Error())
	}

	return shim.Success([]byte("add information successfully"))
}

func (t *TestChaincode) queryFlowerInfoByID(stub shim.ChaincodeStubInterface, args []string) peer.Response {
	if len(args) != 1 {
		return shim.Error("wrong number of amount about parameters")
	}


	b, err := stub.GetState(args[0])
	if err != nil {
		return shim.Error("failed to record info")
	}

	if b == nil {
		return shim.Error("the recorded info is null")
	}


	var flower Irises
	err = json.Unmarshal(b, &flower)
	if err != nil {
		return  shim.Error("wrong occour when deserialization emr info")
	}


	result, err := json.Marshal(flower)
	if err != nil {
		return shim.Error("wrong occour when serialization emr info")
	}
	return shim.Success(result)
}



