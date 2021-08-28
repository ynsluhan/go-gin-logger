package Logger

import (
	"github.com/sirupsen/logrus"
	"log"
	"os"
	"path"
)

// 日志实体类
var logger *logrus.Logger

// 获取日志路径
func GetLogPath() string {
	getwd, err := os.Getwd()
	//
	if err != nil {
		log.Fatal("ERROR 获取项目路径失败:", err)
	}
	//
	return path.Join(getwd, "logs")
}

// 获取logger
func GetLogger() *logrus.Logger {
	return logger
}
