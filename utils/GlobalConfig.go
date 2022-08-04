package utils

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

// GlobalConfig 全局配置参数，供其他模块使用
type GlobalConfig struct {
	Host           string
	Port           int
	Name           string
	Version        string
	MaxConn        int
	MaxPackageSize uint32
}

// ServerConfig 全局配置
var ServerConfig *GlobalConfig

func (config *GlobalConfig) LoadConfig() {
	data, readErr := ioutil.ReadFile("conf/kronus.json")
	if readErr != nil {
		fmt.Println("[LoadConfig] ReadFile error:", readErr)
		return
	}
	parseErr := json.Unmarshal(data, &config)
	if parseErr != nil {
		fmt.Println("[LoadConfig] Unmarshal error:", parseErr)
	}
}

func init() {
	// 默认配置
	ServerConfig = &GlobalConfig{
		Name:           "Kronus Server",
		Version:        "v0.3",
		Host:           "0.0.0.0",
		Port:           8999,
		MaxConn:        1000,
		MaxPackageSize: 4096,
	}
	ServerConfig.LoadConfig()
}
