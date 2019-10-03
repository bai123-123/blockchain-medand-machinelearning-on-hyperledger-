package models

import (
	"demo_test_med/hospitalPriservice"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
	"time"
)
type User struct {
	Id int
	Name string
	Account string
	Pwd string
	IsMed bool
	IsDrugShop bool

}

type EHR struct {
	EHR_NO int `orm:"pk;auto"`
	Name string
	Medical_department string
	Age string
	Time time.Time
	IsVerify bool
}





var PriServiceSetup hospitalPriservice.PrivateServiceSetup

func init()  {




	orm.RegisterDataBase("default","mysql","root:1796mncw@tcp(127.0.0.1:3306)/test?charset=utf8")

	orm.RegisterModel(new(User),new(EHR))

	orm.RunSyncdb("default",false,true)
}