package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context/param"
)

func init() {

    beego.GlobalControllerRouter["MyProject/controllers:MainController"] = append(beego.GlobalControllerRouter["MyProject/controllers:MainController"],
        beego.ControllerComments{
            Method: "Get",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["MyProject/controllers:MainController"] = append(beego.GlobalControllerRouter["MyProject/controllers:MainController"],
        beego.ControllerComments{
            Method: "Post",
            Router: `/new_message`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

}
