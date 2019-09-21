package log1

import (
	"github.com/cihub/seelog"
)

func Seelog() {
	logger, err := seelog.LoggerFromConfigAsFile("seelog.xml")

	defer seelog.Flush()

	if err != nil {
		seelog.Critical("err parsing config log file", err)
		return
	}
	seelog.ReplaceLogger(logger)

	seelog.Error("seelog error")
	seelog.Info("seelog info")
	seelog.Debug("seelog debug")
}

func Log()  {
	Seelog()
}