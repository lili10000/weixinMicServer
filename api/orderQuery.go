package api

import (
	"encoding/json"
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

	recvList := dao.QueryDayInfo(dataStr)
	// countRetn := dao.QuerySellOrderDay(dataStr)
	// sellCountRetn := dao.QuerySellNumDay(dataStr)
	// moneyPrice := dao.QuerySellMoneyPriceDay(dataStr)
	// moneyRecv := dao.QuerySellMoneyRecvDay(dataStr)

	var data OrderCount
	data.Sum = SessionInfo{
		OrderCount: recvList[0],
		SellCount:  recvList[1],
		MoneyPrice: recvList[2],
		MoneyRecv:  recvList[3],
	}
	// fmt.Println(data.Sum)
	// fmt.Sprintf("订单: %s 销售: %s 销售总额: %s 实收: %s", countRetn, sellCountRetn, moneyPrice, moneyRecv)
	sessionList := dao.QuerySession(dataStr)
	for _, session := range sessionList {
		var info SessionInfo
		time := session[1]
		info.Addr = session[0]
		info.Time = session[1]

		recvList = dao.QuerySessionInfo(time)
		info.OrderCount = recvList[0]
		info.SellCount = recvList[1]
		info.MoneyPrice = recvList[2]
		info.MoneyRecv = recvList[3]

		// info.OrderCount = dao.QuerySellOrderSession(time)
		// info.SellCount = dao.QuerySellNumSession(time)
		// info.MoneyPrice = dao.QuerySellMoneyPriceSession(time)
		// info.MoneyRecv = dao.QuerySellMoneyRecvSession(time)
		// sessionInfo := fmt.Sprintf("时间: %s	场次：%s\n订单: %s 销售: %s 销售总额: %s 实收: %s", time, addr, sessionCount, sessionSellCount, sessionMoneyPrice, sessionMoneyRecv)
		data.Info = append(data.Info, info)
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
