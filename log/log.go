package log

import (
	"github.com/sirupsen/logrus"
)

// Fatalf ...
func Fatalf(s string, i ...interface{}) {
	logrus.Fatalf(s, i...)
}

// Errorf ...
func Errorf(s string, i ...interface{}) {
	logrus.Errorf(s, i...)
}

// Printf ...
func Printf(s string, i ...interface{}) {
	logrus.Errorf(s, i...)
}

func init() {
	logrus.SetReportCaller(true)
}
