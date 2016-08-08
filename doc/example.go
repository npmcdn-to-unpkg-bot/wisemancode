package doc

import (
	"beegorouter/controllers"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context"
)

func init() {
	// bee 工具自动生成的固定路由
	beego.Router("/", &controllers.MainController{})

	// beego路由实践，参考文档：http://beego.me/docs/mvc/controller/router.md

	/**
	 * 基础路由实践
	 */

	/**
	 * 基本Get路由实践
	 * 注意事项：需要导入包：github.com/astaxie/beego/context
	 *           否则【context.Context】会编译出错
	 * 测试用例：浏览器里输入：http://localhost:8080/get
	 * 输出结果：在浏览器页面里显示：基本Get路由实践。
	 */
	beego.Get("/get", func(ctx *context.Context) {
		ctx.Output.Body([]byte("基本Get路由实践。"))
	})

	/**
	 * 基本Post路由实践
	 * 注意事项：需要修改或是追加html页面，从页面提交一个post请求
	 * 测试用例：修改index.tpl模板，追加一按钮，指定method="post"，点击该按钮
	 * 输出结果：基本Post路由实践。
	 */
	beego.Post("/post", func(ctx *context.Context) {
		ctx.Output.Body([]byte("基本Post路由实践。"))
	})

	/**
	 * 正则路由实践
	 */

	/**
	  * 正则路由实践
	  * 注意事项：1. 第二个参数需要传递一个控制器：Controller
	  *           2. 【?:id】 中的id也可以换成其他字符，如：【?:abc】
	  * 测试用例：能 匹 配：/api ; /api/123 ; /api/abc ; /api/abc.html ; /api/abc/
	              不能匹配：/api/123/456 ：即/api/之后只能再接一个参数
	  * 输出结果：具体输出结果需要看第二个参数的Get()方法所指定的模板和数据。
	*/
	beego.Router("/api/?:id", &controllers.MainController{})

	/**
	  * 正则路由实践
	  * 注意事项：【:id】前面没有了 ?
	  * 测试用例：能 匹 配：/api/123 ; /api/abc ; /api/abc.html ; /api/abc/
	              不能匹配：/api ; /api/123/456 ：即/api/之后只能再接一个参数
	  * 输出结果：具体输出结果需要看第二个参数的Get()方法所指定的模板和数据。
	*/
	beego.Router("/api/:id", &controllers.MainController{})

	/**
	  * 正则路由实践
	  * 注意事项：【:id】前面没有了 ?
	  * 测试用例：能 匹 配：/api/123 ; /api/0
	              不能匹配：/api ; /api/123/456 ; /api/abc ; /api/123.html
	  * 输出结果：具体输出结果需要看第二个参数的Get()方法所指定的模板和数据。
	*/
	beego.Router("/api/:id:int", &controllers.MainController{})

	/**
	  * 正则路由实践
	  * 注意事项：:id([0-9]+) 之中有个 + 号。有+号时表示可以是多位数字，无+号表示就只能匹配一位数字
	  * 测试用例：能 匹 配：/api/123 ; /api/0
	              不能匹配：/api ; /api/abc ; /api/123/456 ：即/api/之后只能再接一个参数
	  * 输出结果：具体输出结果需要看第二个参数的Get()方法所指定的模板和数据。
	*/
	beego.Router("/api/:id([0-9]+)", &controllers.MainController{})

	/**
	  * 正则路由实践
	  * 注意事项：:[\\w] 之中的w不能替换为别的字符
	  * 测试用例：能 匹 配：/api/123 ; /api/hezhixiong
	              不能匹配：/api ; /api/abc.html
	  * 输出结果：具体输出结果需要看第二个参数的Get()方法所指定的模板和数据。
	*/
	beego.Router("/api/:username([\\w]+)", &controllers.MainController{})

	/**
	  * 正则路由实践
	  * 注意事项：: :username:string表示 username为string型
	  * 测试用例：能 匹 配：/api/123 ; /api/hezhixiong
	              不能匹配：/api ; /api/abc.html
	  * 输出结果：具体输出结果需要看第二个参数的Get()方法所指定的模板和数据。
	*/
	beego.Router("/api/:username:string", &controllers.MainController{})

	/**
	  * 正则路由实践
	  * 注意事项：:无
	  * 测试用例：能 匹 配：/api/123 ; /api/hezhixiong ; /api/abc.html ; /api/abc/123/efg/ddd
	              不能匹配：/api
	  * 输出结果：具体输出结果需要看第二个参数的Get()方法所指定的模板和数据。
	*/
	beego.Router("/api/*.*", &controllers.MainController{})

	/**
	  * 正则路由实践
	  * 注意事项：:无
	  * 测试用例：能 匹 配：/api/123 ; /api/hezhixiong ; /api/abc.html ; /api/abc/123/efg/ddd
	              不能匹配：/api
	  * 输出结果：具体输出结果需要看第二个参数的Get()方法所指定的模板和数据。
	*/
	beego.Router("/api/*", &controllers.MainController{})

	/**
	  * 正则路由实践
	  * 注意事项：:无
	  * 测试用例：能 匹 配：/api/bei_123.html ; /api/bei_0.html
	              不能匹配：/api/bei_.html ; /api/bei_123 ; /api/bei_12a.html
	  * 输出结果：具体输出结果需要看第二个参数的Get()方法所指定的模板和数据。
	*/
	beego.Router("/api/bei_:id([0-9]+).html", &controllers.MainController{})

	/**
	 * 自定义方法及RESTful规则的实践
	 */

	/**
	 * 自定义方法实践
	 * 注意事项：第三个参数"*:MyMethod"之中，不能有空格（如："* : MyMethod"），否则编译报错
	 * 测试用例：客户端进行Get，Post，Put，Delete请求
	 * 输出结果：执行 MainController的MyMethod方法，不再执行RESTful的规则：即Get请求执行Get方法……
	 */
	beego.Router("/bei", &controllers.MainController{}, "*:MyMethod")
	/**
	  * Get请求不再执行Get()函数，而是执行指定的MyGet()函数，没指定的其他请求（如Post），则按照RESTful规则执行
	  * beego.Router("/bei", &controllers.MainController{}, "get:MyGet")
	  *
	  * 多个 HTTP Method 指向同一个函数
	  * beego.Router("/bei", &controllers.MainController{}, "get,post:MyGetPost")
	  *
	  * 不同的 method 对应不同的函数，通过 ；进行分割
	  * beego.Router("/bei", &controllers.MainController{}, "get:GetFunc;post:PostFunc")
	   *
	  * 如果同时存在 * 和对应的 HTTP Method，那么优先执行 HTTP Method 的方法
	  * beego.Router("/bei", &controllers.MainController{}, "*:AllFunc;post:PostFunc")
	*/

	/**
	 * 自动路由
	 */

	/**
	  * 正则路由实践
	  * 注意事项：控制器的名称为：MainController，那么【main】为路径中的第一个参数
	  * 测试用例：路径为：/main/login      将会调用 MainController 中的Login()函数
	              路径为：/main/login/123  将会调用 MainController 中的Login()函数
	              路径为：/hello/login.htm 将会调用 MainController 中的Login()函数
	              路径为：/hello/login     不能调用 MainController 中的Login()函数
	              路径为：/hellologin      不能调用 MainController 中的Login()函数
	  * 输出结果：具体输出结果需要看控制器(MainController)的Login函数所指定的模板和数据。
	*/
	beego.AutoRouter(&controllers.MainController{}) // 把需要的路由注册到自动路由中

	/**
	 * 注解路由
	 */

	/**
	 * 注解路由实践
	 * 注意事项：重要comments：// @router
	 * 测试用例：
	 * 输出结果：具体输出结果需要看控制器(MainController)的函数所指定的模板和数据。
	 */
	beego.Include(&controllers.MainController{})

	/**
	 * namespace路由
	 */

	/**
	 * namespace路由实践
	 * 注意事项：必须要把NewNamespace的对象注册到AddNamespace中去，否则无效
	 * 测试用例：仅仅匹配：/aaa/go
	 * 输出结果：get请求的情况下，执行指定函数Login；其他HTTP method则按照RESTful规则
	 */
	ns_1 := beego.NewNamespace("aaa", beego.NSRouter("/go", &controllers.MainController{}, "get:Login"))
	// beego.AddNamespace(ns_1)

	/**
	 * 域名如果不是：127.0.0.1，则不可以匹配 /bbb/go
	 * 仅仅匹配：/bbb/go
	 */
	ns_2 := beego.NewNamespace("bbb",
		beego.NSCond(func(ctx *context.Context) bool {
			if ctx.Input.Domain() == "127.0.0.1" {
				return true
			}
			return false
		}),
		beego.NSRouter("/go", &controllers.MainController{}, "get:Login"),
	)

	ns_3 := beego.NewNamespace("ccc",
		beego.NSRouter("/go", &controllers.MainController{}, "get:Login"),
		// 条件判断，如果为真，则可以匹配上下文的路由，如果为假，则上下文的路由都不能匹配
		beego.NSCond(func(ctx *context.Context) bool {
			if ctx.Input.Domain() == "127.0.0.1" {
				return true
			}
			return false
		}),
		beego.NSRouter("php", &controllers.MainController{}),
		beego.NSGet("java", func(ctx *context.Context) {
			ctx.Output.Body([]byte("显示Get追加的内容"))
		}),

		// nasespace嵌套示例
		beego.NSNamespace("room",
			beego.NSCond(func(ctx *context.Context) bool {
				// 如果子namespace的判断条件为假，那么仅仅是子namespace的url不能匹配，不影响夫namespace的匹配结果
				if ctx.Input.Request.Method != "GET" {
					return true
				}
				return false
			}),
			beego.NSRouter("/shanghai", &controllers.MainController{}), // 匹配地址：/ccc/room/shanghai
		),
	)

	beego.AddNamespace(ns_1, ns_2, ns_3)
}
