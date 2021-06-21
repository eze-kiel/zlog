package zlog

const (
	debugPrefix = "DEBUG"
	infoPrefix  = "INFO"
	warnPrefix  = "WARNING"
	errorPrefix = "ERROR"
	fatalPrefix = "FATAL"
)

func (l *Logger) SetDebugPrefix(prefix string) {
	l.debugPrefix = prefix
}

func (l *Logger) SetInfoPrefix(prefix string) {
	l.infoPrefix = prefix
}

func (l *Logger) SetWarnPrefix(prefix string) {
	l.warnPrefix = prefix
}

func (l *Logger) SetErrorPrefix(prefix string) {
	l.errorPrefix = prefix
}

func (l *Logger) SetFatalPrefix(prefix string) {
	l.fatalPrefix = prefix
}

func (l *Logger) GetDebugPrefix() string {
	return l.debugPrefix
}

func (l *Logger) GetInfoPrefix() string {
	return l.infoPrefix
}

func (l *Logger) GetWarnPrefix() string {
	return l.warnPrefix
}

func (l *Logger) GetErrorPrefix() string {
	return l.errorPrefix
}

func (l *Logger) GetFatalPrefix() string {
	return l.fatalPrefix
}

func (l *Logger) SetPrefixBorders(first, last string) {
	l.prefixBorders = [2]string{first, last}
}
