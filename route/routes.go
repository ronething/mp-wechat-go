// author: ashing
// time: 2019/12/25 12:08 下午
// mail: axingfly@gmail.com
// Less is more.

package route

import (
	"github.com/gin-gonic/gin"
	"github.com/ronething/mp-wechat-go/controller"
)

func Register(g *gin.Engine)  {
	g.Use(gin.Recovery())
	g.Use(gin.Logger())

	registerApi(g)
}

func registerApi(g *gin.Engine) {
	g.GET("/", controller.Hello)
	g.POST("/", controller.Hello)
}
