// author: ashing
// time: 2019/12/25 11:26 上午
// mail: axingfly@gmail.com
// Less is more.

package config

import (
	"fmt"
	"github.com/spf13/viper"
)

type wechat struct {
	AppID          string
	AppSecret      string
	Token          string
	EncodingAESKey string
}

type app struct {
	Addr string
}

type config struct {
	App app
	Wechat  wechat
}

var MpConfig config

func InitConfig() (err error) {

	v := viper.New()
	v.SetConfigFile("wechat.yaml")

	if err = v.ReadInConfig(); err != nil {
		fmt.Printf("read config err: %+v", err)
		return err
	}

	if err = v.Unmarshal(&MpConfig); err != nil {
		fmt.Printf("init config err: %+v", err)
		return err
	}

	return
}
