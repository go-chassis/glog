package glog

import (
	"log"
	"os"
)

var LogLevel Level = 3

const fatalExitCode = 255
const activeExitCode = 1

type Level int32
type Verbose bool

func Flush() {
}

func V(level Level) Verbose {
	return LogLevel >= level
}

func (v Verbose) Info(args ...interface{}) {
	if v {
		log.Print(args...)
	}
}

func (v Verbose) Infoln(args ...interface{}) {
	if v {
		log.Println(args...)
	}
}

func (v Verbose) Infof(format string, args ...interface{}) {
	if v {
		log.Printf(format, args...)
	}
}

func Info(args ...interface{}) {
	log.Print(args...)
}

func InfoDepth(depth int, args ...interface{}) {
	log.Print(args...)
}

func Infoln(args ...interface{}) {
	log.Println(args...)
}

func Infof(format string, args ...interface{}) {
	log.Printf(format, args...)
}

func Warning(args ...interface{}) {
	log.Print(args...)
}

func WarningDepth(depth int, args ...interface{}) {
	log.Print(args...)
}

func Warningln(args ...interface{}) {
	log.Println(args...)
}

func Warningf(format string, args ...interface{}) {
	log.Printf(format, args...)
}

func Error(args ...interface{}) {
	log.Print(args...)
}

func ErrorDepth(depth int, args ...interface{}) {
	log.Print(args...)
}

func Errorln(args ...interface{}) {
	log.Println(args...)
}

func Errorf(format string, args ...interface{}) {
	log.Printf(format, args...)
}

func Fatal(args ...interface{}) {
	log.Print(args...)
	os.Exit(fatalExitCode)
}

func FatalDepth(depth int, args ...interface{}) {
	log.Print(args...)
	os.Exit(fatalExitCode)
}

func Fatalln(args ...interface{}) {
	log.Println(args...)
	os.Exit(fatalExitCode)
}

func Fatalf(format string, args ...interface{}) {
	log.Printf(format, args...)
	os.Exit(fatalExitCode)
}

func Exit(args ...interface{}) {
	log.Print(args...)
	os.Exit(activeExitCode)
}

func ExitDepth(depth int, args ...interface{}) {
	log.Print(args...)
	os.Exit(activeExitCode)
}

func Exitln(args ...interface{}) {
	log.Println(args...)
	os.Exit(activeExitCode)
}

func Exitf(format string, args ...interface{}) {
	log.Printf(format, args...)
	os.Exit(activeExitCode)
}
