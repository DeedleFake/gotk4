// Package logger provides abstractions over girgen's logging.
package logger

import (
	"log"

	"github.com/fatih/color"
)

// Level vaguely describes the importance of each log entry.
type Level uint8

const (
	// Debug describes logging information that wouldn't be useful for anything
	// except for during debugging.
	Debug Level = iota
	// Skip describes an event that is only used when something is skipped.  The
	// reason for skipping should not be written here and should instead be in
	// Debug.
	Skip
	// Unusual describes an event that is unusual but is not fatal.
	Unusual
	// Error describes an event that is erroneous and will likely not produce a
	// valid output.
	Error
)

func (lvl Level) prefix() string {
	switch lvl {
	case Debug:
		return "debug:"
	case Skip:
		return "skipped:"
	case Unusual:
		return "unusuality:"
	case Error:
		return "error:"
	default:
		return ""
	}
}

func (lvl Level) colorf(f string, v ...any) string {
	switch lvl {
	case Skip:
		return color.BlueString(f, v...)
	case Unusual:
		return color.YellowString(f, v...)
	case Error:
		return color.New(color.Bold, color.FgHiRed).Sprintf(f, v...)
	case Debug:
		fallthrough
	default:
		return color.New(color.Faint).Sprintf(f, v...)
	}
}

// LineLogger describes anything that can log itself.
type LineLogger interface {
	Logln(Level, ...any)
}

// NoopLogger is a logger that doesn't log anything. It is used as a placeholder
// to pass into functions that need logging.
var NoopLogger = noopLogger{}

type noopLogger struct{}

func (noop noopLogger) Logln(Level, ...any) {}

// Prefix prepends the given prefixes into the given value list.
func Prefix(list []any, p any) []any {
	list = append(list, nil)
	copy(list[1:], list)
	list[0] = p
	return list
}

// Stdlog renders the given log entry into the stdlib logger.
func Stdlog(logger *log.Logger, minlevel, level Level, v ...any) {
	if logger == nil || minlevel > level {
		return
	}

	prefix := level.prefix()
	if prefix != "" {
		prefix = level.colorf(prefix)
		v = Prefix(v, prefix)
	}

	logger.Println(v...)
}
