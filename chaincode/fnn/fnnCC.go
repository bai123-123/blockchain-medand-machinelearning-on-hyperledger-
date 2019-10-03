package main

import (
	"encoding/json"
	"fmt"
	"github.com/hyperledger/fabric/core/chaincode/shim"
	"github.com/hyperledger/fabric/protos/peer"
	"regexp"
	"strconv"
	"strings"
)


var retInput = make([][]float64, 0)
var retOutput = make([][]float64, 0)
type Info struct {
	Id string
	data []float64

}

func PutResult(stub shim.ChaincodeStubInterface,  info Info) ([]byte, bool) {
	b, err := json.Marshal(info)
	if err != nil {
		return nil, false
	}
	err = stub.PutState(info.Id,b)
	if err != nil {
		return nil, false
	}

	return b, true
}


func (t *FNNChaincode) addInputData(stub shim.ChaincodeStubInterface, args []string) peer.Response {
	if len(args) !=2{
		return shim.Error("wrong number of paramters given")
	}
	flysnowRegexp := regexp.MustCompile(`[0-9]*\.?[0-9]+`)
	s_content := string(args[0])
	lines := strings.Split(s_content, "\n")
	for _, line := range lines {
		line = strings.TrimRight(line, "\r\n")

		if len(line) == 0 {
			continue
		}
		tup := flysnowRegexp.FindAll([]byte(line),-1)
		X := make([]float64, 0)

		for _, x := range tup {

			f_x, _ := strconv.ParseFloat(string(x), 64)
			X = append(X, f_x)
		}
		retInput = append(retInput, X)

	}
	return shim.Success([]byte("LoadInputdata success"))
}

func (t *FNNChaincode) addOutputData(stub shim.ChaincodeStubInterface, args []string) peer.Response {
	if len(args) !=2{
		return shim.Error("wrong number of paramters given")
	}
	flysnowRegexp := regexp.MustCompile(`[0-9]*\.?[0-9]+`)
	s_content := string(args[0])
	lines := strings.Split(s_content, "\n")
	for _, line := range lines {
		line = strings.TrimRight(line, "\r\n")

		if len(line) == 0 {
			continue
		}
		tup := flysnowRegexp.FindAll([]byte(line),-1)
		X := make([]float64, 0)

		for _, x := range tup {

			f_x, _ := strconv.ParseFloat(string(x), 64)
			X = append(X, f_x)
		}
		retOutput = append(retOutput, X)

	}
	return shim.Success([]byte("LoadOutputdata success"))
}




func (t *FNNChaincode) classifyData(stub shim.ChaincodeStubInterface, args []string) peer.Response {
	inputTrain := retInput
	outputTrain := retOutput
	inputTrain = inputNormalization(inputTrain)    //Sample normalization
	outputTrain = outputNormalization(outputTrain) //Output expectation normalization
	trainNet(inputTrain, outputTrain)
	fmt.Println(yns)
	fmt.Println(counterResults)
	for i:=0;i< len(yns);i++{
		if yns[i]>2&&yns[i]<=3{
			fmt.Println("detect abnormal socre is ",yns[i],"save abnormal data to block")
			temp:=Info{
				Id:   string(i),
				data: inputTrain[i],
			}
			_, bl := PutResult(stub, temp) //create block
			if !bl {
				return shim.Error("wrong when saving abNormalInfo")
			}
		}
	}
	return shim.Success([]byte("Classifydata success"))
}