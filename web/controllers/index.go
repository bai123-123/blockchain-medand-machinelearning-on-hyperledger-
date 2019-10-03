package controllers

import (
	"fmt"
	"github.com/astaxie/beego"
)

type IndexController struct {
	beego.Controller
}
// @router /index [get]
func (this *IndexController) Index(){
	name := this.GetSession("name")
	job :=this.GetSession("job")
	account :=this.GetSession("account")
	this.Data["name"]=name
	this.Data["job"]= job
	this.Data["account"] = account
	fmt.Println(name)
	if this.GetSession("isMed") ==nil{
		this.TplName ="index.html"
		return
	}else {
		this.TplName="index1.html"
		return
	}
	this.Ctx.WriteString("wrong")
}

