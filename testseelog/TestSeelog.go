package main

import (
	"fmt"
	"github.com/cihub/seelog"
	"reflect"
)

func main() {
	defer seelog.Flush()
	logger := getLoger()
	if logger == nil {
		return
	}
	seelog.ReplaceLogger(logger)

	/*seelog.Error("seelog error")
	seelog.Info("seelog info")
	seelog.Debug("seelog debug")*/
	for i := 0; i < 1024; i++ {
		seelog.Error("seelog info: " + fmt.Sprintf("%v", i))
		seelog.Info("seelog info: " + fmt.Sprintf("%v", i))
	}
}

func getLoger() seelog.LoggerInterface {
	logger, err := seelog.LoggerFromConfigAsFile("seelog2.xml")
	if err != nil {
		seelog.Critical("err parsing config log file", err)
		return nil
	}
	return logger
}
