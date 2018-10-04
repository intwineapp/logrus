package logrus

import (
	"fmt"
	"runtime"
	"strings"

	"github.com/sirupsen/logrus"
)

const runtimeSkip = 2

type logger struct {
	logger *logrus.Logger
}

// Logger interface
type Logger interface {
	SetLogFormatter(formatter logrus.Formatter)
	Debug(args ...interface{})
	DebugWithFields(l interface{}, f Fields)
	Info(args ...interface{})
	InfoWithFields(l interface{}, f Fields)
	Warn(args ...interface{})
	WarnWithFields(l interface{}, f Fields)
	Error(args ...interface{})
	ErrorWithFields(l interface{}, f Fields)
	Fatal(args ...interface{})
	FatalWithFields(l interface{}, f Fields)
	Panic(args ...interface{})
	PanicWithFields(l interface{}, f Fields)
	fileInfo() string
}

// Fields wraps logrus.Fields, which is a map[string]interface{}
type Fields logrus.Fields

// New return a new Logger
func New() Logger {
	log := logrus.New()
	logger := &logger{log}
	return logger
}

func (lg *logger) SetLogLevel(level logrus.Level) {
	lg.logger.Level = level
}

func (lg *logger) SetLogFormatter(formatter logrus.Formatter) {
	lg.logger.Formatter = formatter
}

// Debug logs a message at level Debug on the standard logger.
func (lg *logger) Debug(args ...interface{}) {
	if lg.logger.Level >= logrus.DebugLevel {
		entry := lg.logger.WithFields(logrus.Fields{})
		entry.Data["file"] = lg.fileInfo()
		entry.Debug(args...)
	}
}

// DebugWithFields logs a message with fields at level Debug on the standard logger.
func (lg *logger) DebugWithFields(l interface{}, f Fields) {
	if lg.logger.Level >= logrus.DebugLevel {
		entry := lg.logger.WithFields(logrus.Fields(f))
		entry.Data["file"] = lg.fileInfo()
		entry.Debug(l)
	}
}

// Info logs a message at level Info on the standard logger.
func (lg *logger) Info(args ...interface{}) {
	if lg.logger.Level >= logrus.InfoLevel {
		entry := lg.logger.WithFields(logrus.Fields{})
		entry.Data["file"] = lg.fileInfo()
		entry.Info(args...)
	}
}

// InfoWithFields logs a message with fields at level Info on the standard logger.
func (lg *logger) InfoWithFields(l interface{}, f Fields) {
	if lg.logger.Level >= logrus.InfoLevel {
		entry := lg.logger.WithFields(logrus.Fields(f))
		entry.Data["file"] = lg.fileInfo()
		entry.Info(l)
	}
}

// Warn logs a message at level Warn on the standard logger.
func (lg *logger) Warn(args ...interface{}) {
	if lg.logger.Level >= logrus.WarnLevel {
		entry := lg.logger.WithFields(logrus.Fields{})
		entry.Data["file"] = lg.fileInfo()
		entry.Warn(args...)
	}
}

// WarnWithFields logs a message with fields at level Warn on the standard logger.
func (lg *logger) WarnWithFields(l interface{}, f Fields) {
	if lg.logger.Level >= logrus.WarnLevel {
		entry := lg.logger.WithFields(logrus.Fields(f))
		entry.Data["file"] = lg.fileInfo()
		entry.Warn(l)
	}
}

// Error logs a message at level Error on the standard logger.
func (lg *logger) Error(args ...interface{}) {
	if lg.logger.Level >= logrus.ErrorLevel {
		entry := lg.logger.WithFields(logrus.Fields{})
		entry.Data["file"] = lg.fileInfo()
		entry.Error(args...)
	}
}

// ErrorWithFields logs a message with fields at level Error on the standard logger.
func (lg *logger) ErrorWithFields(l interface{}, f Fields) {
	if lg.logger.Level >= logrus.ErrorLevel {
		entry := lg.logger.WithFields(logrus.Fields(f))
		entry.Data["file"] = lg.fileInfo()
		entry.Error(l)
	}
}

// Fatal logs a message at level Fatal on the standard logger.
func (lg *logger) Fatal(args ...interface{}) {
	if lg.logger.Level >= logrus.FatalLevel {
		entry := lg.logger.WithFields(logrus.Fields{})
		entry.Data["file"] = lg.fileInfo()
		entry.Fatal(args...)
	}
}

// FatalWithFields logs a message with fields at level Fatal on the standard logger.
func (lg *logger) FatalWithFields(l interface{}, f Fields) {
	if lg.logger.Level >= logrus.FatalLevel {
		entry := lg.logger.WithFields(logrus.Fields(f))
		entry.Data["file"] = lg.fileInfo()
		entry.Fatal(l)
	}
}

// Panic logs a message at level Panic on the standard logger.
func (lg *logger) Panic(args ...interface{}) {
	if lg.logger.Level >= logrus.PanicLevel {
		entry := lg.logger.WithFields(logrus.Fields{})
		entry.Data["file"] = lg.fileInfo()
		entry.Panic(args...)
	}
}

// PanicWithFields logs a message with fields at level Debug on the standard logger.
func (lg *logger) PanicWithFields(l interface{}, f Fields) {
	if lg.logger.Level >= logrus.PanicLevel {
		entry := lg.logger.WithFields(logrus.Fields(f))
		entry.Data["file"] = lg.fileInfo()
		entry.Panic(l)
	}
}

// GetLogrusLogger return the underlying logrus.Logger
func (lg *logger) GetLogrusLogger() *logrus.Logger {
	return lg.logger
}

// filInfo extract the fileInfo
func (lg *logger) fileInfo() string {
	_, file, line, ok := runtime.Caller(runtimeSkip)
	if !ok {
		file = "<???>"
		line = 1
	} else {
		slash := strings.LastIndex(file, "/")
		if slash >= 0 {
			file = file[slash+1:]
		}
	}
	return fmt.Sprintf("%s:%d", file, line)
}
