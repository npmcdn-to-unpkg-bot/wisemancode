package routers

import (
	"wisemancode/controllers"

	"github.com/astaxie/beego"
	//	"github.com/astaxie/beego/context"
)

func init() {
	beego.Router("/", &controllers.MainController{})
	beego.Router("/wx", &controllers.WXController{})
	beego.Router("/test", &controllers.WXController{}, "*:Test")

}
