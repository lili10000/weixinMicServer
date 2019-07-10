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

	// var countRetn, sellCountRetn, moneyPrice, moneyRecv string
	// if dataStr := values.Get("date"); len(dataStr) > 0 {
	// 	countRetn = dao.QuerySellOrderDay(dataStr)
	// 	sellCountRetn = dao.QuerySellNumDay(dataStr)
	// 	moneyPrice = dao.QuerySellMoneyPriceDay(dataStr)
	// 	moneyRecv = dao.QuerySellMoneyRecvDay(dataStr)
	// } else {
	// 	countRetn = dao.QuerySellOrderToday()
	// 	sellCountRetn = dao.QuerySellNumToday()
	// 	moneyPrice = dao.QuerySellMoneyPriceToday()
	// 	moneyRecv = dao.QuerySellMoneyRecvToday()
	// }
	// count := fmt.Sprintf("%d, 销售数 %d", countRetn, sellCountRetn)
	var retn OrderCountRetn
	retn.Code = errCode.Success
	retn.Msg = ""
	retn.Data = OrderCount{
		Count: fmt.Sprintf("    %s\n销售数:     %s\n销售总额:   %s\n实收金额:   %s", countRetn, sellCountRetn, moneyPrice, moneyRecv),
	}
	retnStr, err := json.Marshal(retn)
	if err != nil {
		w.Write([]byte("系统错误"))
	}
	w.Write([]byte(retnStr))
}
