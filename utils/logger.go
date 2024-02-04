/*
Copyright © 2023 jun<simpleyuan@gmail.com>
*/
package utils

import (
	"io"
	"time"

	rotatelogs "github.com/lestrrat-go/file-rotatelogs"
	"github.com/rifflock/lfshook"
	"github.com/sirupsen/logrus"
)

var (
	Logger   = logrus.New() // 初始化日志对象
	LogEntry *logrus.Entry
)

func init() {
	// 写入日志文件
	logPath := "runtime/log" // 日志存放路径
	// linkName := "logs/latest.log" // 最新日志的软连接路径
	/*
		src, err := os.OpenFile(logPath, os.O_RDWR|os.O_CREATE, 0755) // 初始化日志文件对象
		if err != nil {
			fmt.Println("err: ", err)
		}
	*/
	//log := logrus.New()  // 初始化日志对象
	Logger.Out = io.Discard // 把产生的日志内容写进日志文件中

	// 日志分隔：1. 每天产生的日志写在不同的文件；2. 只保留一定时间的日志（例如：一星期）
	Logger.SetLevel(logrus.DebugLevel) // 设置日志级别
	logWriter, _ := rotatelogs.New(
		logPath+"%Y%m%d%H.log",                   // 日志文件名格式
		rotatelogs.WithMaxAge(7*24*time.Hour),    // 最多保留7天之内的日志
		rotatelogs.WithRotationTime(1*time.Hour), // 一天保存一个日志文件
		// rotatelogs.WithLinkName(linkName),        // 为最新日志建立软连接
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
		TimestampFormat: "2006-01-02 15:04:05", // 格式日志时间
	})
	Logger.AddHook(Hook)
	LogEntry = logrus.NewEntry(Logger).WithField("service", "imoowi-comer")
}
