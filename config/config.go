package config

import (
	"encoding/json"
	"os"
)

//项目配置
type AppConfig struct {
	AppName string `json:"app_name"`
	Port string `json:"port"`
	StaticPath string `json:"static_path"`
	Mode string `json:"mode"`
}

var ServConfig AppConfig

func InitConfi() *AppConfig{
	file, err := os.Open("C:/Users/tuotuo/Desktop/js/GOPROJECT/myapp/config.json")
	if err != nil{
		panic(err.Error())
	}
	decoder := json.NewDecoder(file)
	conf := AppConfig{}
	err = decoder.Decode(&conf)
	if err != nil{
		panic(err.Error())
	}
	return  &conf
}