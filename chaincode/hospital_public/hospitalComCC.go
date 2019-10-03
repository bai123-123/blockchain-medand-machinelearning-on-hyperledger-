package main

import (
	"crypto/md5"
	"encoding/json"
	"fmt"
	"github.com/hyperledger/fabric/core/chaincode/shim"
	"github.com/hyperledger/fabric/protos/peer"
	"io"
	"math/rand"
	"time"

)

func PutEMR(stub shim.ChaincodeStubInterface, emr EMR_common) ([]byte, bool) {


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


func GetEMRInfo(stub shim.ChaincodeStubInterface, EmrNO string) (EMR_common, bool)  {
	var emr EMR_common

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

func (t *HospitalComChaincode) addCommonEMR(stub shim.ChaincodeStubInterface, args []string) peer.Response {

	if len(args) != 2{
		return shim.Error("wrong number of amount about parameters")
	}

	var emr EMR_common
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

func (t *HospitalComChaincode) queryCommonEmrInfoByEmrID(stub shim.ChaincodeStubInterface, args []string) peer.Response {
	if len(args) != 1 {
		return shim.Error("wrong number of amount about parameters")
	}


	b, err := stub.GetState(args[0])
	if err != nil {
		return shim.Error("failed to query info")
	}

	if b == nil {
		return shim.Error("the queried info is null")
	}


	var emr EMR_common
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


func (t *HospitalComChaincode) updateInfo(stub shim.ChaincodeStubInterface, args []string) peer.Response {
	if len(args) != 2{
		return shim.Error("wrong number of amount about parameters")
	}

	var emr EMR_common
	err := json.Unmarshal([]byte(args[0]), &emr)
	if err != nil {
		return  shim.Error("wrong occour when deserialization")
	}


	result, bl := GetEMRInfo(stub, emr.EMRNo)
	if !bl{
		return shim.Error("failed to query info")
	}



	if(result.CommonInfo.AmountFixed>=emr.CommonInfo.AmountCurrent) {

		result.CommonInfo.AmountCurrent=emr.CommonInfo.AmountCurrent

		_, bl = PutEMR(stub, result)
		if !bl {
			return shim.Error("wrong when save info")
		}

		err = stub.SetEvent(args[1], []byte{})
		if err != nil {
			return shim.Error(err.Error())
		}

		return shim.Success([]byte("update current amount successfully"))

	} else {
		return shim.Error("the quality exceed the the quality given by hospital")
	}



}


func (t *HospitalComChaincode) createPushCodeByEMrID(stub shim.ChaincodeStubInterface, args []string) peer.Response {
	if len(args) != 1 {
		return shim.Error("worng parameter")
	}


	b, err := stub.GetState(args[0])
	if err != nil {
		return shim.Error("fail to qurey info by EHR No")
	}

	if b == nil {
		return shim.Error("the EHR does exist")
	}

	time:=time.Now()
	h:=md5.New()
	io.WriteString(h,string(b))
	io.WriteString(h,time.String())
	randomCode :=[]rune(fmt.Sprintf("%x",h.Sum(nil)))

	temp:=make([]rune,4)
	for i:=range temp{
		temp[i]=randomCode[rand.Intn((len(randomCode)))]
	}

	return shim.Success([]byte(string(temp)))
}
