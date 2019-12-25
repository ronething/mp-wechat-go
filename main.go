package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/ronething/mp-wechat-go/config"
	"github.com/ronething/mp-wechat-go/route"
	"net/http"
)

func main() {
	fmt.Printf("hello mp-wechat-go\n")

	var (
		err error
		g   *gin.Engine
	)

	if err = config.InitConfig(); err != nil {
		goto ERR
	}

	gin.SetMode("debug")

	g = gin.New()
	route.Register(g)
	// 启动
	fmt.Printf("start to listening on address: %s\n", config.MpConfig.App.Addr)

	if err = http.ListenAndServe(config.MpConfig.App.Addr, g); err != nil {
		goto ERR
	}

	return

ERR:
	fmt.Printf("%+v", err)
}
