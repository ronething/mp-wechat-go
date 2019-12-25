// author: ashing
// time: 2019/12/25 12:14 下午
// mail: axingfly@gmail.com
// Less is more.

package serivce

import (
	"github.com/gin-gonic/gin"
	"github.com/ronething/mp-wechat-go/config"
	"github.com/silenceper/wechat"
	"github.com/silenceper/wechat/server"
)

// init wechat server
func InitServer(c *gin.Context) *server.Server {

	var (
		wechatCli    *wechat.Wechat
		wechatConfig *wechat.Config
	)
	wechatConfig = &wechat.Config{
		AppID:          config.MpConfig.Wechat.AppID,
		AppSecret:      config.MpConfig.Wechat.AppSecret,
		Token:          config.MpConfig.Wechat.Token,
		EncodingAESKey: config.MpConfig.Wechat.EncodingAESKey,
	}
	wechatCli = wechat.NewWechat(wechatConfig)
	// 传入request和responseWriter
	return wechatCli.GetServer(c.Request, c.Writer)

}
