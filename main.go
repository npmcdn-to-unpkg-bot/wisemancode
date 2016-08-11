package main

import (
	_ "wisemancode/routers"
	//"wisemancode/wechat"

	"github.com/astaxie/beego"
)

func main() {
	//wechat.MainCallbacl()
	beego.Run()
}

//logs.SetLogger("multifile", ')
