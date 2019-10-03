package routers

import (
	"demo_test_med/web/controllers"
	"github.com/astaxie/beego"
)

func init() {
    beego.Include(&controllers.IndexController{})
	beego.Include(&controllers.AddEhrController{})
	beego.Include(&controllers.ListEhrController{})
	beego.Include(&controllers.ListPomController{})
	beego.Include(&controllers.LoginController{})
	beego.Include(&controllers.VerifyPomController{})
	beego.Include(&controllers.ViewEhrController{})
	beego.Include(&controllers.TestController{})
	beego.Include(&controllers.RegisterController{})

}
