package model

import (
	"github.com/astaxie/beego/orm"
)

//Subscribe 订阅者信息
type Subscribe struct {
	Header       *Head  //表头信息
	ToUserName   string //开发者微信号
	FromUserName string //关注者微信号
	CreateTime   int64  //微信创建时间
	MsgType      string //消息类型
	Event        string // 消息类型
}

//NewSubscribe 创建订阅者
func NewSubscribe() *Subscribe {
	sub := new(Subscribe)
	sub.Header = NewHead()
	return sub
}

//NewSubscribeByAgrs 生成带参数的构造函数
func NewSubscribeByAgrs(touserName, fromUserName, msgType, event string, createTime int64) *Subscribe {
	sub := NewSubscribe()
	sub.CreateTime = createTime
	sub.ToUserName = touserName
	sub.FromUserName = fromUserName
	sub.MsgType = msgType
	sub.Event = event
	return sub
}

//TableName 表的名称
func (subscribe *Subscribe) TableName() string {
	return "subscribe"
}
func init() {
	orm.RegisterModel(new(Subscribe))
}

//AddSubscribe 添加事件
func (subscribe *Subscribe) AddSubscribe() (int64, error) {
	o := orm.NewOrm()
	return o.Insert(subscribe)
}
