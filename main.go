package main

import (
	"fmt"
	"github.com/ronething/mp-wechat-go/config"
)

func main() {
	fmt.Printf("hello mp-wechat-go\n")

	var err error

	if err = config.InitConfig(); err != nil {
		goto ERR
	}

	fmt.Println(config.MpConfig.App.Addr)

	return

ERR:
	fmt.Printf("err: %+v", err)
}
