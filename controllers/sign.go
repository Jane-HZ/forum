package controllers

import (
	"github.com/astaxie/beego"
	"forum/models"
)


type ResponseJson struct {
	State int
	Message string
	Data interface{}
}

type SignController struct {
	beego.Controller
}

func (c *SignController) Get() {
	c.TplName = "sign.tpl"
}

func (c *SignController) Post() {
	/*	u := User{}
		if err := c.ParseForm(&u); err != nil {

		} else {
			c.Ctx.WriteString("Welcome！" + u.Name)
		}
		name := strings.Trim(c.GetString("name"), "\t\r\n ")
		password := strings.Trim(c.GetString("password"), "\t\r\n ")
		if name == "" || password == "" {
			logs.Info("数据不能为空")
			c.Ctx.WriteString("用户名或密码不能为空")
			return
		}*/
	/*
		name := c.GetSession("Name")
		if name == nil {
			c.Redirect("/test/sign",302)
			c.Ctx.WriteString("用户名不能为空")
			return
		}
		c.Data["Name"] = name
	}*/
	user := models.User{}
	user.Name = c.Input().Get("name")
	user.Password = c.Input().Get("password")
	response := ResponseJson{State: 0, Message: "ok"}
	if user.Name == "" || user.Password == "" {
		response.State = 500
		response.Message = "用户名或密码不能为空"
	} else {
		if  _,err := user.SaveInformation(); err != nil {
			response.State = 501
			response.Message = "保存失败"
		} else {
			c.SetSession("Name",user.Name)
		}
	}
	c.Data["json"] = response
	c.ServeJSON()

}
