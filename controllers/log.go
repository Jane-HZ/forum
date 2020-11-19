package controllers

import (
	"github.com/astaxie/beego"
	"forum/models"
)


type LogController struct {
	beego.Controller
}

func (c *LogController) Get() {
	c.TplName = "log.tpl"
}

	func(c *LogController) Post() {
		user := models.User{}
		user.Name = c.Input().Get("name")
		user.Password = c.Input().Get("password")
		response := ResponseJson{State: 0, Message: "ok"}
		if user.Name == "" || user.Password == "" {
			response.State = 500
			response.Message = "用户名或密码不能为空"
		} else {
			if _, err := user.GetInformation(); err != nil  {
				response.State = 501
				response.Message =  "该用户尚未注册"
			} else {
				c.SetSession("Name", user.Name)
			}
		}
		c.Data["json"] = response
		c.ServeJSON()
	}
