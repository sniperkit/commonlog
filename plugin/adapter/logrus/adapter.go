// Package logging defines an adapter using the popular "logrus" package
package logrus

/*
import (

	// external
	logrus "github.com/sirupsen/logrus"
	prefixed "github.com/x-cray/logrus-prefixed-formatter"

	// internal
	commonlog "github.com/sniperkit/commonlog/pkg"
	adapter "github.com/sniperkit/commonlog/pkg/adapter"
	level "github.com/sniperkit/commonlog/pkg/level"
)

//
// * Types
//

type adapted logrus.Logger

type loggingAdapter struct {
	commonlog.Logger
	adaptee *adapted
}

//
// * Functions
//

// New constructs a new instance by injecting the adaptee
func New(adaptee *logging.Logger) adapter.LogAdapter {
	adapted := adapted(*adaptee)

	adapted.Formatter = new(prefixed.TextFormatter)
	adapted.Level = logrus.DebugLevel

	return &loggingAdapter{
		commonlog.NewLogger(&adapted),
		&adapted,
	}
}

// Adaptee gets the underyling adaptee
func (a *loggingAdapter) Adaptee() interface{} {
	return a.adaptee
}

// Log an error based on a specified level, a format, and a splat of arguments
func (a *adapted) Log(lvl level.LogLevel, format string, args ...interface{}) {
	// Validate the passed in level
	if ok, err := lvl.IsValid(); !ok {
		panic(err)
	}

	adaptee := (*logrus.Logger)(a)

	switch lvl {
	case level.EMERGENCY:
		// TODO: Handle this level once a generic log method is in the logging.Logger
		fallthrough
	case level.ALERT:
		// TODO: Handle this level once a generic log method is in the logging.Logger
		fallthrough
	case level.CRITICAL:
		adaptee.Criticalf(format, args...)
	case level.ERROR:
		adaptee.Errorf(format, args...)
	case level.WARNING:
		adaptee.Warningf(format, args...)
	case level.NOTICE:
		adaptee.Noticef(format, args...)
	case level.INFO:
		adaptee.Infof(format, args...)
	case level.DEBUG:
		adaptee.Debugf(format, args...)
	}
}
*/
