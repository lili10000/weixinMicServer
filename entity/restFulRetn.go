package entity

type Comm struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
	// data []interface{} `json:"data"`
}

type SessionInfo struct {
	Time       string `json:"time"`
	Addr       string `json:"addr"`
	OrderCount string `json:"orderCount"`
	SellCount  string `json:"sellCount"`
	MoneyPrice string `json:"moneyPrice"`
	MoneyRecv  string `json:"moneyRecv"`
}

type OrderCount struct {
	Sum  SessionInfo   `json:"sum"`
	Info []SessionInfo `json:"detail"`
}

type OrderCountRetn struct {
	Comm
	Data OrderCount `json:"data"`
}
