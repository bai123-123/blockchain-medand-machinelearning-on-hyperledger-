package irisesService

import (
	"encoding/json"
	"fmt"
	"github.com/hyperledger/fabric-sdk-go/pkg/client/channel"
)

func(t *IrisesServiceSetup) LoadTrainData(flower string)([]byte,error){
	eventID :="eventLoadData"
	reg,_ :=regitserEvent(t.Client,t.ChaincodeID,eventID)
	defer t.Client.UnregisterChaincodeEvent(reg)



	req:=channel.Request{ChaincodeID:t.ChaincodeID, Fcn:"addDataAndTrain",Args:[][]byte{[]byte(flower),[]byte(eventID)}}
	response,err :=t.Client.Query(req)

	if err != nil{
		fmt.Println(err)
		return []byte{0x00},err
	}

	return response.Payload,nil

}

func(t *IrisesServiceSetup) AddResult(flower Irises)([]byte, error){
	eventID :="eventPredict"
	reg,_ :=regitserEvent(t.Client,t.ChaincodeID,eventID)
	defer t.Client.UnregisterChaincodeEvent(reg)


	b, err := json.Marshal(flower)
	if err != nil {
		return []byte("none"), fmt.Errorf("something wrong about serialization")
	}
	req:=channel.Request{ChaincodeID:t.ChaincodeID, Fcn:"addResult",Args:[][]byte{b,[]byte(eventID)}}
	response,err :=t.Client.Query(req)

	if err != nil{
		fmt.Println(err)
		return []byte("none"), err
	}

	return response.Payload, nil

}

func (t *IrisesServiceSetup) FindInfoByNo(flowerID string) ([]byte, error){

	req := channel.Request{ChaincodeID: t.ChaincodeID, Fcn: "query", Args: [][]byte{[]byte(flowerID)}}
	respone, err := t.Client.Query(req)
	if err != nil {
		return []byte{0x00}, err
	}

	return respone.Payload, nil
}