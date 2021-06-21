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
	timeFormat = "[15:04:05]"
)

type Logger struct {
	mutex     *sync.Mutex
	level     int
	buffer    []byte
	fd        io.Writer
	logTime   bool
	useColors bool

	debugPrefix string
	infoPrefix  string
	warnPrefix  string
	errorPrefix string
	fatalPrefix string
}

func NewLogger() *Logger {
	return &Logger{
		mutex:     &sync.Mutex{},
		buffer:    []byte{},
		level:     InfoLvl,
		fd:        os.Stderr,
		logTime:   true,
		useColors: true,

		debugPrefix: debugPrefix,
		infoPrefix:  infoPrefix,
		warnPrefix:  warnPrefix,
		errorPrefix: errorPrefix,
		fatalPrefix: fatalPrefix,
	}
}

func (l *Logger) ParseLevel(level string) error {
	switch level {
	case "debug":
		l.level = DebugLvl
	case "info":
		l.level = InfoLvl
	case "warn":
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

func (l Logger) GetLogLevel() int {
	return l.level
}

func (l *Logger) log(pfx, msg string) {
	l.mutex.Lock()
	defer l.mutex.Unlock()

	l.buffer = []byte{}

	if l.logTime {
		l.buffer = append(l.buffer, []byte(time.Now().Format(timeFormat))...)
		l.buffer = append(l.buffer, ' ')
	}

	l.buffer = append(l.buffer, []byte(pfx)...)
	l.buffer = append(l.buffer, ' ')

	l.buffer = append(l.buffer, []byte(msg)...)

	fmt.Fprintf(l.fd, "%s\n", l.buffer)
}
