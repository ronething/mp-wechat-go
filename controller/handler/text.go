// author: ashing
// time: 2019/12/25 12:45 下午
// mail: axingfly@gmail.com
// Less is more.

package handler

import "github.com/silenceper/wechat/message"

func HelloHandler(msg message.MixMessage) *message.Reply {



	//回复消息：演示回复用户发送的消息
	text := message.NewText(msg.Content)
	return &message.Reply{MsgType: message.MsgTypeText, MsgData: text}
}
