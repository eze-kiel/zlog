package zlog

import (
	"errors"
	"fmt"
	"io"
	"os"
	"sync"
	"time"
)

const (
	DebugLvl = iota
	InfoLvl
	WarnLvl
	ErrorLvl
	FatalLvl
)

const (
	timeFormat = "15:04:05"
)

type Logger struct {
	mutex         *sync.Mutex
	level         int
	buffer        []byte
	fd            io.Writer
	logTime       bool
	useColors     bool
	timeFormat    string
	prefixBorders [2]string
	timeBorders   [2]string

	debugPrefix string
	infoPrefix  string
	warnPrefix  string
	errorPrefix string
	fatalPrefix string
}

// NewLogger creates a new logger object with default values set
func NewLogger() *Logger {
	return &Logger{
		mutex:         &sync.Mutex{},
		buffer:        []byte{},
		level:         InfoLvl,
		fd:            os.Stderr,
		logTime:       true,
		useColors:     true,
		timeFormat:    timeFormat,
		prefixBorders: [2]string{"", ""},
		timeBorders:   [2]string{"", ""},

		debugPrefix: debugPrefix,
		infoPrefix:  infoPrefix,
		warnPrefix:  warnPrefix,
		errorPrefix: errorPrefix,
		fatalPrefix: fatalPrefix,
	}
}

// ParseLevel reads from a string an tries to set the logger's level according
// to the value provided. It returns an error if it can not parse the level
func (l *Logger) ParseLevel(level string) error {
	switch level {
	case "debug":
		l.level = DebugLvl
	case "info":
		l.level = InfoLvl
	case "warn", "warning":
		l.level = WarnLvl
	case "err", "error":
		l.level = ErrorLvl
	case "fatal":
		l.level = FatalLvl
	default:
		return errors.New("log level is not recognised")
	}
	return nil
}

// GetLogLevel returns the logger's current log level
func (l Logger) GetLogLevel() int {
	return l.level
}

// UseColors allows the user to create colored logs
func (l *Logger) UseColors(value bool) {
	l.useColors = value
}

// SetOutputFileDescriptor allows the user to choose the fd that will be used to
// log
func (l *Logger) SetOutputFileDescriptor(output io.Writer) {
	l.fd = output
}

// log writes a log line to the fd set
func (l *Logger) log(pfx, msg string) {
	l.mutex.Lock()
	defer l.mutex.Unlock()

	l.buffer = []byte{}

	if l.logTime {
		l.buffer = append(l.buffer, []byte(l.timeBorders[0])...)
		l.buffer = append(l.buffer, []byte(time.Now().Format(timeFormat))...)
		l.buffer = append(l.buffer, []byte(l.timeBorders[1])...)
		l.buffer = append(l.buffer, ' ')
	}

	l.buffer = append(l.buffer, []byte(l.prefixBorders[0])...)
	l.buffer = append(l.buffer, []byte(pfx)...)
	l.buffer = append(l.buffer, []byte(l.prefixBorders[1])...)
	l.buffer = append(l.buffer, ' ')

	l.buffer = append(l.buffer, []byte(msg)...)

	fmt.Fprintf(l.fd, "%s\n", l.buffer)
}
