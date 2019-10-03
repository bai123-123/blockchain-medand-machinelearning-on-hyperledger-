package hospitalPubservice

import (
	"encoding/json"
	"fmt"
	"github.com/hyperledger/fabric-sdk-go/pkg/client/channel"
)

func (t *PublicServiceSetup) SavePubEMR(emr EMR_common) (string, error) {

	eventID := "eventAddEmr_pub"
	reg, notifier := regitserEvent(t.Client, t.ChaincodeID, eventID)
	defer t.Client.UnregisterChaincodeEvent(reg)


	b, err := json.Marshal(emr)
	if err != nil {
		return "", fmt.Errorf("something wrong about serialization")
	}

	req := channel.Request{ChaincodeID: t.ChaincodeID, Fcn: "addCommonEMR", Args: [][]byte{b, []byte(eventID)}}
	respone, err := t.Client.Execute(req)
	if err != nil {
		return "", err
	}

	err = eventResult(notifier, eventID)
	if err != nil {
		return "", err
	}

	return string(respone.TransactionID), nil
}


func (t *PublicServiceSetup) FindEmrInfoByEmrNo(emrNo string) ([]byte, error){

	req := channel.Request{ChaincodeID: t.ChaincodeID, Fcn: "queryPubEmrByEmrNo", Args: [][]byte{[]byte(emrNo)}}
	respone, err := t.Client.Query(req)
	if err != nil {
		return []byte{0x00}, err
	}

	return respone.Payload, nil
}

func (t *PublicServiceSetup) UpdateMedicineQuality(emr EMR_common) (string, error){

	eventID := "eventUpdate_quality"
	reg, notifier := regitserEvent(t.Client, t.ChaincodeID, eventID)
	defer t.Client.UnregisterChaincodeEvent(reg)


	b, err := json.Marshal(emr)
	if err != nil {
		return "", fmt.Errorf("something wrong about serialization")
	}

	req := channel.Request{ChaincodeID: t.ChaincodeID, Fcn: "updateMedicineQuality", Args: [][]byte{b, []byte(eventID)}}
	respone, err := t.Client.Execute(req)
	if err != nil {
		return "", err
	}

	err = eventResult(notifier, eventID)
	if err != nil {
		return "", err
	}

	return string(respone.TransactionID), nil
}

func (t *PublicServiceSetup) CreatePushCode(emrNo string) ([]byte, error){

	req := channel.Request{ChaincodeID: t.ChaincodeID, Fcn: "createPushCode", Args: [][]byte{[]byte(emrNo)}}
	respone, err := t.Client.Query(req)
	if err != nil {
		return []byte{0x00}, err
	}

	return respone.Payload, nil
}