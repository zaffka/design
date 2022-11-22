package log

import (
	"fmt"
	"log"
)

const (
	Info = iota
	Error
)

var levels = [...]string{
	Info:  "[Info]: ",
	Error: "[Error]: ",
}

func Errorf(format string, v ...any) {
	log.Println(levels[Error] + fmt.Sprintf(format, v...))
}

func Infof(format string, v ...any) {
	log.Println(levels[Info] + fmt.Sprintf(format, v...))
}
