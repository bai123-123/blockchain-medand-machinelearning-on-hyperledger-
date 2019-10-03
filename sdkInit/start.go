package sdkInit

import (
	"github.com/hyperledger/fabric-sdk-go/pkg/client/channel"
	"github.com/hyperledger/fabric-sdk-go/pkg/common/errors/retry"
	"fmt"
	"github.com/hyperledger/fabric-sdk-go/pkg/client/resmgmt"

	"github.com/hyperledger/fabric-sdk-go/pkg/common/providers/msp"

	_"github.com/hyperledger/fabric-sdk-go/pkg/core/config"
	mspclient "github.com/hyperledger/fabric-sdk-go/pkg/client/msp"
	"github.com/hyperledger/fabric-sdk-go/pkg/fabsdk"
	"github.com/hyperledger/fabric-sdk-go/pkg/fab/ccpackager/gopackager"
	"github.com/hyperledger/fabric-sdk-go/third_party/github.com/hyperledger/fabric/common/cauthdsl"

)







const ChaincodeVersion  = "1.0"

func SetupSDK(ConfigFile string, initialized bool) (*fabsdk.FabricSDK, error) {

	if initialized {
		return nil, fmt.Errorf("Fabric SDK had been instantiation")
	}

	sdk, err := fabsdk.New(ConfigBackend)
	if err != nil {
		return nil, fmt.Errorf("instantiation Fabric SDK failed: %v", err)
	}

	fmt.Println("Instantiation about Fabric SDK success")
	return sdk, nil
}


func CreateChanneltwo(sdk *fabsdk.FabricSDK, info *InitInfo) error   {

	org2ClientContext := sdk.Context(fabsdk.WithUser(info.Org2Admin), fabsdk.WithOrg(info.Org2Name))
	if org2ClientContext == nil{
		return fmt.Errorf("create context failed")
	}

	resOrg2MgmtClient , err := resmgmt.New(org2ClientContext)
	if err !=nil{
		return fmt.Errorf("cteate org1 client by context failed: %v",err)
	}

	org2MspClient , err := mspclient.New(sdk.Context(), mspclient.WithOrg(info.Org2Name))

	if err !=nil{
		return fmt.Errorf("create Org MSP client faile according to Orgname : %v",err)
	}



	org2AdminIdentity, err := org2MspClient.GetSigningIdentity(info.Org2Admin)
	if err != nil{
		return fmt.Errorf("failed to get sign according to id : %v",err)
	}

	channel1Req := resmgmt.SaveChannelRequest{ChannelID:info.Channel2ID, ChannelConfigPath:info.Channel2Config, SigningIdentities:[]msp.SigningIdentity{org2AdminIdentity}}

	_,err = resOrg2MgmtClient.SaveChannel(channel1Req,resmgmt.WithRetry(retry.DefaultResMgmtOpts), resmgmt.WithOrdererEndpoint(info.OrdererOrgName))
	if err!=nil{
		return fmt.Errorf("fail to create channel %v",err)
	}

	info.Org2ResMgmt = resOrg2MgmtClient
	err = info.Org2ResMgmt.JoinChannel(info.Channel2ID, resmgmt.WithRetry(retry.DefaultResMgmtOpts), resmgmt.WithOrdererEndpoint(info.OrdererOrgName))
	if err != nil {
		return fmt.Errorf("org2 Peers failed to join channel: %v", err)
	}

	fmt.Println("org2 peers join channel success.")
	return nil

}

func CreateChannelthree(sdk *fabsdk.FabricSDK, info *InitInfo) error{

	org3ClientContext := sdk.Context(fabsdk.WithUser(info.Org3Admin), fabsdk.WithOrg(info.Org3Name))
	if org3ClientContext == nil{
		return fmt.Errorf("create context failed")
	}

	resOrg3MgmtClient , err := resmgmt.New(org3ClientContext)
	if err !=nil{
		return fmt.Errorf("cteate org1 client by context failed: %v",err)
	}

	org3MspClient , err := mspclient.New(sdk.Context(), mspclient.WithOrg(info.Org3Name))

	if err !=nil{
		return fmt.Errorf("create Org MSP client faile according to Orgname : %v",err)
	}



	org3AdminIdentity, err := org3MspClient.GetSigningIdentity(info.Org3Admin)
	if err != nil{
		return fmt.Errorf("failed to get sign according to id : %v",err)
	}

	channel1Req := resmgmt.SaveChannelRequest{ChannelID:info.Channel3ID, ChannelConfigPath:info.Channel3Config, SigningIdentities:[]msp.SigningIdentity{org3AdminIdentity}}

	_,err = resOrg3MgmtClient.SaveChannel(channel1Req,resmgmt.WithRetry(retry.DefaultResMgmtOpts), resmgmt.WithOrdererEndpoint(info.OrdererOrgName))
	if err!=nil{
		return fmt.Errorf("fail to create channel %v",err)
	}

	info.Org3ResMgmt = resOrg3MgmtClient
	err = info.Org3ResMgmt.JoinChannel(info.Channel3ID,resmgmt.WithRetry(retry.DefaultResMgmtOpts), resmgmt.WithOrdererEndpoint(info.OrdererOrgName))
	if err != nil {
		return fmt.Errorf("org3 Peers failed to join channel: %v", err)
	}

	fmt.Println("org3 peers join channel success.")
	return nil

}

