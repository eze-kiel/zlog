package zlog

const (
	timeFormat = "15:04:05"
)

// SetTimeFormat allows an user to change default time format
func (l *Logger) SetTimeFormat(format string) {
	l.timeFormat = format
}

// SetTimeBorders allows an user to set the time prefix borders
func (l *Logger) SetTimeBorders(first, last string) {
	l.timeBorders = [2]string{first, last}
}
