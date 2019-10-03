package fnnService

import (
	"fmt"
	"github.com/hyperledger/fabric-sdk-go/pkg/client/channel"
)

func(t *FnnServiceSetup) LoadInputData(data string)([]byte,error){
	eventID :="eventLoadInputData"
	reg,_ :=regitserEvent(t.Client,t.ChaincodeID,eventID)
	defer t.Client.UnregisterChaincodeEvent(reg)
	req:=channel.Request{ChaincodeID:t.ChaincodeID, Fcn:"addInputData",Args:[][]byte{[]byte(data),[]byte(eventID)}}
	response,err :=t.Client.Query(req)
	if err != nil{
		fmt.Println(err)
		return []byte{0x00},err
	}
	return response.Payload,nil
}

func(t *FnnServiceSetup) LoadOutputData(data string)([]byte,error){
	eventID :="eventLoadInputData"
	reg,_ :=regitserEvent(t.Client,t.ChaincodeID,eventID)
	defer t.Client.UnregisterChaincodeEvent(reg)
	req:=channel.Request{ChaincodeID:t.ChaincodeID, Fcn:"addOutputData",Args:[][]byte{[]byte(data),[]byte(eventID)}}
	response,err :=t.Client.Query(req)
	if err != nil{
		fmt.Println(err)
		return []byte{0x00},err
	}
	return response.Payload,nil
}

func(t *FnnServiceSetup) ClassifyData()([]byte,error){
	eventID :="eventClassifyData"
	reg,_ :=regitserEvent(t.Client,t.ChaincodeID,eventID)
	defer t.Client.UnregisterChaincodeEvent(reg)
	req:=channel.Request{ChaincodeID:t.ChaincodeID, Fcn:"classifyData",Args:[][]byte{[]byte(""),[]byte(eventID)}}
	response,err :=t.Client.Query(req)
	if err != nil{
		fmt.Println(err)
		return []byte{0x00},err
	}
	return response.Payload,nil
}