package util

import (
	"github.com/op/go-logging"
	"errors"
	"os"
	"fmt"
)


var single = logging.MustGetLogger("example")
var format = logging.MustStringFormatter(
    `%{color}%{time:15:04:05.000} %{shortfunc} > %{level:.4s} %{id:03x}%{color:reset} %{message}`,
)


type LogUtil struct {
	FilePath   string  // 日志文件路径和名称
}

func (p *LogUtil) Get() (*logging.Logger, error){
	if p.FilePath == "" {
		error := errors.New("filePath is null")
		return nil, error
	}
	if single == nil {
		p.Init(p.FilePath)
	}
	return single, nil
}

func (p *LogUtil) Init(filePath string) {
	p.FilePath = filePath
	logFile, err := os.OpenFile(p.FilePath, os.O_WRONLY,0666)
    if err != nil{
        fmt.Println(err)
    }
	defer logFile.Close()
    backend1 := logging.NewLogBackend(logFile, "", 0)
    backend2 := logging.NewLogBackend(os.Stderr, "", 0)
 
    backend2Formatter := logging.NewBackendFormatter(backend2, format)
    backend1Leveled := logging.AddModuleLevel(backend1)
    backend1Leveled.SetLevel(logging.INFO, "")
 
    logging.SetBackend(backend1Leveled, backend2Formatter)
}
