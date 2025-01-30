package logger

import (
	"fmt"
	"log"
	"os"
	"time"
)

type Logger interface {
	Info(v ...any)
	Warn(v ...any)
	Error(v ...any)
	Fatal(v ...any)
}

type logger struct {
	log *log.Logger
}

func New() Logger {
	prefix := fmt.Sprintf("[%v]", time.UTC)
	flag := log.Flags() | log.LUTC
	l := log.New(os.Stdout, prefix, flag)

	return &logger{
		log: l,
	}
}

func (l *logger) Info(v ...any) {
	l.log.Println("[INFO]:", fmt.Sprint(v...))
}

func (l *logger) Warn(v ...any) {
	l.log.Println("[WARN]:", fmt.Sprint(v...))
}

func (l *logger) Error(v ...any) {
	l.log.Println("[ERROR]:", fmt.Sprint(v...))
}

func (l *logger) Fatal(v ...any) {
	l.log.Fatal("[FATAL]:", fmt.Sprint(v...))
}
