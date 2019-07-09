package entity

type ConfigMysql struct {
	Server      string `json:"server"`
	MaxConn     int `json:"maxConn"`
	MaxIdleConn int `json:"maxIdleConn"`
}




type ConfigJson struct {
	AppName  string            `json:"appName"`
	Mysql    ConfigMysql       `json:"mysql"`
}
