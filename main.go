package main

import (
	"weixin/util"
	// "fmt"
	"weixin/api"
)

func main() {
	var log util.LogUtil
	log.Init("test.log")
	log.Info("start work!")
	api.Start(":10000")
}
