/*
Copyright © 2023 yuanjun<simpleyuan@gmail.com>
*/
package global

import (
	"io"
	"time"

	rotatelogs "github.com/lestrrat-go/file-rotatelogs"
	"github.com/rifflock/lfshook"
	"github.com/sirupsen/logrus"
)

var Log *logrus.Entry

func init() {
	Log = logrus.NewEntry(logrus.StandardLogger())
}

func RegisterLogger(l *logrus.Entry) {
	Log = l
}

func initLog() error {
	// 获取logger相关的配置信息
	config := Config.Sub("logger")

	logger := logrus.New()

	// 设定写入日志的存放路径
	logPath := config.GetString("path") // 日志存放路径
	if logPath == "" {
		logPath = "runtime/log"
	}
	logger.Out = io.Discard // 把产生的日志内容写进日志文件中

	// 设定日志存放的级别
	logLevel := config.GetString("level")
	logLevelInt, err := logrus.ParseLevel(logLevel)
	if err != nil {
		logLevelInt = logrus.DebugLevel // 级别设置错误时，默认设置为debug级别
	}

	// 获取保存周期，最小为7天
	maxAge := config.GetDuration("maxAge")
	if maxAge < 7*24*time.Hour {
		maxAge = 7 * 24 * time.Hour
	}

	// 获取切割周期，最小为一个小时
	rotatimeTime := config.GetDuration("rotationTime")
	if rotatimeTime < time.Hour {
		rotatimeTime = time.Hour
	}

	// 日志分隔：1. 每天产生的日志写在不同的文件；2. 只保留一定时间的日志（例如：一星期）
	logger.SetLevel(logLevelInt) // 设置日志级别
	logWriter, _ := rotatelogs.New(
		logPath+"%Y%m%d%H.log",                    // 日志文件名格式
		rotatelogs.WithMaxAge(maxAge),             // 设置最大保存周期
		rotatelogs.WithRotationTime(rotatimeTime), // 设置日志切割周期，最小为1小时
	)
	writeMap := lfshook.WriterMap{
		logrus.PanicLevel: logWriter,
		logrus.FatalLevel: logWriter,
		logrus.ErrorLevel: logWriter,
		logrus.WarnLevel:  logWriter,
		logrus.InfoLevel:  logWriter, // info级别使用logWriter写日志
		logrus.DebugLevel: logWriter,
		logrus.TraceLevel: logWriter,
	}
	Hook := lfshook.NewHook(writeMap, &logrus.TextFormatter{
		TimestampFormat: time.RFC3339Nano, // 格式日志时间
	})
	logger.AddHook(Hook)
	Log = logrus.NewEntry(logger).WithField("service", ".moduleName")
	// App.SetLogger(Log)
	return nil
}
