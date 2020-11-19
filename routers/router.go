package routers

import (
	"forum/controllers"
	"github.com/astaxie/beego"
)

func init() {
    beego.Router("/", &controllers.MainController{},"get:Get;post:Post")
	beego.Router("/test/sign", &controllers.SignController{},"get:Get;post:Post")
	beego.Router("/test/log", &controllers.LogController{},"get:Get;post:Post")

}
