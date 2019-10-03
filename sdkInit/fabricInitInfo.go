package sdkInit

import "github.com/hyperledger/fabric-sdk-go/pkg/client/resmgmt"

type InitInfo struct {
	Channel1ID     string
	Channel2ID     string
	Channel3ID     string

	Channel1Config string
	Channel2Config string
	Channel3Config string

	Org1Admin      string
	Org1Name       string

	Org2Admin      string
	Org2Name       string

	Org3Admin      string
	Org3Name       string


	OrdererOrgName	string
	Org1ResMgmt *resmgmt.Client
	Org2ResMgmt *resmgmt.Client
	Org3ResMgmt *resmgmt.Client

	ChaincodeOneID	string
	ChaincodeOneGoPath	string
	ChaincodeOnePath	string
	Org1UserName	string


	ChaincodeTwoID	string
	ChaincodeTwoGoPath	string
	ChaincodeTwoPath	string
	Org2UserName	string

	ChaincodeThreeID	string
	ChaincodeThreeGoPath	string
	ChaincodeThreePath	string
	Org3UserName	string


}