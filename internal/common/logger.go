package common

import "fmt"

type LogLevel int

const (
	DEBUG LogLevel = LogLevel(iota)
	INFO  LogLevel = LogLevel(iota)
	WARN  LogLevel = LogLevel(iota)
	ERROR LogLevel = LogLevel(iota)
)

type Logger interface {
	Debug(msg string, args ...any)
	Info(msg string, args ...any)
	Warn(msg string, args ...any)
	Error(msg string, args ...any)
}

type StdoutLogger struct {
	level LogLevel
}

func NewStdoutLogger(level LogLevel) *StdoutLogger {
	return &StdoutLogger{
		level: level,
	}
}

func (l *StdoutLogger) Debug(msg string, args ...any) {
	if l.level > DEBUG {
		return
	}

	fmt.Printf(msg, args...)
}

func (l *StdoutLogger) Info(msg string, args ...any) {
	if l.level > INFO {
		return
	}

	fmt.Printf(msg, args...)
}

func (l *StdoutLogger) Warn(msg string, args ...any) {
	if l.level > WARN {
		return
	}

	fmt.Printf(msg, args...)
}

func (l *StdoutLogger) Error(msg string, args ...any) {
	if l.level > ERROR {
		return
	}

	fmt.Printf(msg, args...)
}
