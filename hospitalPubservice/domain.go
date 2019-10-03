package hospitalPubservice

import (
	"fmt"
	"github.com/hyperledger/fabric-sdk-go/pkg/client/channel"
	"github.com/hyperledger/fabric-sdk-go/pkg/common/providers/fab"
	"time"
)

type EMR_common struct {
	EMRNo string
	MedNo string
	DoctorNo string
	POM string
	Date string
	CommonInfo Common
	VerifyDate string
}

type Common struct {
	Name string
	Gender string
	Age string
	BirthDate string
	Contact string
	Medicine string
	QuantityFixed string
	AmountFixed string
	AmountCurrent string
}



type PublicServiceSetup struct {
	ChaincodeID string
	Client *channel.Client
}

func regitserEvent(client *channel.Client,chanincodeID,eventID string)(fab.Registration,<-chan *fab.CCEvent){
	reg, notifier,err := client.RegisterChaincodeEvent(chanincodeID,eventID)
	if err!=nil{
		fmt.Println("fail to registe chaincode event: %s",err)
	}

	return reg,notifier
}

func eventResult(notifier <-chan *fab.CCEvent, eventID string) error {
	select {
	case ccEvent:=<-notifier:
		fmt.Printf("receive chaincode event %v\n",ccEvent)
	case <-time.After(time.Second*20):
		return fmt.Errorf("can not receive event according to event ID (%s)",eventID)

	}
	return nil
}

