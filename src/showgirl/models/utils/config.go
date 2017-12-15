package utils

import (
	"github.com/astaxie/beego"
	"strings"
)

func GetConfigByString(key string) string {
	return beego.AppConfig.String(key)
}

func GetConfigByInt(key string) int {
	value, err := beego.AppConfig.Int(key)
	if err != nil {
		return -1
	}
	return value
}

func GetConfigByInt64(key string) int64 {
	value, err := beego.AppConfig.Int64(key)
	if err != nil {
		return -1
	}
	return value
}

func SetConfig(key string, val string) error {
	err := beego.AppConfig.Set(key, val)
	return err
}

func GetSection(key string) (section map[string]string) {
	key = strings.ToLower(key)
	section, _ = beego.AppConfig.GetSection(key)
	return
}

func GetConfigByBool(key string) bool {
	value, err := beego.AppConfig.Bool(key)
	if err != nil {
		return false
	}
	return value
}
