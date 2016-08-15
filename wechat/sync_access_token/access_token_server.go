package syncAccessToken

import (
	"sync/atomic"
	"time"
	"unsafe"
	"wisemancode/log"
	"wisemancode/utils"

	"github.com/astaxie/beego/httplib"
)

//AccessTokenServer 定时获取微信校验服务
type AccessTokenServer struct {
	AccessTokenStr       chan string       //ACCESS_TOKEN 字符串 解析后的信息
	accessTokenResultStr chan accessResult //微信返回信息后解析结果
	StartTime            string            //获取时间
	EndTime              string            //失效时间
	LiveTime             int               //存活时间 秒
	Times                time.Duration     //时间间隔

	tokenCache unsafe.Pointer //微信返回的信息指针

}
type accessResult struct {
	accessTokenStrNew string
	err               string
}

//AcessTokenURL 获取维系字符串URL
type AcessTokenURL struct {
	url            string
	aPPID          string
	appsecret      string
	grantType      string
	accessTokenURL string
}

const (
	times time.Duration = 2 * 60 * 60
)

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
func NewAccessTokenServer() (accessTokenServer *AccessTokenServer) {
	accessTokenStr := make(chan string)
	access := make(chan accessResult)
	accessTokenServer = &AccessTokenServer{AccessTokenStr: accessTokenStr, accessTokenResultStr: access, Times: times}
	go accessTokenServer.GetAccessTokenByTime(accessToken)
	return
}

//GetAccessTokenByTime 启动定时器进行获取微信accessToken
func (this *AccessTokenServer) GetAccessTokenByTime(url *AcessTokenURL) {
	//启动定时器到微信服务器中进行获取accessToken
	//	currentAccessToken := this.AccessTokenStr
	//if()
	//accessToken, error := httplib.Get(url.accessTokenURL).String()
	//	if error == nil {
	//		//错误处理
	//		return
	//	}

}

type accessTokenJson struct {
	Token     string `json:"access_token"`
	ExpiresIn int64  `json:"expires_in"`
}

//获取新的token 放到缓存中 别且返回 新的token
func (this *AccessTokenServer) GetNewAccessTokenFromWeiXin(url *AcessTokenURL, currentAccess string) (token *accessTokenJson, err error) {
	if p := (*accessTokenJson)(atomic.LoadPointer(&this.tokenCache)); p != nil && currentAccess != p.Token {
		return p, nil
	}
	data, err := httplib.Get(url.accessTokenURL).String()
	log.Logger.Info("获取微信信息：data=：" + data)
	log.Logger.Info("获取微信信息：err=：", err.Error())
	if err != nil {
		//处理错误信息
		atomic.StorePointer(&this.tokenCache, nil)
		return
	}
	return
}
