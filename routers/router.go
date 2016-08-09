package routers

import (
	"wisemancode/controllers"
	"github.com/astaxie/beego"
)

func init() {
    beego.Router("/", &controllers.MainController{})
	beego.Router("/wx",&controllers.)
}
