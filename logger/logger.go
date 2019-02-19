package logger

import (
	"github.com/sirupsen/logrus"
)

// Fatalf ...
func Fatalf(s string, i ...interface{}) {
	logrus.Fatalf(s, i...)
}

// Fatal ...
func Fatal(s string) {
	logrus.Fatal(s)
}

// Errorf ...
func Errorf(s string, i ...interface{}) {
	logrus.Errorf(s, i...)
}

// Error ...
func Error(s string) {
	logrus.Errorf(s)
}

// Infof ...
func Infof(s string, i ...interface{}) {
	logrus.Infof(s, i...)
}

// Info ...
func Info(s string) {
	logrus.Infof(s)
}

// Printf ...
func Printf(s string, i ...interface{}) {
	logrus.Infof(s, i...)
}

// Print ...
func Print(s string) {
	logrus.Infof(s)
}

// Println ...
func Println(i ...interface{}) {
	logrus.Infoln(i...)
}

// Warnf ...
func Warnf(s string, i ...interface{}) {
	logrus.Warnf(s, i...)
}

// Warn ...
func Warn(s string) {
	logrus.Warnf(s)
}

// Set modifies the logger
func Set(hook ContextHook, setReport bool) {
	logrus.SetReportCaller(setReport)
}
