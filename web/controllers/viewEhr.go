package controllers

import (
	"demo_test_med/hospitalPriservice"
	"encoding/json"
	"fmt"
	"github.com/astaxie/beego"
)

type ViewEhrController struct {
	beego.Controller
}



// @router /viewEhr/ [get]
func (this *IndexController) ViewEhr(){
	var emr hospitalPriservice.EMR_pri
	id:= this.GetString("id")
	fmt.Println(id)
	result, err :=PriServiceSetup.FindEmrInfoByEmrNo(id)
	if err != nil {
		fmt.Println(err.Error())
	} else {

		json.Unmarshal(result, &emr)
		fmt.Println("根据EMR号码查询信息成功：")
		fmt.Println(emr)
	}
	fmt.Println(len(result))
	this.Data["name"] =emr.PPatient.Name
	this.Data["gender"]= emr.PPatient.Gender
	this.Data["age"] = emr.PPatient.Age
	this.Data["bitrhDate"] =emr.PPatient.BirthDate
	this.Data["IdNo"] = emr.PPatient.IdCardNo
	this.Data["address"]=emr.PPatient.Address
	this.Data["ContactNum"] =emr.PPatient.ContactNumber
	this.Data["EcontractNum"] =emr.PPatient.EmergencyContact
	this.Data["doctorName"] = this.GetSession("name")
	this.Data["date"] =emr.Date
	this.Data["Doctor"]=emr.Doctor
	this.Data["EhrNo"]=emr.EMRNo
	this.Data["PMCD"]=emr.PCD
	this.Data["DD"]=emr.DD
	this.Data["allergies"]=emr.Medicine
	this.Data["amount"]=emr.Amount
	this.Data["quantity"]=emr.Quantity


	this.TplName="viewEhr.html"
}
