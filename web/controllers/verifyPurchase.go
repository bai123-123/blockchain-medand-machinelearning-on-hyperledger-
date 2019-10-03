
package controllers

import (
	"demo_test_med/hospitalPubservice"
	"encoding/json"
	"fmt"
	"github.com/astaxie/beego"
)

type VerifyPurchaseController struct {
	beego.Controller
}
// @router /getPushcode [post]
func (this *IndexController) GetMedicineInformationAndPushcode(){
	//this.Redirect(" http://10.140.0.77:8080/index/",302)
	type response struct {
		Name string
		Pushcode string
		Medicine string
		AmountFixed string
		AmountCurrent string
		QuantityFixed string
	}

	type ehr_id struct {
		ehr_id string
	}


	//var ehr ehr_id

	////json数据封装到user对象中
	//err := json.Unmarshal(data, &ehr)
	//if err != nil {
	//	fmt.Println("json.Unmarshal is err:", err.Error())
	//}


	ehrNo := this.GetString("ehr_id")
	fmt.Println(ehrNo)
	fmt.Println("asdasdasdasdasd")
	result1, err1 :=PubServiceSetup.FindEmrInfoByEmrNo(ehrNo)
	result2, err2 :=PubServiceSetup.CreatePushCode(ehrNo)
	var pubEhr hospitalPubservice.EMR_common
	if err1 != nil {
		fmt.Println(err1.Error())

	}else if err2!=nil{
		fmt.Println(err2.Error())
	}else {

		json.Unmarshal(result1, &pubEhr)
		fmt.Println("根据EMR号码查询信息成功：")
		fmt.Println(pubEhr)
	}

	fmt.Println(len(result1))
	if len(result1)==1{
		fmt.Println("ehr number does exist")
		this.Ctx.WriteString("wrong ehrNO")

	}else {
		fmt.Println(pubEhr.CommonInfo.Medicine)

		fmt.Println(pubEhr.EMRNo)
		fmt.Println(result2)
		jsonResponse :=response{
			Name:          pubEhr.CommonInfo.Name,
			Pushcode:      string(result2),
			Medicine:      pubEhr.CommonInfo.Medicine,
			QuantityFixed: pubEhr.CommonInfo.QuantityFixed,
			AmountFixed:   pubEhr.CommonInfo.AmountFixed,
			AmountCurrent: pubEhr.CommonInfo.AmountCurrent,
		}
		this.Data["json"] = &jsonResponse
		this.ServeJSON()

	}
}


// @router /getPushcode [get]
func (this *IndexController) GetVerifyPage(){

	this.TplName="testForm.html"

}

