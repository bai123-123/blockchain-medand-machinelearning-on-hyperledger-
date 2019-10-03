package controllers


import (
	"demo_test_med/hospitalPubservice"
	"fmt"
	"github.com/astaxie/beego"
)

type excutePurchaseController struct {
	beego.Controller
}
// @router /finishPurchase [post]
func (this *IndexController) ExcutePurchase(){

	type response struct {
		fail bool
	}

	ehrNo := this.GetString("ehrNo")
	currentNum :=this.GetString("amountCurrent")

	fmt.Println(ehrNo,"ehrNooooo")
	fmt.Println(currentNum,"currentNummmm")




	temp1:=hospitalPubservice.Common{
		Name:           "",
		Gender:         "",
		Age:            "",
		BirthDate:      "",
		Contact:        "",
		Medicine:       "",
		QuantityFixed:   "",
		AmountFixed:     "",
		AmountCurrent: currentNum,
	}

	temp:=hospitalPubservice.EMR_common{
		EMRNo:      ehrNo,
		MedNo:      "",
		DoctorNo:   "",
		POM:        "",
		Date:       "",
		CommonInfo: temp1,
		VerifyDate: "",
	}

	errorResponse :=response{fail:true}

	result1, err1 :=PubServiceSetup.UpdateMedicineQuality(temp)
	if err1 != nil {
		fmt.Println(err1.Error())

		this.Ctx.WriteString("wrong ehrNO")
	}else if len(result1)==1{
		fmt.Println("wrong number does exist")
		//return fail information
		this.Ctx.WriteString("wrong ehrNO")

	}else {
		//return success information
		this.Data["json"]=&errorResponse
		this.ServeJSON()
		errorResponse.fail=false
	}

}
// @router /finishPurchase [get]
func (this *IndexController) ShowVerifyPage(){

	this.TplName="testFinish.html"

}
