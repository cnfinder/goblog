package routers

import (
	"github.com/cnfinder/goblog/controllers"

	"github.com/astaxie/beego"
)

func init() {

	// 注册 beego 路由
	beego.Router("/", &controllers.MainController{})

	beego.Router("/home", &controllers.HomeController{})
}
