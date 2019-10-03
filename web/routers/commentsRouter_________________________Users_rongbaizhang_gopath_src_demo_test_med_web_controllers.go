package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context/param"
)

func init() {

    beego.GlobalControllerRouter["demo_test_med/web/controllers:IndexController"] = append(beego.GlobalControllerRouter["demo_test_med/web/controllers:IndexController"],
        beego.ControllerComments{
            Method: "ShowLogin",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["demo_test_med/web/controllers:IndexController"] = append(beego.GlobalControllerRouter["demo_test_med/web/controllers:IndexController"],
        beego.ControllerComments{
            Method: "AddEhr",
            Router: `/addEhr`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["demo_test_med/web/controllers:IndexController"] = append(beego.GlobalControllerRouter["demo_test_med/web/controllers:IndexController"],
        beego.ControllerComments{
            Method: "ShowAddEhr",
            Router: `/addEhr`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["demo_test_med/web/controllers:IndexController"] = append(beego.GlobalControllerRouter["demo_test_med/web/controllers:IndexController"],
        beego.ControllerComments{
            Method: "ListEhr",
            Router: `/ehrList`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["demo_test_med/web/controllers:IndexController"] = append(beego.GlobalControllerRouter["demo_test_med/web/controllers:IndexController"],
        beego.ControllerComments{
            Method: "ShowVerifyPage",
            Router: `/finishPurchase`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["demo_test_med/web/controllers:IndexController"] = append(beego.GlobalControllerRouter["demo_test_med/web/controllers:IndexController"],
        beego.ControllerComments{
            Method: "ExcutePurchase",
            Router: `/finishPurchase`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["demo_test_med/web/controllers:IndexController"] = append(beego.GlobalControllerRouter["demo_test_med/web/controllers:IndexController"],
        beego.ControllerComments{
            Method: "GetMedicineInformationAndPushcode",
            Router: `/getPushcode`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["demo_test_med/web/controllers:IndexController"] = append(beego.GlobalControllerRouter["demo_test_med/web/controllers:IndexController"],
        beego.ControllerComments{
            Method: "GetVerifyPage",
            Router: `/getPushcode`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["demo_test_med/web/controllers:IndexController"] = append(beego.GlobalControllerRouter["demo_test_med/web/controllers:IndexController"],
        beego.ControllerComments{
            Method: "Index",
            Router: `/index`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["demo_test_med/web/controllers:IndexController"] = append(beego.GlobalControllerRouter["demo_test_med/web/controllers:IndexController"],
        beego.ControllerComments{
            Method: "Login",
            Router: `/login`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["demo_test_med/web/controllers:IndexController"] = append(beego.GlobalControllerRouter["demo_test_med/web/controllers:IndexController"],
        beego.ControllerComments{
            Method: "ListPom",
            Router: `/pomList`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["demo_test_med/web/controllers:IndexController"] = append(beego.GlobalControllerRouter["demo_test_med/web/controllers:IndexController"],
        beego.ControllerComments{
            Method: "Regist",
            Router: `/regist`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["demo_test_med/web/controllers:IndexController"] = append(beego.GlobalControllerRouter["demo_test_med/web/controllers:IndexController"],
        beego.ControllerComments{
            Method: "ShowRegist",
            Router: `/regist`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["demo_test_med/web/controllers:IndexController"] = append(beego.GlobalControllerRouter["demo_test_med/web/controllers:IndexController"],
        beego.ControllerComments{
            Method: "ShowVerifyPom",
            Router: `/verifyPom`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["demo_test_med/web/controllers:IndexController"] = append(beego.GlobalControllerRouter["demo_test_med/web/controllers:IndexController"],
        beego.ControllerComments{
            Method: "VerifyPom",
            Router: `/verifyPom`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["demo_test_med/web/controllers:IndexController"] = append(beego.GlobalControllerRouter["demo_test_med/web/controllers:IndexController"],
        beego.ControllerComments{
            Method: "ViewEhr",
            Router: `/viewEhr/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["demo_test_med/web/controllers:IndexController"] = append(beego.GlobalControllerRouter["demo_test_med/web/controllers:IndexController"],
        beego.ControllerComments{
            Method: "ViewPom",
            Router: `/viewPom`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["demo_test_med/web/controllers:TestController"] = append(beego.GlobalControllerRouter["demo_test_med/web/controllers:TestController"],
        beego.ControllerComments{
            Method: "Test",
            Router: `/test`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["demo_test_med/web/controllers:TestController"] = append(beego.GlobalControllerRouter["demo_test_med/web/controllers:TestController"],
        beego.ControllerComments{
            Method: "TestPost",
            Router: `/test`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

}
