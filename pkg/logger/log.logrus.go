package logger

import (
	"fmt"
	"os"
	"runtime"

	"github.com/sirupsen/logrus"
)

var logger *logrus.Logger
var loggerEntry *logrus.Entry

func init() {
	customFormatter := &logrus.TextFormatter{
		ForceColors:     true,
		FullTimestamp:   true,
		TimestampFormat: "2006-01-02 15:04:05",
	}
	customFormatter.TimestampFormat = "2006-01-02 15:04:05.00000"
	logger = &logrus.Logger{
		Out:          os.Stdout,
		ExitFunc:     os.Exit,
		Formatter:    customFormatter,
		Hooks:        make(logrus.LevelHooks),
		Level:        logrus.DebugLevel,
		ReportCaller: false,
	}
	loggerEntry = logrus.NewEntry(logger)
}

func Info(message string, fields logrus.Fields) {
	logger.WithFields(fields).Info(message)
}

func InfoF(format string, fields logrus.Fields, args ...interface{}) {
	logger.WithFields(fields).Infof(format, args...)
}

// Debug logs a message at level Debug on the standard logger.
func Debug(args ...interface{}) {
	if logger.Level >= logrus.DebugLevel {
		loggerEntry.WithField("file", fileInfo(2)).Debug(args...)
	}
}

func DebugF(format string, fields logrus.Fields, args ...interface{}) {
	logger.WithFields(fields).Debugf(format, args...)
}

func Error(message string, fields logrus.Fields) {
	logger.WithFields(fields).Error(message)
}

func ErrorF(format string, fields logrus.Fields, args ...interface{}) {
	logger.WithFields(fields).Errorf(format, args...)
}

func Fatal(message string, fields logrus.Fields) {
	logger.WithFields(fields).Fatal(message)
}

func FatalF(format string, fields logrus.Fields, args ...interface{}) {
	logger.WithFields(fields).Fatalf(format, args...)
}

func Panic(message string, fields logrus.Fields) {
	logger.WithFields(fields).Panic(message)
}

func PanicF(format string, fields logrus.Fields, args ...interface{}) {
	logger.WithFields(fields).Panicf(format, args...)
}

func fileInfo(skip int) string {
	_, file, line, ok := runtime.Caller(skip)
	if !ok {
		file = "<???>"
		line = 1
	}
	// else {
	//	slash := strings.LastIndex(file, "/")
	//	if slash >= 0 {
	//		file = file[slash+1:]
	//	}
	// }
	return fmt.Sprintf("%s:%d", file, line)
}
