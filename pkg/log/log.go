package log

var (
	debugLog Logger = &DebugLogger{}
	infoLog  Logger = &InfoLogger{}
	warnLog  Logger = &WarningLogger{}
	errLog   Logger = &ErrorLogger{}
)

type Logger interface {
	logOut(format *string, args ...interface{})
}

func Debug(args ...interface{}) {
	debugLog.logOut(nil, args...)
}

func Debugf(format string, args ...interface{}) {
	debugLog.logOut(&format, args...)
}

func Info(args ...interface{}) {
	infoLog.logOut(nil, args...)
}

func Infof(format string, args ...interface{}) {
	infoLog.logOut(&format, args...)
}

func Warning(args ...interface{}) {
	warnLog.logOut(nil, args...)
}

func Warningf(format string, args ...interface{}) {
	warnLog.logOut(&format, args...)
}

func Error(args ...interface{}) {
	errLog.logOut(nil, args...)
}

func Errorf(format string, args ...interface{}) {
	errLog.logOut(&format, args...)
}
