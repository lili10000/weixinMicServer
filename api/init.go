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
	log.Info(http.ListenAndServeTLS(ip, "1_gacicv.com_bundle.crt", "2_gacicv.com.key", nil))
}

func EndRequestLog(startTime int64, url string) {
	endTime := time.Now().UnixNano() / 1e6
	log.Info(url, " used ", endTime-startTime, "s")
}
