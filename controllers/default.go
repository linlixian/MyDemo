package controllers

import (
	"fmt"

	"github.com/astaxie/beego"
)

type MainController struct {
	beego.Controller
}

func (c *MainController) Map() map[interface{}]interface{} {
	return c.Data
}

// @router / [get]
func (c *MainController) Get() {
	// c.XSRFExpire = 7200
	// c.Data["xsrf_token"] = c.XSRFToken()
	// c.Data["xsrfdata"] = template.HTML(c.XSRFFormHTML())
	c.TplName = "test.html"
}

// @router /new_message [post]
func (c *MainController) Post() {
	get(c)
}

func get(c RepayInterface) {
	defer c.ServeJSON()

	// c.XsrfToken()
	fmt.Println("2222222222222")
	// c.XSRFToken()
	// b := c.CheckXSRFCookie()
	// fmt.Println("1111111111111111", b)
	msg := c.GetString("message")
	// _xsrf := c.GetString("_xsrf")
	switch data := c.(type) {
	case *MainController:
		fmt.Println("11111111111111111--------", data)
	}
	Data := c.Map()
	Data["json"] = map[string]interface{}{"ret": 200, "msg": msg}
}

type RepayInterface interface {
	GetString(key string, def ...string) string
	ServeJSON(encoding ...bool)
	Map() map[interface{}]interface{}
}
