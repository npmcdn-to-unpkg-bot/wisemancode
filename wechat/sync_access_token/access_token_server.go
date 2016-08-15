package sync_access_token

import (
	"wisemancode/utils"
)

//AccessToken
type AccessToken struct {
	Access_token string //ACCESS_TOKEN 字符串
	Start_time   string //获取时间
	End_time     string //失效时间
	Live_time    int    //存活时间 秒
}

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
	accessToken = newAcessTokenURL()
}
func newAcessTokenURL() (urlToken *AcessTokenURL) {
	url := utils.GetWxConfig("access_token_url")
	appid := utils.GetWxConfig("appID")
	appsecret := utils.GetWxConfig("appsecret")
	grantType := utils.GetWxConfig("grant_type")
	accessTokenURL := url + "?grant_type=" + grantType + "&&appid=" + appid + "&secret=" + appsecret
	return &AcessTokenURL{url: url, aPPID: appid, appsecret: appsecret, grantType: grantType, accessTokenURL: accessTokenURL}
}
