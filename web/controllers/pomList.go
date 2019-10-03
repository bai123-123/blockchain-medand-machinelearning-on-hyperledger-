package controllers

import (
	"demo_test_med/web/models"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)

type ListPomController struct {
	beego.Controller
}
// @router /pomList [get]
func (this *IndexController) ListPom(){

	name := this.GetSession("name")
	job :=this.GetSession("job")
	account :=this.GetSession("account")
	this.Data["name"]=name
	this.Data["job"]= job
	this.Data["account"] = account
	o:=orm.NewOrm()
	ehr := new(models.EHR)
	var ehrs []models.EHR
	qs := o.QueryTable(ehr)
	qs.All(&ehrs)


	this.Data["array"]= ehrs


	this.TplName="pomList.html"
}
