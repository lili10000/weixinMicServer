package util

import (
	// "bufio"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"weixin/entity"
)

var config entity.ConfigJson

var configPath = "config/application.json"

// var configPath = "../config/application.json"

type Config struct {
	initFlag bool
}

func (p *Config) init() {
	data, err := ioutil.ReadFile(configPath)
	if err != nil {
		fmt.Println(filePath, err)
		os.Exit(1)
	}
	err = json.Unmarshal(data, &config)
	if err != nil {
		fmt.Println(filePath, err)
		os.Exit(1)
	}

}

func (p *Config) Get() entity.ConfigJson {
	if p.initFlag == false {
		p.init()
	}
	return config
}
