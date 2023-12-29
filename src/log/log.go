package log

import (
	"fmt"
	"strings"
)

const (
	Silent = iota
	Warning
	Error
	Info
	Debug
)

var Level byte
var levelString string

func Init(logLevel string) {
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
		Level = Silent
		levelString = "silent"
	}
}

func Infoln(format string, a ...any) {
	Ln(Info, format, a...)
}

func WarnLn(format string, a ...any) {
	Ln(Warning, format, a...)
}

func Errorln(format string, a ...any) { // Error {
	Ln(Error, format, a...)
	//return fmt.Errorf(format, a...)
}

func Debugln(format string, a ...any) {
	Ln(Debug, format, a...)
}

func Ln(logLevel byte, format string, a ...any) {
	// 根据 level 选择性显示内容
	//if logLevel > Level {
	//	return
	//}
	msg := fmt.Sprintln("[", levelString, "] ", fmt.Sprintf(format, a...))
	fmt.Print(msg)
}
