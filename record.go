package record

import (
	"time"

	"github.com/infinite-iroha/touka"
)

// 日志中间件
// 请保证logger实例被定义
func Middleware() touka.HandlerFunc {
	return func(c *touka.Context) {
		logger := c.GetLogger()
		startTime := time.Now()

		c.Next()

		endTime := time.Now()
		timingResults := endTime.Sub(startTime)

		logger.Infof("%s %s %s %s %s %d %s ", c.ClientIP(), c.Request.Method, c.GetProtocol(), c.Request.URL.Path, c.Request.UserAgent(), c.Writer.Status(), timingResults)
	}
}
