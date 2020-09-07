package logutil

import (
	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/log/level"
	"github.com/pkg/errors"
	"io"
	"os"
)

type LogType string

const (
	FileLogger = "file"
	DebugLevel = "debug"
	InfoLevel = "info"
	WarnLevel = "warn"
	ErrorLevel = "error"
)

type LoggerConfig struct {
	LogType LogType `yaml:"logType"`
	LogLevel string `yaml:"logLevel"`
	Properties map[string]string `yaml:"properties,omitempty"`
}

func newLoggerWriter(c LoggerConfig) (log.Logger, error) {

	var writer io.Writer
	switch c.LogType {
	case FileLogger:
		fn, found := c.Properties[FileLogger]
		if !found {
			fn = os.Args[0] + ".log"
		}
		f, err := os.OpenFile(fn, os.O_RDWR|os.O_CREATE, 0755)
		if err != nil {
			return nil, errors.Wrap(err, "cannot open file " + fn)
		}
		writer = log.NewSyncWriter(f)
	default:
		writer = os.Stdout
	}

	logger := log.NewJSONLogger(writer)
	logger = log.With(logger, "ts", log.DefaultTimestampUTC)
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