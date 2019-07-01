package api

import (
	"net/http"
	"time"
	"weixin/util"
)

var log util.LogUtil

func initDispath() {
	http.HandleFunc("/dataServer/orderByDay", QueryOrderByDay)
}

func Start(ip string) {
	initDispath()
	log.Info(http.ListenAndServe(ip, nil))
}

func EndRequestLog(startTime int64, url string) {
	endTime := time.Now().Unix()
	log.Info(url, " used ", endTime-startTime, "ms")
}
