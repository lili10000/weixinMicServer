package api

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
	"weixin/dao"
	. "weixin/entity"
	"weixin/entity/errCode"
)

func QueryOrderByDay(w http.ResponseWriter, req *http.Request) {
	startTime := time.Now().UnixNano() / 1e6
	defer EndRequestLog(startTime, req.URL.Path)

	values := req.URL.Query()

	dataStr := values.Get("date")
	countRetn := dao.QuerySellOrderDay(dataStr)
	sellCountRetn := dao.QuerySellNumDay(dataStr)
	moneyPrice := dao.QuerySellMoneyPriceDay(dataStr)
	moneyRecv := dao.QuerySellMoneyRecvDay(dataStr)

	var data OrderCount
	data.Sum = fmt.Sprintf("订单: %s 销售: %s 销售总额: %s 实收: %s", countRetn, sellCountRetn, moneyPrice, moneyRecv)
	sessionList := dao.QuerySession(dataStr)
	for _, session := range sessionList {
		addr := session[0]
		time := session[1]
		sessionCount := dao.QuerySellOrderSession(time)
		sessionSellCount := dao.QuerySellNumSession(time)
		sessionMoneyPrice := dao.QuerySellMoneyPriceSession(time)
		sessionMoneyRecv := dao.QuerySellMoneyRecvSession(time)
		sessionInfo := fmt.Sprintf("场次：%s 时间: %s 订单: %s 销售: %s 销售总额: %s 实收: %s", addr, time, sessionCount, sessionSellCount, sessionMoneyPrice, sessionMoneyRecv)
		data.Detail = append(data.Detail, sessionInfo)
	}

	var retn OrderCountRetn
	retn.Code = errCode.Success
	retn.Msg = ""
	retn.Data = data
	retnStr, err := json.Marshal(retn)
	if err != nil {
		w.Write([]byte("系统错误"))
	}
	w.Write([]byte(retnStr))
}
