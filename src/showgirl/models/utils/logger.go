package utils

import (
	"encoding/json"
	"os"
	"runtime"
	"strconv"
	"strings"

	"github.com/astaxie/beego"
)

//获取go线程id,hack方法慎用
func GetGoID() int64 {
	var buf [64]byte
	n := runtime.Stack(buf[:], false)
	idField := strings.Fields(strings.TrimPrefix(string(buf[:n]), "goroutine "))[0]
	id, err := strconv.ParseInt(idField, 10, 64)
	if err != nil {
		return 0
	}
	return id
}

type LMap map[string]interface{}

/*
Go模块日志规范
go模块日志主要由flowid，事件和细节组成
@param flowid: 请求唯一标识
@param event: 事件概要，请以下划线做分割，以方便ES等时间数据库做搜索和匹配, 请不要包含百分号和竖线
@param details: 事件详情，其类型为LMap，日志程序会将其JSON化之后打印出来
*/
func JsonLog(flowid int64, event string, details LMap) string {
	format := "FlowId=" + strconv.FormatInt(flowid, 10) + "|" + event
	if strings.Contains(event, "%") {
		Fatal(flowid, "jsonLog_event_param_error, event:%s", event)
		return format
	}

	detailBytes, err := json.Marshal(details)
	if err != nil {
		Fatal(flowid, "jsonLog_json_Marshal_failed, err:%s", err)
		return format
	}
	format += "|" + string(detailBytes)
	return format
}

func JFatal(flowid int64, event string, details LMap) {
	beego.BeeLogger.Critical(JsonLog(flowid, event, details))
}

func JWarn(flowid int64, event string, details LMap) {
	beego.BeeLogger.Warn(JsonLog(flowid, event, details))
}

func JInfo(flowid int64, event string, details LMap) {
	beego.BeeLogger.Info(JsonLog(flowid, event, details))
}

func JTrace(flowid int64, event string, details LMap) {
	beego.BeeLogger.Trace(JsonLog(flowid, event, details))

}

func JDebug(flowid int64, event string, details LMap) {
	beego.BeeLogger.Debug(JsonLog(flowid, event, details))
}

func getFormat(flowid int64, format string) string {
	return "FlowId=" + strconv.FormatInt(flowid, 10) + "|" + format
}

func Fatal(flowid int64, format string, v ...interface{}) {
	beego.BeeLogger.Critical(getFormat(flowid, format), v...)
}

func Warn(flowid int64, format string, v ...interface{}) {
	beego.BeeLogger.Warning(getFormat(flowid, format), v...)
}

func Info(flowid int64, format string, v ...interface{}) {
	beego.BeeLogger.Informational(getFormat(flowid, format), v...)
}

func Trace(flowid int64, format string, v ...interface{}) {
	beego.BeeLogger.Trace(getFormat(flowid, format), v...)
}

func Debug(flowid int64, format string, v ...interface{}) {
	beego.BeeLogger.Debug(getFormat(flowid, format), v...)
}

func init() {
	var level int
	logLevel := beego.AppConfig.String("log::level")
	//maxsize:4GB
	beego.SetLogger("file", `{"filename":"./log/showgirl.log","daily":true,"maxdays":7,"maxlines":100000000,"maxsize":4294967296}`)
	if strings.ToUpper(os.Getenv("CONSOLE_LOG")) != "ON" {
		beego.BeeLogger.DelLogger("console")
	} else {
		beego.SetLogger("console", "")
	}
	beego.SetLogFuncCall(true)
	beego.BeeLogger.Debug("logLevel:%s", logLevel)
	switch logLevel {
	case "DEBUG":
		level = beego.LevelDebug
	case "INFO":
		level = beego.LevelInformational
	case "WARN":
		level = beego.LevelWarning
	case "FATAL":
		level = beego.LevelError
	default:
		level = beego.LevelInformational
		beego.BeeLogger.Warning("log level not exist!(DEBUG, INFO, WARN, FATAL)")
	}
	// Log levels:
	// LevelEmergency
	// LevelAlert
	// LevelCritical
	// LevelError
	// LevelWarning
	// LevelNotice
	// LevelInformational
	// LevelDebug
	//beego.SetLevel(beego.LevelDebug)
	beego.SetLevel(level)
}