func CreateChannelone(sdk *fabsdk.FabricSDK, info *InitInfo) error{

	org1ClientContext := sdk.Context(fabsdk.WithUser(info.Org1Admin), fabsdk.WithOrg(info.Org1Name))
	if org1ClientContext == nil{
		return fmt.Errorf("create context failed")
	}

	resOrg1MgmtClient , err := resmgmt.New(org1ClientContext)
	if err !=nil{
		return fmt.Errorf("cteate org1 client by context failed: %v",err)
	}

	org1MspClient , err := mspclient.New(sdk.Context(), mspclient.WithOrg(info.Org1Name))

	if err !=nil{
		return fmt.Errorf("create Org MSP client faile according to Orgname : %v",err)
	}



	org1AdminIdentity, err := org1MspClient.GetSigningIdentity(info.Org1Admin)
	if err != nil{
		return fmt.Errorf("failed to get sign according to id : %v",err)
	}

	channel1Req := resmgmt.SaveChannelRequest{ChannelID:info.Channel1ID, ChannelConfigPath:info.Channel1Config, SigningIdentities:[]msp.SigningIdentity{org1AdminIdentity}}

	_,err = resOrg1MgmtClient.SaveChannel(channel1Req,resmgmt.WithRetry(retry.DefaultResMgmtOpts), resmgmt.WithOrdererEndpoint(info.OrdererOrgName))
	//resmgmt.WithRetry(retry.DefaultResMgmtOpts), resmgmt.WithOrdererEndpoint(info.OrdererOrgName)
	if err!=nil{
		return fmt.Errorf("fail to create channel %v",err)
	}

	info.Org1ResMgmt = resOrg1MgmtClient
	err = info.Org1ResMgmt.JoinChannel(info.Channel1ID,resmgmt.WithRetry(retry.DefaultResMgmtOpts), resmgmt.WithOrdererEndpoint(info.OrdererOrgName))
	//, resmgmt.WithRetry(retry.DefaultResMgmtOpts), resmgmt.WithOrdererEndpoint(info.OrdererOrgName)
	if err != nil {
		return fmt.Errorf("org3 Peers failed to join channel: %v", err)
	}

	fmt.Println("org1 peers join channel success.")
	return nil

}



func ChannelOneInstallAndInstantiateCC(sdk *fabsdk.FabricSDK, info *InitInfo) (*channel.Client, error) {
	fmt.Println("start install chaincode......")
	// creates new go lang chaincode package
	cc1Pkg, err := gopackager.NewCCPackage(info.ChaincodeOnePath, info.ChaincodeOneGoPath)
	if err != nil {
		return nil, fmt.Errorf("fail to create chaincode package: %v", err)
	}
	fmt.Println(info.ChaincodeOnePath)
	// contains install chaincode request parameters
	installCCReq := resmgmt.InstallCCRequest{Name: info.ChaincodeOneID, Path: info.ChaincodeOnePath, Version: ChaincodeVersion, Package: cc1Pkg}
	// allows administrators to install chaincode onto the filesystem of a peer
	_, err = info.Org1ResMgmt.InstallCC(installCCReq, resmgmt.WithRetry(retry.DefaultResMgmtOpts))
	if err != nil {
		return nil, fmt.Errorf("fail to install chaincode: %v", err)
	}

	fmt.Println("install chaincode successfully")
	fmt.Println("start instantiation......")

	//  returns a policy that requires one valid
	//ccPolicy := cauthdsl.SignedByAnyPeer([]string{"peer0.org1.example.com"})
	ccPolicy :=cauthdsl.SignedByAnyMember([]string{"Org1MSP"})

	instantiateCCReq := resmgmt.InstantiateCCRequest{Name: info.ChaincodeOneID, Path: info.ChaincodeOnePath, Version: ChaincodeVersion, Args: [][]byte{[]byte("init")}, Policy: ccPolicy}
	// instantiates chaincode with optional custom options (specific peers, filtered peers, timeout). If peer(s) are not specified
	_, err = info.Org1ResMgmt.InstantiateCC(info.Channel1ID, instantiateCCReq, resmgmt.WithRetry(retry.DefaultResMgmtOpts))
	if err != nil {
		return nil, fmt.Errorf("instantiation of chaincode failed: %v", err)
	}

	fmt.Println("instantiation of chaincode successfully")

	clientChannelContext := sdk.ChannelContext(info.Channel1ID, fabsdk.WithUser(info.Org1UserName), fabsdk.WithOrg(info.Org1Name))
	// returns a Client instance. Channel client can query chaincode, execute chaincode and register/unregister for chaincode events on specific channel.
	channelClient1, err := channel.New(clientChannelContext)
	if err != nil {
		return nil, fmt.Errorf("failed to create the client of channel application: %v", err)
	}

	fmt.Println("create channel one client successfully ，you can use this client to invoke chaincode to query or excute event.")

	return channelClient1, nil
}

