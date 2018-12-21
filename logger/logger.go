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

// Printf ...
func Printf(s string, i ...interface{}) {
	logrus.Infof(s, i...)
}

// Print ...
func Print(s string) {
	logrus.Infof(s)
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
	logrus.AddHook(hook)
}
