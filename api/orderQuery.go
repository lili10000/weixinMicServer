package api

import (
	"encoding/json"
	"net/http"
	"time"
	"weixin/dao"
	"fmt"
	. "weixin/entity"
	"weixin/entity/errCode"
)

func QueryOrderByDay(w http.ResponseWriter, req *http.Request) {
	startTime := time.Now().UnixNano() / 1e6
	defer EndRequestLog(startTime, req.URL.Path)

	values := req.URL.Query()
	var countRetn, sellCountRetn int
	if dataStr := values.Get("date"); len(dataStr) > 0 {
		countRetn = dao.QueryOrderDay(dataStr)
		sellCountRetn = dao.QuerySellNumDay(dataStr)
	} else {
		countRetn = dao.QueryOrderToday()
		sellCountRetn = dao.QuerySellNumToday()
	}
	// count := fmt.Sprintf("%d, 销售数 %d", countRetn, sellCountRetn)
	var retn OrderCountRetn
	retn.Code = errCode.Success
	retn.Msg = ""
	retn.Data = OrderCount{
		Count: fmt.Sprintf("%d, 销售数:  %d", countRetn, sellCountRetn),
	}
	retnStr, err := json.Marshal(retn)
	if err != nil {
		w.Write([]byte("系统错误"))
	}
	w.Write([]byte(retnStr))
}