func ChannelTwoInstallAndInstantiateCC(sdk *fabsdk.FabricSDK, info *InitInfo) (*channel.Client, error) {
	fmt.Println("start install chaincode......")
	// creates new go lang chaincode package
	cc2Pkg, err := gopackager.NewCCPackage(info.ChaincodeTwoPath, info.ChaincodeTwoGoPath)
	if err != nil {
		return nil, fmt.Errorf("fail to create chaincode package: %v", err)
	}

	// contains install chaincode request parameters
	installCCReq := resmgmt.InstallCCRequest{Name: info.ChaincodeTwoID, Path: info.ChaincodeTwoPath, Version: ChaincodeVersion, Package: cc2Pkg}
	// allows administrators to install chaincode onto the filesystem of a peer
	_, err = info.Org2ResMgmt.InstallCC(installCCReq, resmgmt.WithRetry(retry.DefaultResMgmtOpts))
	if err != nil {
		return nil, fmt.Errorf("fail to install chaincode: %v", err)
	}

	fmt.Println("install chaincode successfully")
	fmt.Println("start instantiation......")

	//  returns a policy that requires one valid
	ccPolicy := cauthdsl.SignedByAnyMember([]string{"Org2MSP"})

	instantiateCCReq := resmgmt.InstantiateCCRequest{Name: info.ChaincodeTwoID, Path: info.ChaincodeTwoPath, Version: ChaincodeVersion, Args: [][]byte{[]byte("init")}, Policy: ccPolicy}
	// instantiates chaincode with optional custom options (specific peers, filtered peers, timeout). If peer(s) are not specified
	_, err = info.Org2ResMgmt.InstantiateCC(info.Channel2ID, instantiateCCReq, resmgmt.WithRetry(retry.DefaultResMgmtOpts))
	if err != nil {
		return nil, fmt.Errorf("instantiation of chaincode failed: %v", err)
	}

	fmt.Println("instantiation of chaincode successfully")

	clientChannelContext := sdk.ChannelContext(info.Channel2ID, fabsdk.WithUser(info.Org2UserName), fabsdk.WithOrg(info.Org2Name))
	// returns a Client instance. Channel client can query chaincode, execute chaincode and register/unregister for chaincode events on specific channel.
	channelClient2, err := channel.New(clientChannelContext)
	if err != nil {
		return nil, fmt.Errorf("failed to create the client of channel application: %v", err)
	}

	fmt.Println("create channel two client successfully ，you can use this client to invoke chaincode to query or excute event.")

	return channelClient2, nil
}

func ChannelThreeInstallAndInstantiateCC(sdk *fabsdk.FabricSDK, info *InitInfo) (*channel.Client, error) {
	fmt.Println("开始安装链码......")
	// creates new go lang chaincode package
	cc3Pkg, err := gopackager.NewCCPackage(info.ChaincodeThreePath, info.ChaincodeThreeGoPath)
	if err != nil {
		return nil, fmt.Errorf("创建链码包失败: %v", err)
	}

	// contains install chaincode request parameters
	installCCReq := resmgmt.InstallCCRequest{Name: info.ChaincodeThreeID, Path: info.ChaincodeThreePath, Version: ChaincodeVersion, Package: cc3Pkg}
	// allows administrators to install chaincode onto the filesystem of a peer
	_, err = info.Org3ResMgmt.InstallCC(installCCReq, resmgmt.WithRetry(retry.DefaultResMgmtOpts))
	if err != nil {
		return nil, fmt.Errorf("安装链码失败: %v", err)
	}

	fmt.Println("指定的链码安装成功")
	fmt.Println("开始实例化链码......")

	//  returns a policy that requires one valid
	ccPolicy := cauthdsl.SignedByAnyMember([]string{"Org3MSP"})

	instantiateCCReq := resmgmt.InstantiateCCRequest{Name: info.ChaincodeThreeID, Path: info.ChaincodeThreePath, Version: ChaincodeVersion, Args: [][]byte{[]byte("init")}, Policy: ccPolicy}
	// instantiates chaincode with optional custom options (specific peers, filtered peers, timeout). If peer(s) are not specified
	_, err = info.Org3ResMgmt.InstantiateCC(info.Channel3ID, instantiateCCReq, resmgmt.WithRetry(retry.DefaultResMgmtOpts))
	if err != nil {
		return nil, fmt.Errorf("实例化链码失败: %v", err)
	}

	fmt.Println("链码实例化成功")

	clientChannelContext := sdk.ChannelContext(info.Channel3ID, fabsdk.WithUser(info.Org3UserName), fabsdk.WithOrg(info.Org3Name))
	// returns a Client instance. Channel client can query chaincode, execute chaincode and register/unregister for chaincode events on specific channel.
	channelClient3, err := channel.New(clientChannelContext)
	if err != nil {
		return nil, fmt.Errorf("创建应用通道客户端失败: %v", err)
	}

	fmt.Println("通道客户端创建成功，可以利用此客户端调用链码进行查询或执行事务.")

	return channelClient3, nil
}















