package syncAccessToken

import (
	"encoding/json"
	"io"
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
	times time.Duration = 2 * time.Hour
)

var (
	tokenURL *AcessTokenURL
)

//Tokener 获取微信服务器的accessToken
type Tokener interface {
	Token() (token string, err error)                      //获取微信accessToken
	RefreshToken(current string) (token string, err error) //刷新微信accessToken
}

func init() {
	log.Logger.Info("获取accessToken的配置")
	tokenURL = newAccessTokenURL()
	log.Logger.Info("获取accessToken的配置 %+v", tokenURL)
}
func newAccessTokenURL() (urlToken *AcessTokenURL) {
	log.Logger.Info("获取accessToken的配置==============start===================")
	urlToken = newAcessTokenURL()
	log.Logger.Info("accessToken 配置文件信息：%+v", urlToken)
	log.Logger.Info("获取accessToken的配置==============end===================")
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
	err               error
}

//AccessTokenServer 定时获取微信校验服务
type AccessTokenServer struct {
	CurrentAccessToken chan string       //ACCESS_TOKEN 字符串 解析后的信息
	accessTokenResult  chan accessResult //微信返回信息后解析结果
	LiveTime           int               //存活时间 秒
	Times              time.Duration     //时间间隔
	json               unsafe.Pointer    //微信返回的数据结构指针
}

//AccessTokenJson 从微信获取accessToken数据
type AccessTokenJson struct {
	Token     string `json:"access_token"`
	ExpiresIn int64  `json:"expires_in"`
}

func NewDefaultAccessTokenServer() (acc *AccessTokenServer) {
	acc = &AccessTokenServer{
		CurrentAccessToken: make(chan string),
		accessTokenResult:  make(chan accessResult),
		Times:              times,
	}
	log.Logger.Info("accessToken 获取，数据结构初始化：%+v", acc)
	log.Logger.Info("accessToken 启动定时器，进行请求accessToken")
	go acc.GetAndAccessTokenServer()
	return acc
}

//获取并且更新微信accessToken
func (this *AccessTokenServer) GetAndAccessTokenServer() {
	//创建一个定时器
	log.Logger.Info("GetAndAccessTokenServer创建定时器，进行请求accessToken")
TIMES_TICKER:
	nicker := time.NewTicker(this.Times)
	for {
		log.Logger.Info("GetAndAccessTokenServer,for 循环")
		select {
		case current := <-this.CurrentAccessToken:
			//数据进行更新
			log.Logger.Info("当前CurrentAccessToken:" + current)
			acc, err := this.updateAccessToken(current)
			log.Logger.Info("已经请求微信服务器获取accessToken:%+v", acc)
			if err != nil {
				log.Logger.Info("获取微信服务器错误")
				this.accessTokenResult <- accessResult{err: err}
				log.Logger.Info("获取微信服务器错误accessTokenResult %+v", this)
				break
			}
			this.accessTokenResult <- accessResult{accessTokenStrNew: acc.Token}
			log.Logger.Info("获取微信accessToken %+v", this)
		case <-nicker.C:
			log.Logger.Info("获取微信accessToken 启动定时器时间nicker.C %+v", nicker.C)
			acc, err := this.updateAccessToken("")
			log.Logger.Info("获取微信accessToken 定时器获取数据 错误nicker.C %+v", err)
			log.Logger.Info("获取微信accessToken 定时器获取数据nicker.C %+v", acc)
			if err != nil {
				this.accessTokenResult <- accessResult{err: err}
				log.Logger.Info("获取微信accessToken 定时器获取数据错误nicker.C %+v", this)
				break
			}
			t := time.Duration(acc.ExpiresIn) * time.Second
			log.Logger.Info("获取数据时间 ExpiresIn  %d", acc.ExpiresIn)
			log.Logger.Info("获取数据时，标准时间差  %d", this.Times-t)
			if this.Times-t > 5*time.Second {
				log.Logger.Info("获取数据时间太长%d", acc.ExpiresIn)
				log.Logger.Info("获取数据时间太长,重新启动定时器")
				this.Times = t
				nicker.Stop()
				goto TIMES_TICKER
			}
			log.Logger.Info("获取微信accessToken 定时器获取数据AccessTokenServer %+v", this)
			//this.accessTokenResult <- accessResult{accessTokenStrNew: acc.Token}被堵塞了
			log.Logger.Info("获取微信accessToken 定时器获取数据AccessTokenServer %+v", this)

		}
	}
}

//更新缓存
func (this *AccessTokenServer) updateAccessToken(current string) (jsonToken *AccessTokenJson, err error) {
	log.Logger.Info("链接微信服务器获取accessToken")
	if len(current) != 0 {
		log.Logger.Info("当前缓存数据：" + current)
		if p := (*AccessTokenJson)(atomic.LoadPointer(&this.json)); p != nil && current != p.Token {
			log.Logger.Info("当前缓存数据：%+v", p)
			return p, nil
		}
	}
	log.Logger.Info("与微信服务器建立链接")
	msg, err := httplib.Get(tokenURL.accessTokenURL).String()
	log.Logger.Info("获取微信服务器数据：" + msg)
	if err != nil {
		log.Logger.Info("获取微信服务器数据错误：%+v", err)
		atomic.StorePointer(&this.json, nil)
		return nil, err
	}
	var j AccessTokenJson
	r := strings.NewReader(msg)
	log.Logger.Info("获取微信服务器数据,r   %+v", r)
	dec := json.NewDecoder(r)
	log.Logger.Info("获取微信服务器数据,dec   %+v", dec)
	//log.Logger.Info("获取微信服务器数据,dec   %+v", dec.Decode(&j))
	for {
		if e := dec.Decode(&j); e == io.EOF {
			log.Logger.Info("获取微信服务器数据,解析数据：%+v", j)
			atomic.StorePointer(&this.json, unsafe.Pointer(&j))
			return &j, nil
		} else if e != nil {
			log.Logger.Info("获取微信服务器数据,解析数据错误：%+v", e)
			atomic.StorePointer(&this.json, nil)
			return nil, e
		}
	}

	return
}

//从当前缓存中获取最新的AccessToken 是一个原子操作
func (this *AccessTokenServer) Token() (token string, err error) {
	log.Logger.Info("获取accessToken Token()")
	if p := (*AccessTokenJson)(atomic.LoadPointer(&this.json)); p != nil {
		log.Logger.Info("获取accessToken Token()  %+v", p)
		token, err = p.Token, error(nil)
	}
	token, err = this.RefreshToken("")
	log.Logger.Info("获取accessToken Token()  %+v", token)
	return
}

//刷新就是把当前的accesstoken 换成最新的accessToken 很大可能是一致
func (this *AccessTokenServer) RefreshToken(current string) (token string, err error) {
	log.Logger.Info("获取accessToken RefreshToken()  %+s", current)
	this.CurrentAccessToken <- current
	res := <-this.accessTokenResult
	log.Logger.Info("获取accessToken RefreshToken()  %+v", res)
	return res.accessTokenStrNew, res.err
}
