package config

import (
	"encoding/json"
	"io/ioutil"
)

// 配置信息
type Config struct {
	ListenAddr string `json:"listen_addr"`
	AuthToken  string `json:"auth_token"`
}

// 全局配置
var AppConfig Config

// 加载配置
func Load(configFile string) (err error) {
	data, err := ioutil.ReadFile(configFile)
	if err != nil {
		return
	}
	return json.Unmarshal(data, &AppConfig)
}
