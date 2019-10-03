package main

import (
	"demo_test_med/irisesService"
	"time"
	"io/ioutil"
	"strings"
	"log"
	//"demo_test_med/org1Service"
	"demo_test_med/sdkInit"
	"fmt"
	"os"
)

const (
	configFile = "config.yaml"
	initialized = false
	CC1 = "hospitalPri"
	CC2 = "flower"
	SimpleCC3 = "testMongoDB12345"
)




var CurrentType string

func main()  {
	initInfo := &sdkInit.InitInfo{

		//Channel1ID: "orgonechannel",
		//Channel1Config: os.Getenv("GOPATH") + "/src/demo_test_med/fixtures/artifacts/orgone.channel.tx",
		//
		Channel2ID: "orgtwochannel",
		Channel2Config: os.Getenv("GOPATH") + "/src/demo_test_med/fixtures/artifacts/orgtwo.channel.tx",

		Channel3ID: "orgthreechannel",
		Channel3Config: os.Getenv("GOPATH") + "/src/demo_test_med/fixtures/artifacts/orgthree.channel.tx",

		//Org1Admin:"Admin",
		//Org1Name:"Org1",
		//
		Org2Admin:"Admin",
		Org2Name:"Org2",

		Org3Admin:"Admin",
		Org3Name:"Org3",

		OrdererOrgName: "peer0.org2.example.com",

		//ChaincodeOneID: CC1,
		//ChaincodeOneGoPath: os.Getenv("GOPATH"),
		//ChaincodeOnePath: "demo_test_med/chaincode/hospital_private",
		//Org1UserName:"Admin",
		//
		//ChaincodeTwoID: CC2,
		//ChaincodeTwoGoPath: os.Getenv("GOPATH"),
		//ChaincodeTwoPath: "demo_test_med/chaincode/IrisesInput",
		//Org2UserName:"Admin",

		ChaincodeThreeID: SimpleCC3,
		ChaincodeThreeGoPath: os.Getenv("GOPATH"),
		ChaincodeThreePath: "demo_test_med/chaincode/irises",
		Org3UserName:"Admin",
	}



	sdk, err := sdkInit.SetupSDK(configFile, initialized)
	if err != nil {
		fmt.Printf(err.Error())
		return
	}



	//err = sdkInit.CreateChannelone(sdk, initInfo)
	//if err != nil {
	//	fmt.Println(err.Error())
	//	return
	//}
	////
	//err = sdkInit.CreateChanneltwo(sdk, initInfo)
	//if err != nil {
	//	fmt.Println(err.Error())
	//	return
	//}

	err = sdkInit.CreateChannelthree(sdk, initInfo)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	//
	//channelClient1, err := sdkInit.ChannelOneInstallAndInstantiateCC(sdk, initInfo)
	//if err != nil {
	//	fmt.Println(err.Error())
	//	return
	//}
	//fmt.Println(channelClient1,"asdasdasdas")
	//
	//
	//channelClient2, err := sdkInit.ChannelTwoInstallAndInstantiateCC(sdk, initInfo)
	//if err != nil {
	//	fmt.Println(err.Error())
	//	return
	//}
	//fmt.Println(channelClient2)

	channelClient3, err := sdkInit.ChannelThreeInstallAndInstantiateCC(sdk, initInfo)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println(channelClient3)




	//===========================================//

	serviceSetup:=irisesService.IrisesServiceSetup{
		ChaincodeID: SimpleCC3,
		Client:      channelClient3,
	}
	//flower :=irisesService.Irises{
	//	FlowerID:     "001",
	//	Calyx_length: "1",
	//	Calyx_width:  "1",
	//	Petal_length: "1",
	//	Petal_width:  "1",
	//}


	//========================================================================================//

	type Trainer struct {
		Sepal_length float64
		Sepal_width float64
		Petal_length float64
		Petal_width float64
		Lable string
	}



	//======================//
	f, _ := os.Open("iris2.data")
	defer f.Close()
	content, _ := ioutil.ReadAll(f)
	//fmt.Println(content)
	s_content := string(content)
	fmt.Println(s_content)
	result,err :=serviceSetup.LoadTrainData(s_content)
	if err!=nil{
		fmt.Println("fail to load data :")
	}else {
		//var flower irisesService.Irises
		//json.Unmarshal(result,&flower)
		fmt.Println(string(result))
	}
	time.Sleep(9000)
   //
   //
   //
   //
   //
   //
   //
   //var Calyx_length string
	//var Calyx_width  string
	//var Petal_length string
	//var Petal_width  string
	//i:=0
	//for {
	//	fmt.Println("please input Irises data....")
	//	fmt.Scan(&Calyx_length,&Calyx_width,&Petal_length,&Petal_width)
   //
	//	example:= irisesService.Irises{
	//		FlowerID:     string(i),
	//		Calyx_length: Calyx_length,
	//		Calyx_width:  Calyx_width,
	//		Petal_length: Petal_length,
	//		Petal_width:  Petal_width,
	//	}
	//	results,_ :=serviceSetup.AddResult(example)
	//	if i>0&&string(results)!=CurrentType{
	//		fmt.Println("detect change")
	//		changeChainCode()
	//		break
	//	}
	//	CurrentType=string(results)
	//	fmt.Println(string(results))
	//	i++
	//}
//=================================================================================================================//
//	serviceSetup:=fnnService.FnnServiceSetup{
//		ChaincodeID: SimpleCC3,
//		Client:      channelClient3,
//	}
//	f, _ := os.Open("INPUT.txt")
//	defer f.Close()
//	content, _ := ioutil.ReadAll(f)
//	s_content := string(content)
//	_, err = serviceSetup.LoadInputData(s_content)
//	if err!=nil{
//		fmt.Println("fail to load input data :")
//		}
//
//
//	f1, _ := os.Open("output.txt")
//	defer f1.Close()
//	content1, _ := ioutil.ReadAll(f1)
//	s_content1 := string(content1)
//	_, err = serviceSetup.LoadOutputData(s_content1)
//	if err!=nil{
//		fmt.Println("fail to load output data :")
//	}
//
//	_, err = serviceSetup.ClassifyData()
//	if err!=nil{
//		fmt.Println("fail to classify data :")
//	}


	//var newData [][]float64
	////
	//f1, _ := os.Open("with_label_features.csv")
	//defer f1.Close()
	//content1, _ := ioutil.ReadAll(f)
	//s_content1 := string(content1)
	//lines := strings.Split(s_content1, "\n")
	//
	//line :=lines[1]
	//
	//	line = strings.TrimRight(line, "\r\n")
	//	tup := strings.Split(line, ",")
	//	time := tup[2:len(tup)-1]
	//	user := tup[len(tup)-1]
	//	X := make([]float64, 0)
	//	var x string
	//	x = strings.Join(time,x)
	//	f_x, _ := strconv.ParseFloat(x, 64)
	//	f_y, _ := strconv.ParseFloat(user, 64)
	//	X = append(X, f_x)
	//	X = append(X,f_y)
	//	newData =append(newData,X)
	//
	//newData =append(newData,X)
	defer sdk.Close()


}

func changeChainCode() {
	input, err := ioutil.ReadFile("chaincode/irises/irisesCC.go")
	if err != nil {
		log.Fatalln(err)
	}

	lines := strings.Split(string(input), "\n")

	for i, line := range lines {
		if strings.Contains(line, "//**??") {
			lines[i] = "    flower.IrisType = output"
		}else if strings.Contains(line, "//**!!"){
			lines[i] = "    IrisType string"
		}
	}
	output := strings.Join(lines, "\n")
	err = ioutil.WriteFile("chaincode/irises/irisesCC.go", []byte(output), 0644)
	if err != nil {
		log.Fatalln(err)
	}
}

