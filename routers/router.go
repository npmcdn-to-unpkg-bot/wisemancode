package routers

import (
	"wisemancode/controllers"

	"github.com/astaxie/beego"
	//	"github.com/astaxie/beego/context"
)

func init() {
	beego.Router("/", &controllers.MainController{})
	beego.Router("/wx", &controllers.WXController{})
	//beego.Get("/wx", Get_)

}

//var Get_ = func(ctx *context.Context) {
//	ctx.Output.Body([]byte("基本Get路由实践。"))
//}
