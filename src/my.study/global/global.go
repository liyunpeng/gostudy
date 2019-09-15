package global

import (
	"log"
	"os"
)

var (
	Logger, LoggerFile = getLogger()
)

func getLogger() (*log.Logger,  *os.File) {
	file, _:= os.Create("debuglog.txt")

	logger := log.New(file, "[goStudy] ", log.Ldate|log.Ltime)

	return logger, file
}