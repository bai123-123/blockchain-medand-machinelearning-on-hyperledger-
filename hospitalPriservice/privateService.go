package hospitalPriservice

import (
	"encoding/json"
	"fmt"
	"github.com/hyperledger/fabric-sdk-go/pkg/client/channel"
)

func (t *PrivateServiceSetup) SavePriEMR(emr EMR_pri) (string, error) {

	eventID := "eventAddEmr_pri"
	reg, notifier := regitserEvent(t.Client, t.ChaincodeID, eventID)
	defer t.Client.UnregisterChaincodeEvent(reg)


	b, err := json.Marshal(emr)
	if err != nil {
		return "", fmt.Errorf("something wrong about serialization")
	}

	req := channel.Request{ChaincodeID: t.ChaincodeID, Fcn: "addPriEMR", Args: [][]byte{b, []byte(eventID)}}
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


func (t *PrivateServiceSetup) FindEmrInfoByEmrNo(emrNo string) ([]byte, error){

	req := channel.Request{ChaincodeID: t.ChaincodeID, Fcn: "queryPriEmrByEmrNo", Args: [][]byte{[]byte(emrNo)}}
	respone, err := t.Client.Query(req)
	if err != nil {
		return []byte{0x00}, err
	}

	return respone.Payload, nil
}
