package config

import (
	"encoding/json"
	"io/ioutil"
	"vue-next-admin-go/modles"
)

var Config = modles.Configuration{}

func InitConfiguration() {
	configFile := "./config.json"
	// ReadFile函数会读取文件的全部内容，并将结果以[]byte类型返回
	data, err := ioutil.ReadFile(configFile)

	if err != nil {
		panic(err)
	}
	// 读取的数据为json格式，需要进行解码
	err = json.Unmarshal(data, &Config)
	if err != nil {
		panic(err)
	}
	return
}
