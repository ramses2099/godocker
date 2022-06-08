package logging

import (
	"os"

	"github.com/sirupsen/logrus"
)

var AppLog *LoggerClass

type LoggerClass struct {
	Log *logrus.Logger
}

func SetUpLogging() {
	AppLog = &LoggerClass{}
	AppLog.Log = logrus.New()

	file, err := os.OpenFile("logs/logs.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		AppLog.Log.Fatal(err)
	}

	AppLog.Log.SetOutput(file)
	AppLog.Log.SetFormatter(&logrus.JSONFormatter{})

	// AppLog.Log.SetLevel(logrus.WarnLevel)
}

func (logger *LoggerClass) WriteLogsWarn(msg string, fields map[string]interface{}) {
	logger.Log.WithFields(fields).Warn(msg)
}

func (logger *LoggerClass) WriteLogsError(msg string, fields map[string]interface{}) {
	logger.Log.WithFields(fields).Error(msg)
}

func (logger *LoggerClass) WriteLogsInfo(msg string, fields map[string]interface{}) {
	logger.Log.WithFields(fields).Info(msg)
}

func (logger *LoggerClass) WriteLogsDebug(msg string, fields map[string]interface{}) {
	logger.Log.WithFields(fields).Debug(msg)
}
