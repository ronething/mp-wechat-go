// author: ashing
// time: 2019/12/25 12:07 下午
// mail: axingfly@gmail.com
// Less is more.

package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/ronething/mp-wechat-go/controller/handler"
	"github.com/ronething/mp-wechat-go/serivce"
)

func Hello(c *gin.Context) {
	server := serivce.InitServer(c)
	// 设置接收消息的处理方法
	server.SetMessageHandler(handler.HelloHandler)

	// 处理消息接收以及回复
	if err := server.Serve(); err != nil {
		fmt.Printf("server err: %+v\n", err)
		return
	}
	// 发送回复的消息
	server.Send()
}
