package message

import (
	"encoding/xml"
	"errors"
	"wisemancode/log"
	"wisemancode/utils"
)

//MsgType 消息类型
type MsgType string

//消息
const (
	Text        MsgType = "text"       //文本消息
	Image       MsgType = "image"      //图文消息
	Voice       MsgType = "voice"      //语音消息
	Video       MsgType = "video"      //视频消息
	Shortvideo  MsgType = "shortvideo" //小视频消息
	LocationMsg MsgType = "location"   //地理位置消息
	Link        MsgType = "link"       //链接消息
	Event       MsgType = "event"      //事件消息
	NULL        MsgType = "NULL"       //消息为空
)

// EventType 事件类型
type EventType string

//事件
const (
	Subscribe     EventType = "subscribe"   //订阅事件
	Unsubscribe   EventType = "unsubscribe" //订阅事件
	Scan          EventType = "scan"        //扫描事件
	LocationEvent EventType = "location"    //上报地理位置事件
	Click         EventType = "click"       //自定义菜单事件
)

//Message 消息体 定义微信返回的信息
type Message struct {
	XMLName           xml.Name  `xml:"xml"`
	ToUserName        string    `xml:"ToUserName"`   //开发者微信号
	FromUserName      string    `xml:"FromUserName"` //发送方帐号（一个OpenID）
	CreateTime        int64     `xml:"CreateTime"`   //消息创建时间 （整型）
	MsgType           MsgType   `xml:"MsgType"`      //消息类型
	MsgID             int64     `xml:"MsgId"`        //消息id，64位整型
	MediaID           string    `xml:"MediaId"`      // 语音、图片、 视频 等消息ID
	Content           string    `xml:"Content"`      //文本消息
	PicURL            string    `xml:"PicUrl"`       //图片链接（由系统生成）
	VoiceFormat       string    `xml:"Format"`       //语音格式
	VoiceRecognition  string    `xml:"Recognition"`  //语音识别结果，UTF8编码
	VideoThumbMediaID string    `xml:"ThumbMediaId"` //视频消息缩略图的媒体id，可以调用多媒体文件下载接口拉取数据。
	LocationLabel     string    `xml:"Label"`        //地理位置信息
	LocationX         float64   `xml:"Location_X"`   //地理位置维度
	LocationY         float64   `xml:"Location_Y"`   //地理位置维度
	LocationScale     float64   `xml:"Scale"`        //地图缩放大小
	LinkTitle         string    `xml:"Title"`        //链接消息标题
	LinkDescription   string    `xml:"Description"`  //链接消息描述
	LinkURL           string    `xml:"Url"`          //链接消息URL
	Event             EventType `xml:"Event"`        //事件类型
	EventKey          string    `xml:"EventKey"`     //事件KEY值
	EventTicket       string    `xml:"Ticket"`       //二维码的ticket，可用来换取二维码图片
	EventLatitude     float64   `xml:"Latitude"`     //地理位置纬度
	EventLongitude    float64   `xml:"Longitude"`    //地理位置纬度
	EventPrecision    float64   `xml:"Precision"`    //地理位置精度
}

// MsgRSender 接收消息,发送消息 需要完善
type MsgRSender interface {
	Receive(xmlContent string) (msgType MsgType, err error)
	Send()
}

var adapters = make(map[MsgType]MsgRSender)

//Register 注册适配 每一个消息类型都是一个适配
func Register(name MsgType, adapter MsgRSender) {
	if adapter == nil {
		panic("MsgRSender: adapter is nil")
	}
	if _, ok := adapters[name]; ok {
		log.Logger.Error("MsgRSender: adapter Register twice")
		return
	}
	adapters[name] = adapter
}

//NewMessage 实例化
func NewMessage() (msg *Message) {
	return &Message{}
}

//Receive 消息接收，并返回消息类型
func (msg *Message) Receive(xmlContent string) (msgType MsgType, err error) {
	log.Logger.Info("开始接收消息,并且解析：%s", xmlContent)
	if len(xmlContent) == 0 {
		log.Logger.Error("接收的消息len==0")
		return NULL, errors.New("接收的消息len==0")
	}

	err = utils.ParseXML(xmlContent, msg)
	if err != nil {
		return NULL, err
	}
	msgType = msg.MsgType
	log.Logger.Info("开始接收消息，解析结束：%+v", msg)
	return msgType, error(nil)
}

//Send 向微信服务器发送消息
func (msg *Message) Send() {

}

//WXMessage 结束消息门面
func WXMessage(xmlContent string) {
	var msg = NewMessage()
	msgType, err := msg.Receive(xmlContent)
	if err != nil {

	}
	adapter := adapters[msgType]
	adapter.Send()
}
