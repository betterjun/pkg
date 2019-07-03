/*
	包logs，日志文件管理模块，基础组件，定义写日志的接口。
*/
package logs

import (
	"fmt"
	"os"
	"os/exec"
	"path"
	"runtime"

	"github.com/astaxie/beego/logs"
)

// 系统全局日志
var sys_logger *logs.BeeLogger

// 错误全局日志
var err_logger *logs.BeeLogger

// 请求全局日志
var req_logger *logs.BeeLogger

// 初始化全局日志对象
func init() {
	file, _ := exec.LookPath(os.Args[0])
	cpath := path.Dir(file)

	sys_logger = NewLogger(path.Join(cpath, "log", "sys_log.txt"), true)
	err_logger = NewLogger(path.Join(cpath, "log", "err_log.txt"), true)
	req_logger = NewLogger(path.Join(cpath, "log", "req_log.txt"), false)
}

// 返回日志对象,以访问所有的方法。
func GetSysLogger() *logs.BeeLogger {
	return sys_logger
}

// 初始化操作对象，并设置控制台显示以及文件记录。
func NewLogger(filename string, console bool) *logs.BeeLogger {
	log := logs.NewLogger(10240) // 缓存大小：10240
	if console {
		log.SetLogger("console", "")
	}
	json := `{"filename":"` + filename + `"}`
	log.SetLogger("file", json)
	return log
}

// skip为需要跳过的堆栈层数
func getStackInfo(skip int) string {
	pc, file, line, _ := runtime.Caller(skip)
	_, f := path.Split(file)
	return fmt.Sprintf("[%s:%d:%s] ", f, line, runtime.FuncForPC(pc).Name())
}

//---------------------------------------------------------------------------
// 全局普通日志便捷访问函数
// 输出跟踪信息。
func Trace(format string, v ...interface{}) {
	sys_logger.Trace(getStackInfo(2)+format, v...)
}

// 输出调试信息。
func Debug(format string, v ...interface{}) {
	sys_logger.Debug(getStackInfo(2)+format, v...)
}

// 输出运行信息。
func Info(format string, v ...interface{}) {
	sys_logger.Info(getStackInfo(2)+format, v...)
}

// 输出错误消息。
func Warn(format string, v ...interface{}) {
	sys_logger.Warn(getStackInfo(2)+format, v...)
}

// 输出错误消息。
func Error(format string, v ...interface{}) {
	newfmt := getStackInfo(2) + format
	sys_logger.Error(newfmt, v...)
	err_logger.Error(newfmt, v...)
}

// 输出危险消息。
func Critical(format string, v ...interface{}) {
	newfmt := getStackInfo(2) + format
	sys_logger.Critical(newfmt, v...)
	err_logger.Critical(newfmt, v...)
}

// 输出Http restful请求消息，在单独的日志文件中记录。
func LogRequest(format string, v ...interface{}) {
	req_logger.Debug(getStackInfo(2)+format, v...)
}
