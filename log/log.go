package log

import (
	"os"
	"github.com/Sirupsen/logrus"
	"github.com/matyix/echo/conf"
)

// Logger defines a set of methods for writing application logs. Derived from and
// inspired by logrus.Entry.
type Logger interface {
	Debug(args ...interface{})
	Debugf(format string, args ...interface{})
	Debugln(args ...interface{})
	Error(args ...interface{})
	Errorf(format string, args ...interface{})
	Errorln(args ...interface{})
	Fatal(args ...interface{})
	Fatalf(format string, args ...interface{})
	Fatalln(args ...interface{})
	Info(args ...interface{})
	Infof(format string, args ...interface{})
	Infoln(args ...interface{})
	Panic(args ...interface{})
	Panicf(format string, args ...interface{})
	Panicln(args ...interface{})
	Print(args ...interface{})
	Printf(format string, args ...interface{})
	Println(args ...interface{})
	Warn(args ...interface{})
	Warnf(format string, args ...interface{})
	Warning(args ...interface{})
	Warningf(format string, args ...interface{})
	Warningln(args ...interface{})
	Warnln(args ...interface{})
}

var logger *logrus.Logger

func init() {
	logger = newLogrusLogger(conf.Config())
}

func NewLogger(cfg conf.Provider) *logrus.Logger {
	return newLogrusLogger(cfg)
}

func newLogrusLogger(cfg conf.Provider) *logrus.Logger {
	l := logrus.New()
	if cfg.GetBool("json_logs") {
		l.Formatter = new(logrus.JSONFormatter)
	}
	l.Out = os.Stderr

	switch cfg.GetString("loglevel") {
	case "debug":
		l.Level = logrus.DebugLevel
	case "warning":
		l.Level = logrus.WarnLevel
	case "info":
		l.Level = logrus.InfoLevel
	default:
		l.Level = logrus.DebugLevel
	}
	return l
}

type Fields map[string]interface{}

func (f Fields) WithKV(k string, v interface{}) Fields {
	f[k] = v
	return f
}

func (f Fields) WithFields(fields Fields) Fields {
	for k, v := range fields {
		f[k] = v
	}
	return f
}

func WithFields(fields Fields) Logger {
	return logger.WithFields(logrus.Fields(fields))
}

//Generated
// Debug package-level convenience method.
func Debug(args ...interface{}) {
	logger.Debug(args...)
}

// Debugf package-level convenience method.
func Debugf(format string, args ...interface{}) {
	logger.Debugf(format, args...)
}

// Debugln package-level convenience method.
func Debugln(args ...interface{}) {
	logger.Debugln(args...)
}

// Error package-level convenience method.
func Error(args ...interface{}) {
	logger.Error(args...)
}

// Errorf package-level convenience method.
func Errorf(format string, args ...interface{}) {
	logger.Errorf(format, args...)
}

// Errorln package-level convenience method.
func Errorln(args ...interface{}) {
	logger.Errorln(args...)
}

// Fatal package-level convenience method.
func Fatal(args ...interface{}) {
	logger.Fatal(args...)
}

// Fatalf package-level convenience method.
func Fatalf(format string, args ...interface{}) {
	logger.Fatalf(format, args...)
}

// Fatalln package-level convenience method.
func Fatalln(args ...interface{}) {
	logger.Fatalln(args...)
}

// Info package-level convenience method.
func Info(args ...interface{}) {
	logger.Info(args...)
}

// Infof package-level convenience method.
func Infof(format string, args ...interface{}) {
	logger.Infof(format, args...)
}

// Infoln package-level convenience method.
func Infoln(args ...interface{}) {
	logger.Infoln(args...)
}

// Panic package-level convenience method.
func Panic(args ...interface{}) {
	logger.Panic(args...)
}

// Panicf package-level convenience method.
func Panicf(format string, args ...interface{}) {
	logger.Panicf(format, args...)
}

// Panicln package-level convenience method.
func Panicln(args ...interface{}) {
	logger.Panicln(args...)
}

// Print package-level convenience method.
func Print(args ...interface{}) {
	logger.Print(args...)
}

// Printf package-level convenience method.
func Printf(format string, args ...interface{}) {
	logger.Printf(format, args...)
}

// Println package-level convenience method.
func Println(args ...interface{}) {
	logger.Println(args...)
}

// Warn package-level convenience method.
func Warn(args ...interface{}) {
	logger.Warn(args...)
}

// Warnf package-level convenience method.
func Warnf(format string, args ...interface{}) {
	logger.Warnf(format, args...)
}

// Warning package-level convenience method.
func Warning(args ...interface{}) {
	logger.Warning(args...)
}

// Warningf package-level convenience method.
func Warningf(format string, args ...interface{}) {
	logger.Warningf(format, args...)
}

// Warningln package-level convenience method.
func Warningln(args ...interface{}) {
	logger.Warningln(args...)
}

// Warnln package-level convenience method.
func Warnln(args ...interface{}) {
	logger.Warnln(args...)
}
