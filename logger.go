package Logger

import (
	nested "github.com/antonfisher/nested-logrus-formatter"
	"github.com/gin-gonic/gin"
	rotatelogs "github.com/lestrrat-go/file-rotatelogs"
	"github.com/sirupsen/logrus"
	config "github.com/ynsluhan/go-config"
	"log"
	"os"
	"path"
	"strconv"
	"time"
)

var conf *config.Config

func init() {
	//
	conf = config.GetConf()
	if conf.Server.EnableLogger {
		initLogger()
	}
}

func initLogger() {
	var dir = conf.Server.LoggerDir
	var fileName string
	//
	if len(dir) == 0 {
		getwd, _ := os.Getwd()
		dir = path.Join(getwd, "logs")
	}
	// 日志文件
	fileName = path.Join(dir, "bo.log")
	// 写入文件
	src, err := os.OpenFile(fileName, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatalf("[logger] ERROR 日志文件创建失败 error: %s, file: %s", err, fileName)
	}
	// 实例化
	logger = logrus.New()
	// 设置输出
	logger.Out = src
	// 设置日志级别
	logger.SetLevel(logrus.InfoLevel)
	// 设置普通日志输出模式  json模式注释
	logger.SetFormatter(
		&nested.Formatter{
			// 禁用键值对日志类型
			HideKeys: true,
			// 格式
			TimestampFormat: "2006-01-02 15:04:05",
			// 字段顺序  logger.WithFields(logrus.Fields{"client_ip": ip})  必须要使用该日志模式
			FieldsOrder: []string{"client_ip", "req_method", "req_uri", "status_code", "latency_time"},
			// 日志禁用颜色
			NoColors: true,
		},
	)
	// 设置 rotatelogs
	logWriter, err := rotatelogs.New(
		// 分割后的文件名称
		path.Join(dir, "bo-%Y-%m-%d_%H.log"),
		// 生成软链，指向最新日志文件
		rotatelogs.WithLinkName(fileName),
		// 设置最大保存时间(30天)
		rotatelogs.WithMaxAge(30*24*time.Hour),
		// 设置日志切割时间间隔(1天)
		rotatelogs.WithRotationTime(24*time.Hour),
	)
	// 普通写入设置日志切割  json模式注释
	logger.SetOutput(logWriter)
}

// 使用日志框架logrus日志记录到文件
func DefaultLogger() gin.HandlerFunc {
	// 使用logrus 日志系统
	return func(c *gin.Context) {
		// 开始时间
		//startTime := time.Now()
		start := time.Now()
		// 计算执行时间
		cost := time.Since(start)
		// 请求方式
		reqMethod := c.Request.Method
		// 请求路由
		reqUri := c.Request.RequestURI
		// 状态码
		statusCode := c.Writer.Status()
		// 请求IP
		clientIP := c.ClientIP()
		// 自定义日志
		logger.Info(clientIP+" ", reqMethod+" ", reqUri+" ", strconv.Itoa(statusCode)+" ", cost)
		c.Next()
	}
}
