package main

import (
	"wisemancode/log"
	_ "wisemancode/routers"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

//"wisemancode/wechat"
func init() {
	log.Logger.Info("初始化数据库=======start============")
	orm.RegisterDriver("mysql", orm.DRMySQL)
	//orm.RegisterDataBase("default", "mysql", "root:root@tcp(192.168.224.128:3306)/cloudta?charset=utf8")
	orm.RegisterDataBase("default", "mysql", "root:root@tcp(127.0.0.1:3306)/wx?charset=utf8")
	log.Logger.Info("初始化数据库=======end=================")
}
func main() {
	//wechat.MainCallbacl()
	beego.Run()

}

//logs.SetLogger("multifile", ')
