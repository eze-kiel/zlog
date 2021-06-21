package zlog

import (
	"os"

	"github.com/fatih/color"
)

// Debug logs a message at the debug level
func (l *Logger) Debug(message string) {
	if l.level <= DebugLvl {
		if l.useColors {
			l.log(color.BlueString(l.debugPrefix), message)
		} else {
			l.log(l.debugPrefix, message)
		}
	}
}

// Info logs a message at the info level
func (l *Logger) Info(message string) {
	if l.level <= InfoLvl {
		if l.useColors {
			l.log(color.GreenString(l.infoPrefix), message)
		} else {
			l.log(l.infoPrefix, message)
		}
	}
}

// Warn logs a message at the warn level
func (l *Logger) Warn(message string) {
	if l.level <= WarnLvl {
		if l.useColors {
			l.log(color.YellowString(l.warnPrefix), message)
		} else {
			l.log(l.warnPrefix, message)
		}
	}
}

// Error logs a message at the error level
func (l *Logger) Error(message string) {
	if l.level <= ErrorLvl {
		if l.useColors {
			l.log(color.RedString(l.errorPrefix), message)
		} else {
			l.log(l.errorPrefix, message)
		}
	}
}

// Fatal logs a message at the fatal level
func (l *Logger) Fatal(message string) {
	if l.level <= FatalLvl {
		if l.useColors {
			l.log(color.RedString(l.fatalPrefix), message)
		} else {
			l.log(l.fatalPrefix, message)
		}
		os.Exit(1)
	}
}
