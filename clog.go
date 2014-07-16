/*
Colorful logger with support for different log levels.

Five log levels are predefined: DEBUG, INFO, WARNING, ERROR and PANIC. The
logger will only output log messages with level equal to or higher than
limit specified in setup. The messages can optionally be shown in color
(turned off by default).

All messages are shown with a RFC3339 timestamp.

The logger can be setup directly using Setup(). Alternatively, using
SetupFromEnv(), the settings can be picked from environment variables
LOG_LEVEL (should be one of the predefined levle names) and
LOG_COLOR (should be "true" or "false").

The logger provides Log() function which takes a level, and a message. The
convenience functions Debug(), Info(), Warning(), Error() and Panic() are
also provided. There are also variants of these functions which support
passing a format string and arguments instead of a single message string:
Logf(), Debugf(), Infof(), Warningf(), Errorf() and Panicf().

When logging a message with a PANIC level, the logger will raise a panic
with the specified message immediately after logging it.

The output by default goes to os.Stderr. This can be changed by using
SetOutput(). Note that SetOutput() must be called after Setup() (or
SetupFromEnv()).

Example use:

    import "clog"

    clog.Setup(clog.WARNING, true)
    clog.Debug("hello world!")
    clog.Warning("Hello")
    clog.Panicf("The end is %s!", "nigh")
*/
package clog

import (
	"fmt"
	"io"
	"os"
	"strings"
	"time"
)

type LogLevel int

// Available log levels
const (
	DEBUG LogLevel = iota
	INFO
	WARNING
	ERROR
	PANIC
)

var colorCodes = [PANIC + 1]string{
	"\x1b[34m",
	"",
	"\x1b[33m",
	"\x1b[31m",
	"\x1b[1;31m",
}

var levelNames = [PANIC + 1]string{
	"DEBUG",
	"INFO",
	"WARNING",
	"ERROR",
	"PANIC",
}

const noColor = "\x1b[0m"

type config struct {
	level    LogLevel
	useColor bool
	output   io.Writer
}

var cfg config

// Setup sets up the logger using provided level and color settings.
func Setup(level LogLevel, useColor bool) {
	cfg.level = level
	cfg.useColor = useColor
	cfg.output = os.Stderr
}

// SetOutput sets the output of the logger to go to the specified writer.
func SetOutput(output io.Writer) {
	cfg.output = output
}

// SetupFromEnv sets up the logger based on the LOG_LEVEL and LOG_COLOR
// environment variables.
func SetupFromEnv() {
	l := DEBUG

	ln := strings.ToUpper(os.Getenv("LOG_LEVEL"))
	c := strings.ToUpper(os.Getenv("LOG_COLOR")) == "TRUE"

	for idx, name := range levelNames {
		if name == ln {
			l = DEBUG + LogLevel(idx)
			break
		}
	}

	Setup(l, c)
}

// Log logs a message with the specified log level.
func Log(level LogLevel, msg string) {
	if level < cfg.level || level > PANIC {
		return
	}

	line := fmt.Sprint(time.Now().Format(time.RFC3339), " ", levelNames[level-DEBUG], " ", msg)

	if cfg.useColor {
		line = fmt.Sprintf("%s%s%s", colorCodes[level-DEBUG], line, noColor)
	}

	fmt.Fprintln(cfg.output, line)

	if level >= PANIC {
		panic(msg)
	}
}

// Logf logs a message with the specified log level. The function takes a
// format string and arguments and passes it through fmt.Sprintf() to get
// the message string.
func Logf(level LogLevel, f string, args ...interface{}) {
	msg := fmt.Sprintf(f, args...)
	Log(level, msg)
}

// Debug is a convenience function equivalent to Log(DEBUG, msg)
func Debug(msg string) {
	Log(DEBUG, msg)
}

// Info is a convenience function equivalent to Log(INFO, msg)
func Info(msg string) {
	Log(INFO, msg)
}

// Warning is a convenience function equivalent to Log(WARNING, msg)
func Warning(msg string) {
	Log(WARNING, msg)
}

// Error is a convenience function equivalent to Log(ERROR, msg)
func Error(msg string) {
	Log(ERROR, msg)
}

// Panic is a convenience function equivalent to Log(PANIC, msg)
func Panic(msg string) {
	Log(PANIC, msg)
}

// Debugf is a convenience function equivalent to Logf(DEBUG, fmt, args...)
func Debugf(f string, args ...interface{}) {
	Logf(DEBUG, f, args...)
}

// Infof is a convenience function equivalent to Logf(INFO, fmt, args...)
func Infof(f string, args ...interface{}) {
	Logf(INFO, f, args...)
}

// Warningf is a convenience function equivalent to Logf(WARNING, fmt, args...)
func Warningf(f string, args ...interface{}) {
	Logf(WARNING, f, args...)
}

// Errorf is a convenience function equivalent to Logf(ERROR, fmt, args...)
func Errorf(f string, args ...interface{}) {
	Logf(ERROR, f, args...)
}

// Panicf is a convenience function equivalent to Logf(PANIC, fmt, args...)
func Panicf(f string, args ...interface{}) {
	Logf(PANIC, f, args...)
}
