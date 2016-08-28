package main

import (
	"wisemancode/log"
	_ "wisemancode/routers"

	"wisemancode/wechat/model"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

//"wisemancode/wechat"
func init() {
	log.Logger.Info("初始化数据库=======start============")
	orm.RegisterDriver("mysql", orm.DRMySQL)
	orm.RegisterDataBase("default", "mysql", "root:root@tcp(127.0.0.1:3306)/wx?charset=utf8")
	//开启开发模式
	orm.Debug = true
	//需要对数据库进行整理
	log.Logger.Info("初始化数据库=======end=================")
}
func main() {
	//wechat.MainCallbacl()
	u := model.NewSubscribeByAgrs("HEYIHSI", "heyishi", "event", "sub", 1)
	u.AddSubscribe()
	beego.Run()

}

//logs.SetLogger("multifile", ')
