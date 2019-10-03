package controllers

import (
	"demo_test_med/hospitalPriservice"
	"demo_test_med/hospitalPubservice"
	"demo_test_med/web/models"
	"encoding/json"
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"strconv"
	"time"
)

type VerifyPomController struct {
	beego.Controller
}

var emr hospitalPriservice.EMR_pri

// @router /verifyPom [get]
func (this *IndexController) ShowVerifyPom(){
	id:= this.GetString("id")
	result, err :=PriServiceSetup.FindEmrInfoByEmrNo(id)
	if err != nil {
		fmt.Println(err.Error())
	} else {

		json.Unmarshal(result, &emr)
		fmt.Println("根据EMR号码查询信息成功：")
		fmt.Println(emr)
	}
	this.Data["name"] =emr.PPatient.Name
	this.Data["gender"]= emr.PPatient.Gender
	this.Data["age"] = emr.PPatient.Age
	this.Data["bitrhDate"] =emr.PPatient.BirthDate
	this.Data["date"] =emr.Date
	this.Data["Doctor"]=emr.Doctor
	this.Data["EhrNo"]=emr.EMRNo
	this.Data["PMCD"]=emr.PCD
	this.Data["medicine"]=emr.Medicine
	this.Data["amount"]=emr.Amount
	this.Data["quantity"]=emr.Quantity
	this.Data["EHRid"] = emr.EMRNo


	this.TplName="verifyPom.html"
}

// @router /verifyPom [post]
func (this *IndexController) VerifyPom(){

	pom:= this.GetString("POM")
	id := this.GetString("ehrID")





	fmt.Println(id)
	result, err :=PriServiceSetup.FindEmrInfoByEmrNo(id)
	if err != nil {
		fmt.Println(err.Error())
	} else {

		json.Unmarshal(result, &emr)

		fmt.Println(emr)
	}
	account :=this.GetSession("account")

	temp1 :=hospitalPubservice.Common{
		Name:             emr.PPatient.Name,
		Gender:           emr.PPatient.Gender,
		Age:              emr.PPatient.Age,
		BirthDate:        emr.PPatient.BirthDate,
		Contact:          emr.PPatient.ContactNumber,
		Medicine:        emr.Medicine,
		AmountCurrent:    "0",
		AmountFixed:     emr.Amount,
		QuantityFixed:    emr.Quantity,
	}

	temp := hospitalPubservice.EMR_common{
		EMRNo:      emr.EMRNo,
		MedNo:      account.(string),
		DoctorNo:   emr.DoctorNo,
		POM:        pom,
		Date:       emr.Date,
		VerifyDate: time.Now().String(),
		CommonInfo: temp1,
	}

	PubServiceSetup.SavePubEMR(temp)

	ehrNo ,_:=strconv.Atoi(id)

	o := orm.NewOrm()
	ehr := models.EHR{EHR_NO: ehrNo}
	if o.Read(&ehr) == nil {
		ehr.IsVerify = true
		if _, err := o.Update(&ehr); err == nil {
			fmt.Println("verify success")
		}
	}



	this.Redirect("/pomList",302)
}
