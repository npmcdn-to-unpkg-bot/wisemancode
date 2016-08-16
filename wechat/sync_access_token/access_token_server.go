package syncAccessToken

import (
	"encoding/json"
	"strings"
	"sync/atomic"
	"time"
	"unsafe"
	"wisemancode/log"
	"wisemancode/utils"

	"github.com/astaxie/beego/httplib"
)

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
	tokenURL *AcessTokenURL
)

func init() {
	tokenURL = newAccessTokenURL()
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

//accessResult解析获取的结果
type accessResult struct {
	accessTokenStrNew string
	err               string
}

//AccessTokenServer 定时获取微信校验服务
type AccessTokenServer struct {
	AccessTokenStr       chan string       //ACCESS_TOKEN 字符串 解析后的信息
	accessTokenResultStr chan accessResult //微信返回信息后解析结果
	LiveTime             int               //存活时间 秒
	Times                time.Duration     //时间间隔
	json                 unsafe.Pointer    //微信返回的数据结构指针
}

//AccessTokenJson 从微信获取accessToken数据
type AccessTokenJson struct {
	Token     string `json:"access_token"`
	ExpiresIn int64  `json:"expires_in"`
}

func NewDefaultAccessTokenServer() (acc *AccessTokenServer) {
	acc = &AccessTokenServer{
		AccessTokenStr:       make(chan string),
		accessTokenResultStr: make(chan accessResult),
		Times:                times,
	}
	go acc.GetAndAccessTokenServer()
	return
}

//获取并且更新微信accessToken
func (this *AccessTokenServer) GetAndAccessTokenServer() {
	//创建一个定时器
	nicker := time.NewTicker(this.Times)
	for {
		select {
		case current := <-this.AccessTokenStr:
			//数据进行更新
			this.updateAccessToken(current)
		}
	}
}
func (this *AccessTokenServer) updateAccessToken(current string) (jsonToken *AccessTokenJson, err error) {
	if len(current) != 0 {
		if p := (*AccessTokenJson)(atomic.LoadPointer(&this.json)); p != nil && current != p.Token {
			return p, nil
		}
		msg, err := httplib.Get(tokenURL.accessTokenURL).String()
		if err != nil {
			atomic.StorePointer(&this.json, nil)
			return nil, err
		}
		var j AccessTokenJson
		js := json.NewDecoder(strings.NewReader(msg))
		if e = js.Decode(&j); e != nil {
			atomic.StorePointer(&this.json, nil)
			return nil, err
		}

	}
}
