package main

import (
	"demo_test_med/chaincode/irises/fxsjy/RF.go/RF"
	"encoding/json"
	"github.com/hyperledger/fabric/core/chaincode/shim"
	"github.com/hyperledger/fabric/protos/peer"
	"strconv"
	"strings"

)
var number_training_data int
var number_test_data int
var train_inputs [][]interface{}
var train_targets []string
var test_inputs [][]interface{}
var test_targets []string
var Forest *RF.Forest
var CurrentType int
var time int

type Irises struct {
	FlowerID string
	Calyx_length string
	Calyx_width string
	Petal_length string
	Petal_width string

	//attribute to be added
	//
	IrisType string
   	//
	//attribute to be added
}

type currentStatus struct {
	flowerType string
	ifChangeable bool
}

func PutPredictedResult(stub shim.ChaincodeStubInterface, flower Irises) ([]byte, bool) {
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

func (t *IrisesChaincode) queryFlowerInfoByID(stub shim.ChaincodeStubInterface, args []string) peer.Response {
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

func (t *IrisesChaincode) addTraindataAndTrain(stub shim.ChaincodeStubInterface, args []string) peer.Response {
	if len(args) !=2{
		return shim.Error("wrong number of paramters given")
	}
	s_content := string(args[0])
	lines := strings.Split(s_content, "\n")

	// set up variables for random forest
	inputs := make([][]interface{}, 0)
	targets := make([]string, 0)
	for _, line := range lines {

		line = strings.TrimRight(line, "\r\n")

		if len(line) == 0 {
			continue
		}
		tup := strings.Split(line, ",")
		pattern := tup[:len(tup)-1]
		target := tup[len(tup)-1]
		X := make([]interface{}, 0)
		for _, x := range pattern {
			f_x, _ := strconv.ParseFloat(x, 64)
			X = append(X, f_x)
		}
		inputs = append(inputs, X)

		targets = append(targets, target)
	}

	Forest = RF.DefaultForest(inputs, targets, 100) //100 trees

	return shim.Success([]byte("addTraindata success"))
}

func (t *IrisesChaincode) addResult(stub shim.ChaincodeStubInterface, args []string) peer.Response {

	if len(args) != 2{
		return shim.Error("wrong number of amount about parameters")
	}
	var flower Irises
	err := json.Unmarshal([]byte(args[0]), &flower)
	if err != nil {
		return shim.Error("wrong occour when deserialization")
	}
	var data [4]string
	data[0] = flower.Calyx_length
	data[1] = flower.Calyx_width
	data[2]=flower.Petal_length
	data[3]=flower.Petal_width
	input := make([]interface{}, 0)
	for _, x := range data {
		f_x, _ := strconv.ParseFloat(x, 64)
		input = append(input, f_x)
	}
	output := Forest.Predicate(input)
	time++
	judge,err := strconv.Atoi(output)
	responseFlower :=currentStatus{
		flowerType: output,
		ifChangeable: false,
	}
	if time>1&&CurrentType!=judge{
		responseFlower.ifChangeable = true
	}
	CurrentType =judge
	//logic to be added
	//
	flower.IrisType = output
	//
	//logic to be added
	PutPredictedResult(stub, flower)
	err = stub.SetEvent(args[1], []byte{})
	if err != nil {
		return shim.Error(err.Error())
	}
	return shim.Success([]byte(output))
}





