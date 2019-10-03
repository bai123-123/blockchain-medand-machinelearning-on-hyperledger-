package main

import (

	"encoding/json"
	"hash/crc32"
	"strconv"

	"github.com/hyperledger/fabric/core/chaincode/shim"
	"github.com/hyperledger/fabric/protos/peer"
)

func PutInfo(stub shim.ChaincodeStubInterface, info UserInfo) ([]byte, bool) {



	b, err := json.Marshal(info)
	if err != nil {
		return nil, false
	}

	// 保存edu状态
	err = stub.PutState(info.EMRNo, b)
	if err != nil {
		return nil, false
	}

	return b, true
}

// 根据身份证号码查询信息状态
// args: entityID
func GetInfo(stub shim.ChaincodeStubInterface, EMRNO string) (UserInfo, bool)  {
	var info UserInfo
	// 根据身份证号码查询信息状态
	b, err := stub.GetState(EMRNO)
	if err != nil {
		return info, false
	}

	if b == nil {
		return info, false
	}

	// 对查询到的状态进行反序列化
	err = json.Unmarshal(b, &info)
	if err != nil {
		return info, false
	}

	// 返回结果
	return info, true
}



// 添加信息
// args: educationObject
// 身份证号为 key, Education 为 value
func (t *PharmacyChaincode) addInfo(stub shim.ChaincodeStubInterface, args []string) peer.Response {

	if len(args) != 2{
		return shim.Error("给定的参数个数不符合要求")
	}

	var info UserInfo
	err := json.Unmarshal([]byte(args[0]), &info)
	if err != nil {
		return shim.Error("反序列化信息时发生错误")
	}



	temp, _ := GetInfo(stub, info.EMRNo)
	if temp.MaxNum<= temp.CurrentNum {
		return shim.Error("number exceed")
	}

	_, bl := PutInfo(stub, info)
	if !bl {
		return shim.Error("保存信息时发生错误")
	}

	err = stub.SetEvent(args[1], []byte{})
	if err != nil {
		return shim.Error(err.Error())
	}

	return shim.Success([]byte("信息添加成功"))
}


func (t *PharmacyChaincode) updateInfo(stub shim.ChaincodeStubInterface, args []string) peer.Response {
	if len(args) != 2{
		return shim.Error("给定的参数个数不符合要求")
	}

	var info UserInfo
	err := json.Unmarshal([]byte(args[0]), &info)
	if err != nil {
		return  shim.Error("反序列化edu信息失败")
	}

	// 根据身份证号码查询信息
	result, bl := GetInfo(stub, info.EMRNo)
	if !bl{
		return shim.Error("根据身份证号码查询信息时发生错误")
	}

	result.Name = info.Name
	result.EMRNo = info.EMRNo
	result.Category = info.Category
	result.MaxNum = info.MaxNum
	result.CurrentNum = info.CurrentNum


	_, bl = PutInfo(stub, result)
	if !bl {
		return shim.Error("保存信息信息时发生错误")
	}

	err = stub.SetEvent(args[1], []byte{})
	if err != nil {
		return shim.Error(err.Error())
	}

	return shim.Success([]byte("信息更新成功"))
}

func (t *PharmacyChaincode) queryInfoByEMrID(stub shim.ChaincodeStubInterface, args []string) peer.Response {
	if len(args) != 1 {
		return shim.Error("给定的参数个数不符合要求")
	}

	// 根据身份证号码查询edu状态
	b, err := stub.GetState(args[0])
	if err != nil {
		return shim.Error("根据身份证号码查询信息失败")
	}

	if b == nil {
		return shim.Error("根据身份证号码没有查询到相关的信息")
	}

	// 对查询到的状态进行反序列化
	var info UserInfo
	err = json.Unmarshal(b, &info)
	if err != nil {
		return  shim.Error("反序列化edu信息失败")
	}

	// 获取历史变更数据
	iterator, err := stub.GetHistoryForKey(info.EMRNo)
	if err != nil {
		return shim.Error("根据指定的身份证号码查询对应的历史变更数据失败")
	}
	defer iterator.Close()

	// 迭代处理
	var historys []HistoryItem
	var hisInfo UserInfo
	for iterator.HasNext() {
		hisData, err := iterator.Next()
		if err != nil {
			return shim.Error("获取edu的历史变更数据失败")
		}

		var historyItem HistoryItem
		historyItem.TxId = hisData.TxId
		json.Unmarshal(hisData.Value, &hisInfo)

		if hisData.Value == nil {
			var empty UserInfo
			historyItem.User = empty
		}else {
			historyItem.User = hisInfo
		}

		historys = append(historys, historyItem)

	}

	info.Historys = historys

	// 返回
	result, err := json.Marshal(info)
	if err != nil {
		return shim.Error("序列化edu信息时发生错误")
	}
	return shim.Success(result)
}

func (t *PharmacyChaincode) createPushCodeByEMrID(stub shim.ChaincodeStubInterface, args []string) peer.Response {
	if len(args) != 1 {
		return shim.Error("给定的参数个数不符合要求")
	}

	// 根据身份证号码查询edu状态
	b, err := stub.GetState(args[0])
	if err != nil {
		return shim.Error("根据身份证号码查询信息失败")
	}

	if b == nil {
		return shim.Error("根据身份证号码没有查询到相关的信息")
	}

	v:=int(crc32.ChecksumIEEE(b))
	if v< 0{
		v = v * -1
	}
	type pushcode struct {
		code string
	}
	pushNum :=strconv.Itoa(v)
	pushStr := pushNum[-4]
	var pushCode pushcode
	pushCode.code = string(pushStr)

	result, err := json.Marshal(pushCode)
	if err != nil {
		return shim.Error("序列化edu信息时发生错误")
	}

	return shim.Success(result)
}