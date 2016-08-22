package message

//TextMsg 文本消息处理
type TextMsg struct{}

//Receive 接收消息 默认调用Message 实现
func (text *TextMsg) Receive(xmlContent string) (msgType MsgType, err error) {
	msg := NewMessage()
	return msg.Receive(xmlContent)
}

//Send 发送文本消息
func (text *TextMsg) Send() {

}

func init() {
	Register(Text, &TextMsg{})
}
