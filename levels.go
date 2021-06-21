package zlog

import (
	"os"

	"github.com/fatih/color"
)

func (l *Logger) Debug(message string) {
	if l.level <= DebugLvl {
		if l.useColors {
			l.log(color.BlueString(l.debugPrefix), message)
		} else {
			l.log(l.debugPrefix, message)
		}
	}
}

func (l *Logger) Info(message string) {
	if l.level <= InfoLvl {
		if l.useColors {
			l.log(color.GreenString(l.infoPrefix), message)
		} else {
			l.log(l.infoPrefix, message)
		}
	}
}

func (l *Logger) Warn(message string) {
	if l.level <= WarnLvl {
		if l.useColors {
			l.log(color.YellowString(l.warnPrefix), message)
		} else {
			l.log(l.warnPrefix, message)
		}
	}
}

func (l *Logger) Error(message string) {
	if l.level <= ErrorLvl {
		if l.useColors {
			l.log(color.RedString(l.errorPrefix), message)
		} else {
			l.log(l.errorPrefix, message)
		}
	}
}

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
