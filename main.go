package main

import (
	"log"

	"github.com/cnfinder/goblog/models"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)

func init() {
	// 注册数据库  在  main  之前执行
	models.RegisterDB()
}
func main() {

	// 读取配置文件
	uname := beego.AppConfig.String("uname")
	pwd := beego.AppConfig.String("pwd")
	log.Printf("uname=%s,pwd=%s", uname, pwd)

	// 开启 ORM 调试模式

	orm.Debug = true
	// 自动建表  默认的数据库 而不是default数据库
	orm.RunSyncdb("default", false, true)

	beego.Run(":8888")

}
