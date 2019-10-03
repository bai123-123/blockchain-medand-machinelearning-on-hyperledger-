package controllers

import "github.com/astaxie/beego"

type TestController struct {
	beego.Controller
}

// @router /test [get]
func (c *TestController) Test(){
	c.Data["data"]= "testGet"
	c.TplName="test.html"
}
// @router /test [post]
func (c *TestController) TestPost(){
	c.Data["data"]= "testPost"
	c.TplName="test.html"
}