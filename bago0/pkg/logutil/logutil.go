package logutil

import (
	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/log/level"
	"github.com/pkg/errors"
	"io"
	"os"
	"time"
)

type LogType string

const (
	FileLogger    = "file"
	ConsoleLogger = "console"
)

const (
	DebugLevel = "debug"
	InfoLevel  = "info"
	WarnLevel  = "warn"
	ErrorLevel = "error"
)

type LoggerConfig struct {
	LogType    LogType           `yaml:"logType"`
	LogLevel   string            `yaml:"logLevel"`
	Properties map[string]string `yaml:"properties,omitempty"`
}

var osStdout io.Writer = os.Stdout

var osOpenFile = func (name string, flag int, perm os.FileMode) (io.Writer, error) {
	f, err := os.OpenFile(name, flag, perm)
	return f, err
}

var fnLogTime = func() time.Time { return time.Now().UTC() }

func newLoggerWriter(c LoggerConfig) (log.Logger, error) {

	var writer io.Writer
	switch c.LogType {
	case FileLogger:
		fn, found := c.Properties[FileLogger]
		if !found {
			fn = os.Args[0] + ".log"
		}
		f, err := osOpenFile(fn, os.O_RDWR|os.O_CREATE, 0755)
		if err != nil {
			return nil, errors.Wrap(err, "cannot open file "+fn)
		}
		writer = log.NewSyncWriter(f)
	default:
		writer = osStdout
	}

	logger := log.NewJSONLogger(writer)
	logger = log.With(logger, "ts",
		log.TimestampFormat(
			fnLogTime,
			time.RFC3339,
		))
	return logger, nil
}

func NewLogger(c LoggerConfig) (log.Logger, error) {
	l, err := newLoggerWriter(c)
	if err != nil {
		return nil, err
	}
	var o level.Option
	switch c.LogLevel {
	case DebugLevel:
		o = level.AllowDebug()
	case InfoLevel:
		o = level.AllowInfo()
	case WarnLevel:
		o = level.AllowWarn()
	case ErrorLevel:
		o = level.AllowError()
	default:
		o = level.AllowAll()
	}
	l = level.NewFilter(l, o)
	return l, nil
}
