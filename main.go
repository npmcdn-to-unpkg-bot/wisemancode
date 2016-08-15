package main

import (
	_ "wisemancode/routers"

	"github.com/astaxie/beego"
)

//"wisemancode/wechat"

func main() {
	//wechat.MainCallbacl()
	beego.Run()

}

//logs.SetLogger("multifile", ')
