package loggger

import (
	"log"
	"os"
)

type Logger interface {
	Debug(...interface{})
	Error(...interface{})
	Fatal(...interface{})
	Info(...interface{})
	Panic(...interface{})
	Warn(...interface{})
}

type logger struct {
	*log.Logger
}

const (
	ErrMsgEmptyFormat   = "empty log format"
	ErrMsgInvalidFormat = "invalid log format"
	ErrMsgNilMsg        = "log message is nil"

	LevelDebug = iota
	LevelError
	LevelFatal
	LevelInfo
	LevelPanic
	LevelWarn
)

func New(loggers ...*log.Logger) Logger {
	// todo: support multiple loggers (e.g. zap)
	if len(loggers) > 0 {
		return &logger{
			loggers[0],
		}
	}

	return &logger{
		log.Default(),
	}
}

func (l *logger) Debug(v ...interface{}) {
	l.print(LevelDebug, v)
}

func (l *logger) Error(v ...interface{}) {
	l.print(LevelError, v)
}

func (l *logger) Fatal(v ...interface{}) {
	l.print(LevelFatal, v)
	os.Exit(1)
}

func (l *logger) Info(v ...interface{}) {
	l.print(LevelInfo, v)
}

func (l *logger) Panic(v ...interface{}) {
	l.print(LevelPanic, v)
	panic(v)
}

func (l *logger) Warn(v ...interface{}) {
	l.print(LevelWarn, v)
}

func (l *logger) print(level uint8, parts ...interface{}) {
	var format string
	var ok bool

	// todo: replace Printf (too expensive) and add coloring
	switch len(parts) {
	case 0:
		l.Printf("%d %v", LevelWarn, ErrMsgNilMsg)
	case 1:
		format, ok = parts[0].(string)
		if !ok {
			l.Printf("%d %v", LevelWarn, ErrMsgInvalidFormat)
		} else if len(format) == 0 {
			l.Printf("%d %v", LevelWarn, ErrMsgEmptyFormat)
		} else {
			l.Printf("%s %v", l.level(level), parts[0])
		}
	default:
		format, ok = parts[0].(string)
		if !ok {
			l.Printf("%s %v", l.level(level), ErrMsgInvalidFormat)
			return
		} else {
			// todo: for each part, add %v
			l.Printf("%s %s %v\n", l.level(level), format, parts[1:])
		}
	}
}

func (l *logger) level(level uint8) string {
	switch level {
	case LevelDebug:
		return "[DEBUG]"
	case LevelError:
		return "[ERROR]"
	case LevelFatal:
		return "[FATAL]"
	case LevelPanic:
		return "[PANIC]"
	case LevelWarn:
		return "[WARN]"
	default:
		return "[INFO]"
	}
}
