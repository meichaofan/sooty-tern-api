package app

import (
	"os"
	"path/filepath"
	"sooty-tern/internal/app/config"
	"sooty-tern/pkg/logger"
	"sooty-tern/pkg/util"
)

// InitLogger 初始化日志
func InitLogger() (func(), error) {
	logger.SetTraceIDFunc(util.MustUUID)
	c := config.GetGlobalConfig().Log
	logger.SetLevel(c.Level)
	logger.SetFormatter(c.Format)
	var file *os.File
	if c.Output != "" {
		switch c.Output {
		case "stdout":
			logger.SetOutput(os.Stdout)
		case "stderr":
			logger.SetOutput(os.Stderr)
		case "file":
			if name := c.OutputFile; name != "" {
				date := util.FormatCurrentDay("-%d-%d-%d")
				logFileName := name + date + ".log"
				os.MkdirAll(filepath.Dir(logFileName), 0777)
				f, err := os.OpenFile(logFileName, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0666)
				if err != nil {
					return nil, err
				}
				logger.SetOutput(f)
				file = f
			}
		}
	}
	return func() {
		if file != nil {
			file.Close()
		}
	}, nil
}
