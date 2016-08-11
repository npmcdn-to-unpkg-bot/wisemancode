package log

import (
	"github.com/astaxie/beego/logs"
)

var (
	Logger *logs.BeeLogger
)

func init() {
	Logger = logs.NewLogger(10000)
	Logger.SetLogger("console", "")
	Logger.SetLogger("file", `{"filename":"test.log"}`)

	Logger.SetLevel(logs.LevelDebug)
	Logger.EnableFuncCallDepth(true)

}
