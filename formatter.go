// Package easy allows to easily format output of Logrus logger
package easy

import (
	"strconv"
	"strings"

	"github.com/sirupsen/logrus"
)

const (
	// Default log format will output [INFO] 2006-01-02 15:04:05.9999 path/file Line:1 - Log message
	defaultLogFormat       = "[%level%] %time% %func% Line:%line% - %msg%\n"
	defaultTimestampFormat = "2006-01-02 15:04:05.9999"
)

// Formatter implements logrus.Formatter interface.
type Formatter struct {
	// Timestamp format
	TimestampFormat string
	// Available standard keys: time, msg, level, func, line
	// Also can include custom fields but limited to strings.
	// All of fields need to be wrapped inside %% i.e %time% %msg%
	LogFormat string
}

// Format building log message.
func (f *Formatter) Format(entry *logrus.Entry) ([]byte, error) {
	output := f.LogFormat
	if output == "" {
		output = defaultLogFormat
	}

	timestampFormat := f.TimestampFormat
	if timestampFormat == "" {
		timestampFormat = defaultTimestampFormat
	}

	output = strings.Replace(output, "%time%", entry.Time.Format(timestampFormat), 1)
	output = strings.Replace(output, "%msg%", entry.Message, 1)
	level := strings.ToUpper(entry.Level.String())
	output = strings.Replace(output, "%level%", level, 1)
	output = strings.Replace(output, "%func%", entry.Caller.Function, 1)
	output = strings.Replace(output, "%line%", strconv.Itoa(entry.Caller.Line), 1)

	for k, val := range entry.Data {
		switch v := val.(type) {
		case string:
			output = strings.Replace(output, "%"+k+"%", v, 1)
		case int:
			s := strconv.Itoa(v)
			output = strings.Replace(output, "%"+k+"%", s, 1)
		case bool:
			s := strconv.FormatBool(v)
			output = strings.Replace(output, "%"+k+"%", s, 1)
		}
	}

	return []byte(output), nil
}
