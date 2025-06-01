package record

import (
	"time"

	"github.com/WJQSERVER-STUDIO/logger"
	"github.com/infinite-iroha/touka"
)

var (
	logw       = logger.Logw
	logDump    = logger.LogDump
	logDebug   = logger.LogDebug
	logInfo    = logger.LogInfo
	logWarning = logger.LogWarning
	logError   = logger.LogError
)

// 日志中间件
// 请保证logger实例被定义
func Middleware() touka.HandlerFunc {
	return func(c *touka.Context) {
		startTime := time.Now()

		c.Next()

		endTime := time.Now()
		timingResults := endTime.Sub(startTime)

		logInfo("%s %s %s %s %s %d %s ", c.ClientIP(), c.Request.Method, c.GetProtocol(), c.Request.URL.Path, c.Request.UserAgent(), c.Writer.Status(), timingResults)
	}
}
