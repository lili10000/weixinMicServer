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
	countRetn := 0
	if dataStr := values.Get("date"); len(dataStr) > 0 {
		countRetn = dao.QueryOrderDay(dataStr)
	} else {
		countRetn = dao.QueryOrderToday()
	}

	var retn OrderCountRetn
	retn.Code = errCode.Success
	retn.Msg = ""
	retn.Data = OrderCount{
		Count: countRetn,
	}
	retnStr, err := json.Marshal(retn)
	if err != nil {
		w.Write([]byte("系统错误"))
	}
	w.Write([]byte(retnStr))
}
