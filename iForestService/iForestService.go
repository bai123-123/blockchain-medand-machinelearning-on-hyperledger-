package iForestService

import (
	"encoding/json"
	"fmt"
	"github.com/hyperledger/fabric-sdk-go/pkg/client/channel"
)



func(t *IforestServiceSetup) LoadTrainData(flower string)([]byte,error){
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

func(t *IforestServiceSetup) AddResult(info CheckedInfo)([]byte, error){
	eventID :="eventPredict"
	reg,_:=regitserEvent(t.Client,t.ChaincodeID,eventID)
	defer t.Client.UnregisterChaincodeEvent(reg)


	b, err := json.Marshal(info)
	if err != nil {
		return []byte("none"), fmt.Errorf("something wrong about serialization")
	}
	req:=channel.Request{ChaincodeID:t.ChaincodeID, Fcn:"addResult",Args:[][]byte{b,[]byte(eventID)}}
	response,err :=t.Client.Query(req)

	if err != nil{
		fmt.Println(err)
		return []byte("none"), err
	}
	//err = eventResult(notifier, eventID)
	//if err != nil {
	//	return []byte("error"), err
	//}
	return response.Payload, nil

}

func (t *IforestServiceSetup) FindInfoByNo(flowerID string) ([]byte, error){

	req := channel.Request{ChaincodeID: t.ChaincodeID, Fcn: "query", Args: [][]byte{[]byte(flowerID)}}
	respone, err := t.Client.Query(req)
	if err != nil {
		return []byte{0x00}, err
	}

	return respone.Payload, nil
}
