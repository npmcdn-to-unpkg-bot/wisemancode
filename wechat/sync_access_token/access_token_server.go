package sync_access_token

import (
	"wisemancode/utils"
)

type AccessToken struct {
	Access_token string //ACCESS_TOKEN 字符串
	Start_time   string //获取时间
	End_time     string //失效时间
	Live_time    int    //存活时间 秒
}

type AcessTokenUrl struct {
	url              string
	aPPID            string
	appsecret        string
	grant_type       string
	access_token_url string
}

var (
	accessToken *AcessTokenUrl
)

func init() {
	accessToken = newAccessTokenUrl()
}
func newAccessTokenUrl() (urlToken *AcessTokenUrl) {
	url := utils.GetWxConfig("access_token_url")
	appid := utils.GetWxConfig("appID")
	appsecret := utils.GetWxConfig("appsecret")
	grant_type := utils.GetWxConfig("grant_type")
	access_token_url := url + "?grant_type=" + grant_type + "&&appid=" + appid + "&secret=" + appsecret
	return &AcessTokenUrl{url: url, aPPID: appid, appsecret: appsecret, grant_type: grant_type, access_token_url: access_token_url}
}
