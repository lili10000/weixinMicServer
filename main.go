package main

import (
	"weixin/util"
	"fmt"
	// "net/http"
)

func main() {
	// tool := new(util.LogUtil);
	var tool util.LogUtil
	tool.FilePath = "test.log"
	// tool.init("test.log")
	log,err := tool.Get()
	if err != nil{
		fmt.Println(err)
	}

	log.Info("Hello, World!")
	// fmt.Println("Hello, World!")
}
