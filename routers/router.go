package routers

import (
	"MyProject/controllers"

	"github.com/astaxie/beego"
)

func init() {
	//=======================================解决自动路由自身bug====================================
	exceptMethod := []string{"GetInt8", "GetInt16", "GetInt32", "GetInt64", "GetFiles",
		"XSRFToken", "CheckXSRFCookie", "XSRFFormHTML", "HandlerFunc", "Mapping", "URLMapping", "renderTemplate", "URLFor"}
	for _, fn := range exceptMethod {
		beego.ExceptMethodAppend(fn)
	}
	beego.Include(&controllers.MainController{})
}
