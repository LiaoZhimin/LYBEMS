package utils

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"sync"
)

//GlobalConfig 全局配置 包括 数据库配置 和 自定义
type GlobalConfig struct {
	Database DatabaseConfig `json:"database"`
	Self     SelfConfig     `json:"self"`
}

// DatabaseConfig 数据库配置
type DatabaseConfig struct {
	Types  string `json:"types"`
	Local  string `json:"local"`
	Online string `json:"online"`
}

// SelfConfig 自定义
type SelfConfig struct {
	Port string `json:"port"`
	Flag int    `json:"flag"`
	Tag  int    `json:"tag"`
}

var (
	globalconfig *GlobalConfig
	configMux    sync.RWMutex
)

// Config 函数 用于返回 全局配置对象
func Config() *GlobalConfig {
	return globalconfig
}

// InitConfigJSON 传入配置文件路径，解析获取
func InitConfigJSON(fliepath string) error {
	var config GlobalConfig
	file, err := ioutil.ReadFile(fliepath)
	if err != nil {
		fmt.Println("配置文件读取错误,找不到配置文件", err)
		return err
	}

	if err = json.Unmarshal(file, &config); err != nil {
		fmt.Println("配置文件读取失败", err)
		return err
	}

	configMux.Lock()
	globalconfig = &config
	configMux.Unlock()
	return nil
}
