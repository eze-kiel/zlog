package zlog

const (
	debugPrefix = "DEBUG"
	infoPrefix  = "INFO"
	warnPrefix  = "WARNING"
	errorPrefix = "ERROR"
	fatalPrefix = "FATAL"
)

// SetDebugPrefix sets the debug prefix string to prefix
func (l *Logger) SetDebugPrefix(prefix string) {
	l.debugPrefix = prefix
}

// SetInfoPrefix sets the info prefix string to prefix
func (l *Logger) SetInfoPrefix(prefix string) {
	l.infoPrefix = prefix
}

// SetWarnPrefix sets the warn prefix string to prefix
func (l *Logger) SetWarnPrefix(prefix string) {
	l.warnPrefix = prefix
}

// SetErrorPrefix sets the error prefix string to prefix
func (l *Logger) SetErrorPrefix(prefix string) {
	l.errorPrefix = prefix
}

// SetFatalPrefix sets the fatal prefix string to prefix
func (l *Logger) SetFatalPrefix(prefix string) {
	l.fatalPrefix = prefix
}

// GetDebugPrefix returns the debug prefix string
func (l *Logger) GetDebugPrefix() string {
	return l.debugPrefix
}

// GetInfoPrefix returns the info prefix string
func (l *Logger) GetInfoPrefix() string {
	return l.infoPrefix
}

// GetWarnPrefix returns the warn prefix string
func (l *Logger) GetWarnPrefix() string {
	return l.warnPrefix
}

// GetErrorPrefix returns the error prefix string
func (l *Logger) GetErrorPrefix() string {
	return l.errorPrefix
}

// GetFatalPrefix returns the fatal prefix string
func (l *Logger) GetFatalPrefix() string {
	return l.fatalPrefix
}

// SetPrefixBorders allows to set the level prefix borders
func (l *Logger) SetPrefixBorders(first, last string) {
	l.prefixBorders = [2]string{first, last}
}
