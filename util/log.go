package util

import (
	"errors"
	"fmt"
	"os"

	"github.com/op/go-logging"
)

var single *logging.Logger
var filePath string
var logFile *os.File

type LogUtil struct{}

func (p *LogUtil) Get() (*logging.Logger, error) {
	if filePath == "" {
		error := errors.New("filePath is null")
		return nil, error
		// filePath = "test.log"
	}
	if single == nil {
		p.InitLog(filePath)
	}
	return single, nil
}

func (p *LogUtil) InitLog(path string) {
	var err error
	single = logging.MustGetLogger("example")
	var format = logging.MustStringFormatter(
		`%{color}%{time:15:04:05.000} %{shortfunc} > %{level:.4s} %{id:03x} %{message}`,
	)

	filePath = path
	logFile, err = os.OpenFile(filePath, os.O_CREATE|os.O_APPEND|os.O_RDWR, 666)
	if err != nil {
		fmt.Println(filePath, path, err)
		os.Exit(3)
	}
	backend1 := logging.NewLogBackend(logFile, "", 0)
	backend1Formatter := logging.NewBackendFormatter(backend1, format)
	backend1Leveled := logging.AddModuleLevel(backend1Formatter)
	backend1Leveled.SetLevel(logging.INFO, "")
	logging.SetBackend(backend1Formatter)

	// backend2 := logging.NewLogBackend(os.Stderr, "", 0)
	// backend2Formatter := logging.NewBackendFormatter(backend2, format)
	// logging.SetBackend(backend1Leveled, backend2Formatter)
}

func (p *LogUtil) Info(detail ...interface{}) {
	if single == nil {
		fmt.Println(detail)
		p.InitLog(filePath)
	}
	single.Info(detail...)
}
func (p *LogUtil) Warning(detail ...interface{}) {
	if single == nil {
		fmt.Println(detail)
		p.InitLog(filePath)
	}
	single.Warning(detail...)
}
func (p *LogUtil) Error(detail ...interface{}) {
	if single == nil {
		fmt.Println(detail)
		p.InitLog(filePath)
	}
	single.Error(detail...)
}
