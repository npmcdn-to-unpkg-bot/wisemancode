package syncAccessToken

import (
	"wisemancode/log"
	"wisemancode/utils"
)

//AccessToken 微信校验串
type AccessToken struct {
	AccessTokenStr string //ACCESS_TOKEN 字符串
	StartTime      string //获取时间
	EndTime        string //失效时间
	LiveTime       int    //存活时间 秒
}

//AcessTokenURL 获取维系字符串URL
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
	accessToken = newAccessTokenURL()
}
func newAccessTokenURL() (urlToken *AcessTokenURL) {
	log.Logger.Info("生产实例")
	urlToken = newAcessTokenURL()
	return
}
func newAcessTokenURL() (urlToken *AcessTokenURL) {
	url := utils.GetWxConfig("access_token_url")
	appid := utils.GetWxConfig("appID")
	appsecret := utils.GetWxConfig("appsecret")
	grantType := utils.GetWxConfig("grant_type")
	accessTokenURL := url + "?grant_type=" + grantType + "&&appid=" + appid + "&secret=" + appsecret
	return &AcessTokenURL{url: url, aPPID: appid, appsecret: appsecret, grantType: grantType, accessTokenURL: accessTokenURL}
}
