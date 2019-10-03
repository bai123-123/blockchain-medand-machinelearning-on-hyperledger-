package controllers

import (
	"demo_test_med/web/models"
	"fmt"
	"strconv"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"time"
	"demo_test_med/hospitalPriservice"


)

type AddEhrController struct {
	beego.Controller
}




// @router /addEhr [get]
func (this *IndexController) ShowAddEhr(){
	namee := this.GetSession("name")
	job :=this.GetSession("job")
	account :=this.GetSession("account")
	this.Data["name"]=namee
	this.Data["job"]= job
	this.Data["account"] = account

	this.TplName="addEhr.html"


}

// @router /addEhr [post]
func (this *IndexController) AddEhr(){
	name:=this.GetString("name")
	gender:=this.GetString("gender")
	age:=this.GetString("age")
	birthDate:=this.GetString("birthDate")
	idCardNo:=this.GetString("idCardNo")
	maritalState:=this.GetString("maritalState")
	address:=this.GetString("address")
	contactNum:=this.GetString("contactNum")
	econtactNum:=this.GetString("econtactNum")
	email:=this.GetString("email")
	MP:=this.GetString("MP")
	PCD:=this.GetString("PCD")
	PMH:=this.GetString("PMH")
	DD:=this.GetString("DD")
	allergies:=this.GetString("allergies")
	quantity:=this.GetString("quantity")
	amount:=this.GetString("amount")
	account :=this.GetSession("account")
	namee := this.GetSession("name")
	currentTime :=time.Now()
	stringTime :=currentTime.String()


	var No string
	o := orm.NewOrm()
	temp := models.EHR{

		Name:               name,
		Medical_department: MP,
		Age:                age,
		Time:              currentTime,
	}
	o.Insert(&temp)
	o.Raw("select max(e_h_r__n_o) from e_h_r").QueryRow(&No)
	num ,_:= strconv.Atoi(No)
	fmt.Println(num)

	temp1 := hospitalPriservice.Patient{
		Name:             name,
		Gender:           gender,
		Age:              age,
		BirthDate:        birthDate,
		EmergencyContact: econtactNum,
		MartialState:       maritalState,
		Address:            address,
		ContactNumber:    contactNum,
		Email:            email,
		IdCardNo:           idCardNo,
	}
	temp_ehr :=hospitalPriservice.EMR_pri{
		EMRNo:     strconv.Itoa(num) ,
		Medicine:          allergies,
		Amount:            amount,
		Quantity:          quantity,
		DoctorNo:           account.(string),
		Doctor:             namee.(string),
		Date:        stringTime,
		PCD:           PCD,
		PMH:           PMH,
		DD:            DD,
		PPatient: temp1,
		Medical_department: MP,

	}

	fmt.Println(temp_ehr)
	msg, err :=PriServiceSetup.SavePriEMR(temp_ehr)
	if err != nil {
		fmt.Println(err.Error())
	}else {
		fmt.Println("信息发布成功, 交易编号为: " + msg)
	}
	this.Redirect("/ehrList",302)


}