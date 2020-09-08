package logutil

import (
	"bytes"
	"github.com/go-kit/kit/log/level"
	"github.com/pkg/errors"
	"io"
	"os"
	"strings"
	"testing"
	"time"
)

var buffer = new(bytes.Buffer)

func mockOpenFile(name string, _ int, _ os.FileMode) (io.Writer, error) {
	if name == "test.log" {
		return buffer, nil
	} else {
		return nil, io.EOF
	}
}

var mockStdout = buffer

func mockLogTime() time.Time {
	return time.Time{}
}

func init() {
	osOpenFile = mockOpenFile
	osStdout = mockStdout
	fnLogTime = mockLogTime
}

func IsSameError(e1, e2 error) bool {
	if e1 == nil {
		return e2 == nil
	} else {
		if e2 == nil {
			return false
		} else {
			return e1.Error() == e2.Error()
		}
	}
}

func TestNewLogger(t *testing.T) {

	var data = []struct {
		config   LoggerConfig
		err      error
		expected string
	}{
		{
			config: LoggerConfig{
				LogType: ConsoleLogger,
				LogLevel: InfoLevel,
			},
			expected: `{"level":"warn","step":"0","ts":"0001-01-01T00:00:00Z"}
{"level":"info","step":"1","ts":"0001-01-01T00:00:00Z"}`,
		},
		{
			config: LoggerConfig{
				LogType: FileLogger,
				LogLevel: InfoLevel,
				Properties: map[string]string{
					"file": "test.log",
				},
			},
			expected: `{"level":"warn","step":"0","ts":"0001-01-01T00:00:00Z"}
{"level":"info","step":"1","ts":"0001-01-01T00:00:00Z"}`,
		},
		{
			config: LoggerConfig{
				LogType: FileLogger,
				LogLevel: InfoLevel,
				Properties: map[string]string{
					"file": "fake.log",
				},
			},
			err: errors.Wrap(io.EOF, "cannot open file fake.log"),
			expected: "",
		},
	}

	t.Log("Given the need to test logger.")
	{

		for _, d := range data {
			buffer.Reset()
			t.Logf("\tWhen using \"%+v\" ", d)
			{
				logger, err := NewLogger(d.config)
				if IsSameError(err, d.err) {
					t.Logf("\tShould expect to create logger %+v.", d.config)
				} else {
					t.Errorf("\tShould expect to create console logger %+v but received error: %v.", d.config, err)
				}

				if logger == nil { // if err received then logger is nil
					return
				}

				level.Warn(logger).Log("step", "0")
				level.Info(logger).Log("step", "1")
				level.Debug(logger).Log("step", "2")

				actual := strings.TrimSpace(buffer.String())
				if actual == d.expected {
					t.Logf("\tShould expect to write log \"%s\"", actual)
				} else {
					t.Errorf("\tShould expect to write log \"%s\" but received \"%s\"", d.expected, actual)
				}
			}
		}
	}
}
