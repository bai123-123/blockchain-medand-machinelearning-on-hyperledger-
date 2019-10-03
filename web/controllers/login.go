package controllers

import (
	"demo_test_med/web/models"
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)

type LoginController struct {
	beego.Controller
}
// @router / [get]
func (this *IndexController) ShowLogin(){
	this.TplName="login.html"
}


// @router /login [post]
func (this *IndexController) Login(){
	account := this.GetString("account")
	pwd := this.GetString("passwd")
	fmt.Println(account,pwd)
	o:=orm.NewOrm()
	user := models.User{}
	user.Account =account
	err := o.Read(&user,"Account")
	if err!=nil{
		this.Ctx.WriteString("wrong account")
		return
	}
	fmt.Println(user.Name)
	var JobType string
	fmt.Println(user.IsMed)
	if user.IsMed{
		JobType ="Pharmacist"
		this.SetSession("isMed",1)
	}else {
		JobType ="Doctor"
		this.SetSession("isMed",nil)
	}
	this.SetSession("name",user.Name)
	this.SetSession("account",user.Account)
	this.SetSession("job",JobType)

	this.Redirect("/index",302)
}