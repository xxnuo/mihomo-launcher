package log

import (
	"fmt"
	"strings"
)

var (
	Level       = DefaultLoglevelE
	levelString = DefaultLoglevel
)

func SetLevelE(logLevel int) {
	Level = logLevel
}

func SetLevelS(logLevel string) {
	levelString = strings.ToLower(logLevel)
	switch levelString {
	case "silent":
		Level = Silent
	case "warning":
		Level = Warning
	case "error":
		Level = Error
	case "info":
		Level = Info
	case "debug":
		Level = Debug
	default:
		Level = DefaultLoglevelE
		levelString = DefaultLoglevel
	}
}

func Infoln(format string, a ...any) {
	Ln(Info, format, a...)
}

func Warnln(format string, a ...any) {
	Ln(Warning, format, a...)
}

func Errorln(format string, a ...any) { // Error {
	Ln(Error, format, a...)
}

func Debugln(format string, a ...any) {
	Ln(Debug, format, a...)
}

func Ln(logLevel int, format string, a ...any) {
	// 根据 level 选择性显示内容
	if logLevel > Level {
		return
	}
	msg := fmt.Sprintln("[", levelString, "] ", fmt.Sprintf(format, a...))
	fmt.Print(msg)
}
