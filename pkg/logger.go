/**
 * CommonLog
 *
 * Copyright © 2015 Trevor N. Suarez (Rican7)
 */

// Package commonlog defines interfaces for common logging interactions, with an
// emphasis on Syslog level compatibility
package commonlog

import (
	level "github.com/sniperkit/commonlog/pkg/level"
)

/**
 * Types
 */

// A Logger defines the method set for a common logging implementation, with
// methods defined for each standard log level as defined in package level
type Logger interface {
	LevelLogger

	Emergency(string, ...interface{})
	Alert(string, ...interface{})
	Critical(string, ...interface{})
	Error(string, ...interface{})
	Warning(string, ...interface{})
	Notice(string, ...interface{})
	Info(string, ...interface{})
	Debug(string, ...interface{})
}

// LevelLogger is an interface containing a single method that logs a message
// and context/formatting values to a specified a level
type LevelLogger interface {
	Log(level.LogLevel, string, ...interface{})
}
