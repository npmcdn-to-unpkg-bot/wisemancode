package utils

import (
	"encoding/xml"
	"wisemancode/log"

	"github.com/astaxie/beego"
)

//获取配置文件和相关工具方法的集合
//type Wxconfig struct {
//	WxConfigs map[string]string
//}

//var (
//	//微信配置文件
//	WXcon *Wxconfig
//)

//func init() {

//	//初始化微信配置文件
//	WXcon = NewWxconfig()

//}

////创建微信配置文件实例
//func NewWxconfig() *Wxconfig {
//	config := make(map[string]string)
//	con := &Wxconfig{WxConfigs: config}
//	return con
//}
const configName string = "wxconfig"

func GetWxConfig(key string) (value string) {
	if len(key) == 0 {
		log.Logger.Error("key is nil")
		return ""
	}

	return userConfig(configName, key)
}
func userConfig(conName, key string) (value string) {
	if len(key) == 0 || len(conName) == 0 {
		log.Logger.Error("key is nil")
		return ""
	}
	value = beego.AppConfig.String(conName + "::" + key)
	log.Logger.Info("key is " + key + "||value is " + value)
	return
}

//parseXML 解析xml
func ParseXML(content string, v interface{}) (err error) {
	log.Logger.Info("需要解析的XML数据：%s", content)
	//	buf := make([]byte, len(content))
	//	buf = append(buf, content...)
	err = xml.Unmarshal([]byte(content), v)
	if err != nil {
		log.Logger.Info("解析xml发生错误：%+v", err)
	}
	log.Logger.Info("解析xml结果：%+v", v)
	return
}
