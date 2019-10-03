package controllers

import (
	"demo_test_med/hospitalPubservice"
	"encoding/json"
	"fmt"
	"github.com/astaxie/beego"
)

type ViewPomController struct {
	beego.Controller
}


// @router /viewPom [get]
func (this *IndexController) ViewPom(){
	id:= this.GetString("id")
	fmt.Println(id)
	result, err :=PubServiceSetup.FindEmrInfoByEmrNo(id)
	var pubEhr hospitalPubservice.EMR_common
	if err != nil {
		fmt.Println(err.Error())
	} else {

		json.Unmarshal(result, &pubEhr)
		fmt.Println("根据EMR号码查询信息成功：")
		fmt.Println(pubEhr)
	}
	this.Data["name"] =pubEhr.CommonInfo.Name
	this.Data["gender"]= pubEhr.CommonInfo.Gender
	this.Data["age"] = pubEhr.CommonInfo.Age
	this.Data["bitrhDate"] =pubEhr.CommonInfo.BirthDate
	this.Data["date"] =pubEhr.Date

	this.Data["EhrNo"]=pubEhr.EMRNo
	this.Data["POM"]=pubEhr.POM
	this.Data["medicine"]=pubEhr.CommonInfo.Medicine
	this.Data["amount"]=pubEhr.CommonInfo.AmountFixed
	this.Data["quantity"]=pubEhr.CommonInfo.QuantityFixed



	this.TplName="viewPom.html"
}
