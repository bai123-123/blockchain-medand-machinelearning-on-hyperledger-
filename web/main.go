package main

import (
	"demo_test_med/hospitalPriservice"
	"demo_test_med/hospitalPubservice"
	"demo_test_med/sdkInit"
	_ "demo_test_med/web/models"
	_ "demo_test_med/web/routers"
	"encoding/json"
	"os"

	"fmt"
	"github.com/astaxie/beego"
	"demo_test_med/web/controllers"

)
const (
	configFile = "config.yaml"
	initialized = false
	CC1 = "hospitalPri"
	CC2 = "hospitalPubASD"
	SimpleCC3 = "simpleCC3"
)
var AAA int

func main() {
	beego.BConfig.WebConfig.Session.SessionOn = true
	fmt.Println("ssssssss")
	initInfo := &sdkInit.InitInfo{

		Channel1ID: "orgonechannel",
		Channel1Config: os.Getenv("GOPATH") + "/src/demo_test_med/fixtures/artifacts/orgone.channel.tx",

		Channel2ID: "orgtwochannel",
		Channel2Config: os.Getenv("GOPATH") + "/src/demo_test_med/fixtures/artifacts/orgtwo.channel.tx",

		Channel3ID: "orgthreechannel",
		Channel3Config: os.Getenv("GOPATH") + "/src/demo_test_med/fixtures/artifacts/orgthree.channel.tx",

		Org1Admin:"Admin",
		Org1Name:"Org1",

		Org2Admin:"Admin",
		Org2Name:"Org2",

		Org3Admin:"Admin",
		Org3Name:"Org3",

		OrdererOrgName: "peer0.org2.example.com",

		ChaincodeOneID: CC1,
		ChaincodeOneGoPath: os.Getenv("GOPATH"),
		ChaincodeOnePath: "demo_test_med/chaincode/hospital_private",
		Org1UserName:"Admin",

		ChaincodeTwoID: CC2,
		ChaincodeTwoGoPath: os.Getenv("GOPATH"),
		ChaincodeTwoPath: "demo_test_med/chaincode/hospital_public",
		Org2UserName:"Admin",

		ChaincodeThreeID: SimpleCC3,
		ChaincodeThreeGoPath: os.Getenv("GOPATH"),
		ChaincodeThreePath: "demo_test_med/chaincode/org1Chaincode",
		Org3UserName:"Admin",
	}



	sdk, err := sdkInit.SetupSDK(configFile, initialized)
	if err != nil {
		fmt.Printf(err.Error())
		return
	}



	err = sdkInit.CreateChannelone(sdk, initInfo)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	//
	err = sdkInit.CreateChanneltwo(sdk, initInfo)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	channelClient1, err := sdkInit.ChannelOneInstallAndInstantiateCC(sdk, initInfo)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println(channelClient1,"asdasdasdas")


	channelClient2, err := sdkInit.ChannelTwoInstallAndInstantiateCC(sdk, initInfo)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println(channelClient2)

	defer sdk.Close()
	controllers.PriServiceSetup=hospitalPriservice.PrivateServiceSetup{
		ChaincodeID:CC1,
		Client:channelClient1,
	}
	temp1 := hospitalPriservice.Patient{
		Name:             "name",
		Gender:           "gender",
		Age:              "age",
		BirthDate:        "birthdate",
		EmergencyContact: "Econtact",
		MartialState:       "ma",
		Address:            "address",
		ContactNumber:    "contactNum",
		Email:            "abc@abc.com",
		IdCardNo:           "IDno",
	}


	temp :=hospitalPriservice.EMR_pri{
		EMRNo:      "123",

		PPatient:temp1,
		PCD:                "pcd",
		PMH:                "pmh",
		DD:                 "dd",
		Medical_department: "med_depart",

		Medicine:        "medicine",
		Quantity:          "quality",
		Amount:            "amount",
		DoctorNo:         "DoctorNo",
		Doctor:           "doctor",

		Date:             "date",

	}
	msg, err := controllers.PriServiceSetup.SavePriEMR(temp)
	if err != nil {
		fmt.Println(err.Error())
	}else {
		fmt.Println ( msg)
	}
	result, err := controllers.PriServiceSetup.FindEmrInfoByEmrNo("123")
	if err != nil {
		fmt.Println(err.Error())
	} else {
		var emr hospitalPriservice.EMR_pri
		json.Unmarshal(result, &emr)
		fmt.Println("success：")
		fmt.Println(emr)
	}

	controllers.PubServiceSetup=hospitalPubservice.PublicServiceSetup{
		ChaincodeID: CC2,
		Client:      channelClient2,
	}
	test1 := hospitalPubservice.Common{
		Name:             "name",
		Gender:           "gender",
		Age:              "age",
		BirthDate:        "birthdate",
		Contact: "econtact",
		Medicine:        "allergies",
		QuantityFixed:          "ammount",
		AmountCurrent:     "current",
		AmountFixed:       "",
	}
	test := hospitalPubservice.EMR_common{
		EMRNo:      "emrno",
		MedNo:      "medno",
		DoctorNo:   "doctorno",
		POM:        "pom",
		Date:       "date",
		CommonInfo: test1,
	}
	msg1, err1 := controllers.PubServiceSetup.SavePubEMR(test)
	if err1 != nil {
		fmt.Println(err1.Error())
	}else {
		fmt.Println ( msg1)
	}
	result1, err := controllers.PubServiceSetup.FindEmrInfoByEmrNo("234")
	if err != nil {
		fmt.Println(err.Error())
	} else {
		var emr hospitalPubservice.EMR_common
		json.Unmarshal(result1, &emr)

		fmt.Println(emr)
	}
	defer sdk.Close()






	beego.Run()



}

