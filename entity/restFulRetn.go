package entity

type Comm struct {
	Code  int        `json:"code"`
	Msg  string        `json:"msg"`
	// data []interface{} `json:"data"`
}

type OrderCount struct{
	Count string `json:"count"`
}

type OrderCountRetn struct{
	Comm
	Data OrderCount `json:"data"`
}