package Logger

import (
	"github.com/gin-gonic/gin"
	"time"
)

/**
 * @Author yNsLuHan
 * @Description: 获取基本参数
 * @Time 2021-06-28 10:31:32
 * @param c
 * @return time.Duration
 * @return string
 * @return string
 * @return int
 * @return string
 */
func LogsFormatArgs(c *gin.Context) (time.Duration, string, string, int, string) {
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
	return cost, reqMethod, reqUri, statusCode, clientIP
}
