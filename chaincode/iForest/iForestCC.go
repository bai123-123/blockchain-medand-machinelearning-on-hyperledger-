package main

import (

	"demo_test_med/chaincode/iForest/go-iforest/iforest"
	"encoding/json"
	"fmt"
	"github.com/hyperledger/fabric/core/chaincode/shim"
	"github.com/hyperledger/fabric/protos/peer"
	"strconv"
	"strings"

)

var inputData [][]float64
var forest  *iforest.Forest

type User struct {
	UserID string
	isAbnormal bool
	score float64
}

func PutResult(stub shim.ChaincodeStubInterface,  user User) ([]byte, bool) {
	b, err := json.Marshal(user)
	if err != nil {
		return nil, false
	}
	err = stub.PutState(user.UserID, b)
	if err != nil {
		return nil, false
	}

	return b, true
}

func (t *IforestChaincode) addData(stub shim.ChaincodeStubInterface, args []string) peer.Response {
	if len(args) !=2{
		return shim.Error("wrong number of paramters given")
	}
	s_content := string(args[0])
	lines := strings.Split(s_content, "\n")

	return shim.Success([]byte("addTraindata success"))
}

func (t *IforestChaincode) addResult(stub shim.ChaincodeStubInterface, args []string) peer.Response {

	//
	//userIDs := make([]string, 0)
	//
	//// set up variables for random forest and ParaseData
	//for _, line := range lines {
	//	line = strings.TrimRight(line, "\r\n")
	//	if len(line) == 0 {
	//		continue
	//	}
	//	tup := strings.Split(line, ",")
	//	userIDs=append(userIDs, tup[1])
	//	pattern := tup[2:]
	//	X := make([]float64, 0)
	//	for _, x := range pattern {
	//		f_x, _ := strconv.ParseFloat(x, 64)
	//		X = append(X, f_x)
	//	}
	//	inputData =append(inputData,X)
	//}
	//treesNumber := 100
	//subsampleSize := 130
	//outliersRatio := 0.09
	//routinesNumber := 10
	//forest = iforest.NewForest(treesNumber, subsampleSize, outliersRatio)
	//forest.Train(inputData)
	//forest.Test(inputData)
	//forest.TestParallel(inputData, routinesNumber)
	//anomalyScores := forest.AnomalyScores
	//labelsTest := forest.Labels
	//fmt.Println(anomalyScores)
	//fmt.Println(labelsTest)
	//
	//for i:=0;i< len(labelsTest);i++{
	//	if labelsTest[i]==1{
	//		fmt.Println("detect abnormal abnormalUserId is :",userIDs[i]," create block with abnormal user")
	//		abnormalUser:=User{
	//			UserID:  userIDs[i],
	//			score:   anomalyScores[i],
	//		}
	//		_, bl := PutResult(stub, abnormalUser) //create block
	//		if !bl {
	//			return shim.Error("wrong when saving abNormalInfo")
	//		}
	//
	//
	//	}
	//
	//}

	//if len(args) != 2{
	//	return shim.Error("wrong number of amount about parameters")
	//}
	//var checkedInfo struct{
	//	data [][]float64
	//}
	//err := json.Unmarshal([]byte(args[0]), &checkedInfo)
	//if err != nil {
	//	return shim.Error("wrong occour when deserialization")
	//}
	//fmt.Println("ghghhghghhgghgghghghhg")
	//fmt.Println(args[0])
	//labels, _,err:= forest.Predict(checkedInfo.data)
	//if labels[0] ==1{
	//	var temp  User
	//	temp.score=strconv.FormatFloat(scores[0],'E',-1,32)
	//	temp.UserID =strconv.FormatFloat(checkedInfo.data[0][1],'E',-1,32)
	//	PutPredictedResult(stub, temp)
	//	err = stub.SetEvent(args[1], []byte{})
	//	if err != nil {
	//		return shim.Error(err.Error())
	//	}
	//	return shim.Success([]byte(strconv.Itoa(labels[0])))
	//}else {
	//	return shim.Success([]byte(strconv.Itoa(labels[0])))
	//}

	//temp:=strconv.Itoa(labels[0])

	return shim.Success([]byte(""))
}






