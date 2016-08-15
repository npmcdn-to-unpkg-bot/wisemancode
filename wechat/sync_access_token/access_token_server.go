package syncAccessToken

import (
	"wisemancode/log"
	"wisemancode/utils"
)

<<<<<<< HEAD
//AccessToken 微信校验串
=======
//AccessToken
>>>>>>> f880d1c0dc0a089e9ca24c71d097b41f4e351c16
type AccessToken struct {
	AccessTokenStr string //ACCESS_TOKEN 字符串
	StartTime      string //获取时间
	EndTime        string //失效时间
	LiveTime       int    //存活时间 秒
}

<<<<<<< HEAD
//AcessTokenURL 获取维系字符串URL
=======
>>>>>>> f880d1c0dc0a089e9ca24c71d097b41f4e351c16
type AcessTokenURL struct {
	url            string
	aPPID          string
	appsecret      string
	grantType      string
	accessTokenURL string
}

var (
	accessToken *AcessTokenURL
)

func init() {
<<<<<<< HEAD
	accessToken = newAccessTokenURL()
}
func newAccessTokenURL() (urlToken *AcessTokenURL) {
	log.Logger.Info("生产实例")
=======
	accessToken = newAcessTokenURL()
}
func newAcessTokenURL() (urlToken *AcessTokenURL) {
>>>>>>> f880d1c0dc0a089e9ca24c71d097b41f4e351c16
	url := utils.GetWxConfig("access_token_url")
	appid := utils.GetWxConfig("appID")
	appsecret := utils.GetWxConfig("appsecret")
	grantType := utils.GetWxConfig("grant_type")
	accessTokenURL := url + "?grant_type=" + grantType + "&&appid=" + appid + "&secret=" + appsecret
	return &AcessTokenURL{url: url, aPPID: appid, appsecret: appsecret, grantType: grantType, accessTokenURL: accessTokenURL}
}
