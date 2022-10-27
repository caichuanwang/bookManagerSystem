package log

import (
	"github.com/sirupsen/logrus"
	rotate "gopkg.in/natefinch/lumberjack.v2"
	"strings"
)

var Std *logrus.Logger

func Init() {
	std1 := logrus.New()
	std1.SetReportCaller(true)
	std1.SetFormatter(&logrus.JSONFormatter{TimestampFormat: "2006-01-02 15:04:05"})
	std1.WithField("key", "11111")
	Std = std1
	InitLogger("log/out")
}
func InitLogger(filePath string) {
	Std.SetOutput(NewRotateFile(filePath, "/bms.log", 1 /**/))
}

func NewRotateFile(filePath, fileName string, maxSize int) *rotate.Logger {
	return &rotate.Logger{
		Filename: strings.TrimRight(filePath, "/") + fileName,
		MaxSize:  maxSize,
		MaxAge:   30,
	}
}
