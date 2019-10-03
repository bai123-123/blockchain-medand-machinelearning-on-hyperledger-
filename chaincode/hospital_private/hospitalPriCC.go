package main

import (
	"encoding/json"
	"github.com/hyperledger/fabric/core/chaincode/shim"
	"github.com/hyperledger/fabric/protos/peer"
)

func PutEMR(stub shim.ChaincodeStubInterface, emr EMR_pri) ([]byte, bool) {


	b, err := json.Marshal(emr)
	if err != nil {
		return nil, false
	}


	err = stub.PutState(emr.EMRNo, b)
	if err != nil {
		return nil, false
	}

	return b, true
}


func GetEMRInfo(stub shim.ChaincodeStubInterface, EmrNO string) (EMR_pri, bool)  {
	var emr EMR_pri

	b, err := stub.GetState(EmrNO)
	if err != nil {
		return emr, false
	}

	if b == nil {
		return emr, false
	}


	err = json.Unmarshal(b, &emr)
	if err != nil {
		return emr, false
	}


	return emr, true
}

func (t *HospitalPriChaincode) addEMR(stub shim.ChaincodeStubInterface, args []string) peer.Response {

	if len(args) != 2{
		return shim.Error("wrong number of amount about parameters")
	}

	var emr EMR_pri
	err := json.Unmarshal([]byte(args[0]), &emr)
	if err != nil {
		return shim.Error("wrong occour when deserialization")
	}


	_, exist := GetEMRInfo(stub, emr.EMRNo)
	if exist {
		return shim.Error("the EHR number already exist")
	}

	_, bl := PutEMR(stub, emr)
	if !bl {
		return shim.Error("wrong when saving EHR")
	}

	err = stub.SetEvent(args[1], []byte{})
	if err != nil {
		return shim.Error(err.Error())
	}

	return shim.Success([]byte("add information successfully"))
}

func (t *HospitalPriChaincode) queryEmrInfoByEmrID(stub shim.ChaincodeStubInterface, args []string) peer.Response {
	if len(args) != 1 {
		return shim.Error("wrong number of amount about parameters")
	}

	// 根据身份证号码查询edu状态
	b, err := stub.GetState(args[0])
	if err != nil {
		return shim.Error("failed to record info")
	}

	if b == nil {
		return shim.Error("the recorded info is null")
	}

	// 对查询到的状态进行反序列化
	var emr EMR_pri
	err = json.Unmarshal(b, &emr)
	if err != nil {
		return  shim.Error("wrong occour when deserialization emr info")
	}


	result, err := json.Marshal(emr)
	if err != nil {
		return shim.Error("wrong occour when serialization emr info")
	}
	return shim.Success(result)
}


