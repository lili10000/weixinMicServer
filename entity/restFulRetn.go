package entity

type Comm struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
	// data []interface{} `json:"data"`
}

type OrderCount struct {
	Sum    string   `json:"sum"`
	Detail []string `json:"detail"`
}

type OrderCountRetn struct {
	Comm
	Data OrderCount `json:"data"`
}
