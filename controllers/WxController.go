package controllers

import (
	"fmt"

	//"io/ioutil"

	"wisemancode/log"
	"wisemancode/utils"
	"wisemancode/wechat/message"
	"wisemancode/wechat/sync_access_token"

	"github.com/astaxie/beego"
)

//创建微信认证服务
type WXController struct {
	beego.Controller
}

func (wxCon *WXController) Post() {
	wxData := fmt.Sprintf("%s", wxCon.Ctx.Input.RequestBody)
	log.Logger.Info("RequestBody:%s", wxData)
	//s := fmt.Sprintf("%s", wxCon.Ctx.Input.CopyBody())
	//log.Logger.Info("CopyBody:%s", s)
	message.WXMessage(wxData)
}

//http://localhost:8088/wx?timestamp=timestamp&signature=signature&nonce=nonce&echostr=echostr
func (wxCon *WXController) Get() {
	log.Logger.Info("微信服务器进行验证")

	log.Logger.Info("接受微信服务器参数：" + wxCon.Ctx.Request.URL.String())
	signature := wxCon.GetString("signature")
	timestamp := wxCon.GetString("timestamp")
	nonce := wxCon.GetString("nonce")
	echostr := wxCon.GetString("echostr")
	wxCon.TplName = "index.tpl"
	//获取配置数据
	token := utils.GetWxConfig("token")
	if len(token) == 0 {
		log.Logger.Error("Config is fund token")
		return
	}
	log.Logger.Info("微信服务器发来的签名串：" + signature)
	signa, e := utils.Sign(token, timestamp, nonce)
	if e != nil {
		log.Logger.Error(e.Error())
		return
	}
	if signa == signature {
		log.Logger.Info("微信服务器验证成功")
		wxCon.Ctx.WriteString(echostr)
		return
	} else {
		log.Logger.Error("微信服务起验证签名失败")
		return
	}
}

func (wxCon *WXController) Test() {
	log.Logger.Info("接受微信服务器参数：" + wxCon.Ctx.Request.URL.String())
	log.Logger.Info("test 数据")
	wxCon.TplName = "index.tpl"
	var acc *syncAccessToken.AccessTokenServer = syncAccessToken.NewDefaultAccessTokenServer()
	log.Logger.Info("test 数据===================== %+v", acc)
	acc.Token()
	return
}
