package controllers

import (
	"demo_test_med/web/models"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"

)

type RegisterController struct {
	beego.Controller
}
// @router /regist [get]
func (this *IndexController) ShowRegist(){
	this.TplName="register.html"
}

// @router /regist [post]
func (this *IndexController) Regist(){
	account :=this.GetString("account")
	pwd:=this.GetString("pwd")
	name :=this.GetString("name")
	isMed :=this.GetString("isMed")
	var TureOrfalse bool
	switch isMed {
	case "false":
		TureOrfalse = false
	case "true":
		TureOrfalse= true


	}
	o := orm.NewOrm()
	temp := models.User{

		Name:    name,
		Account: account,
		Pwd:     pwd,
		IsMed:   TureOrfalse,
	}
	o.Insert(&temp)
	this.TplName = "register.html"
}
