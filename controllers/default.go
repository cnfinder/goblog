package controllers

import (
	"github.com/astaxie/beego"
)

type MainController struct {
	beego.Controller
}

type User struct {
	UserId   string
	UserName string
	Name     string
	age      int
}

func (c *MainController) Get() {
	c.Data["Website"] = "goblog"
	c.Data["Email"] = "171721626@qq.com"
	c.TplName = "index.tpl"
	c.Data["A"] = "a"

	c.Data["user"] = User{UserId: "10001", UserName: "dragon", Name: "少龙", age: 10}
	c.Data["condition"] = false

	arr := []int{1, 2, 3, 4, 5}
	c.Data["arr"] = arr
	c.Data["html"] = "<div style=color:#ff0000>andrew</div>"

}
