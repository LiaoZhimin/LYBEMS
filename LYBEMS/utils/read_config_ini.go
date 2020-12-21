package utils

import (
	"fmt"
	"strconv"

	"github.com/Unknwon/goconfig"
)

var cfg *goconfig.ConfigFile

// GetConfigIni 传入ini配置文件路径，解析获取配置信息
func GetConfigIni(filepath string) (err error) {
	config, err := goconfig.LoadConfigFile(filepath)
	if err != nil {
		fmt.Println("配置文件读取错误,找不到配置文件", err)
		return err
	}
	cfg = config
	return nil
}

// GetDatabase 获取ini文件中的数据库配置信息
func GetDatabase() (types, local, online string, err error) {
	if types, err = cfg.GetValue("database", "types"); err != nil {
		fmt.Println("配置文件中不存在types", err)
		return types, local, online, err
	}
	if local, err = cfg.GetValue("database", "local"); err != nil {
		fmt.Println("配置文件中不存在local", err)
		return types, local, online, err
	}
	if online, err = cfg.GetValue("database", "online"); err != nil {
		fmt.Println("配置文件中不存在online", err)
		return types, local, online, err
	}
	return types, local, online, nil
}

// GetSelf 获取ini文件中的 自定义的配置
func GetSelf() (port string, flag, tag int, err error) {
	if port, err = cfg.GetValue("self", "port"); err != nil {
		fmt.Println("配置文件中不存在port", err)
		return port, flag, tag, err
	}

	flagTemp, err := cfg.GetValue("self", "flag")
	if err != nil {
		fmt.Println("配置文件中不存在flag", err)
		return port, flag, tag, err
	}
	flag, err = strconv.Atoi(flagTemp)
	if err != nil {
		fmt.Println("配置文件中flag类型有误", err)
		return port, flag, tag, err
	}

	tagTemp, err := cfg.GetValue("self", "tag")
	if err != nil {
		fmt.Println("配置文件中不存在tag", err)
		return port, flag, tag, err
	}
	tag, err = strconv.Atoi(tagTemp)
	if err != nil {
		fmt.Println("配置文件中tag类型有误", err)
		return port, flag, tag, err
	}

	return port, flag, tag, nil
}
