package main

import logger "github.com/mmungdong/chatgpt-web/pkg/utils/log"

func main() {
	logger.CfgConsoleLogger(true, true)
	logger.Info("hello world")
	logger.Debug("debug hello world")
	logger.Warn("warn hello world")
}
