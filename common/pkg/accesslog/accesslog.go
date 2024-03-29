package accesslog

import (
	"github.com/kataras/golog"
	"github.com/kataras/iris/v12/context"
	"github.com/kataras/iris/v12/middleware/accesslog"
	rotatelogs "github.com/lestrrat-go/file-rotatelogs"
	"time"
)

// GetAccessLog 获取accesslog实例,相比iris官方accesslog 多了req_id,user两个一级key
func GetAccessLog(ServerName, Environment, InstanceKey string) *accesslog.AccessLog {
	pathToAccessLog := "./log/access_log.%Y%m%d%H%M"
	writer, err := rotatelogs.New(
		pathToAccessLog,
		rotatelogs.WithMaxAge(24*time.Hour),
		rotatelogs.WithRotationTime(time.Hour),
	)
	if err != nil {
		golog.Fatal(err)
	}

	ac := accesslog.New(writer)

	// 1. Register a field.
	// add req_id
	ac.AddFields(func(context *context.Context, fields *accesslog.Fields) {
		fields.Set("req_id", context.GetID())
	})
	//add user
	ac.AddFields(func(c *context.Context, fields *accesslog.Fields) {
		if len(c.GetHeader("X-User-Name")) == 0 {
			fields.Set("user", "-")
			return
		}
		fields.Set("user", c.GetHeader("X-User_Name"))
	})
	//add query
	ac.AddFields(func(c *context.Context, fields *accesslog.Fields) {
		urlParams := c.Request().URL.RawQuery
		fields.Set("url_params", urlParams)
	})

	// The default configuration:
	ac.TimeFormat = "2006-01-02 15:04:05"
	ac.LatencyRound = time.Millisecond
	ac.IP = true
	ac.Async = true
	ac.Blank = []byte("-")
	ac.BytesReceivedBody = false
	ac.BytesSentBody = false
	ac.BytesReceived = false
	ac.BytesSent = false
	ac.BodyMinify = true
	ac.RequestBody = true
	ac.ResponseBody = true
	ac.KeepMultiLineError = true
	ac.PanicLog = accesslog.LogHandler

	ac.SetFormatter(&UniformJSON{
		HumanTime:   true,
		ServerName:  ServerName,
		Environment: Environment,
		InstanceKey: InstanceKey,
	})

	return ac
}
