package logger

import (
	"fmt"
	"github.com/fatih/color"
)

type DBLogger struct {
	counter int
}

func New() *DBLogger {
	return &DBLogger{counter: 0}
}

func (l *DBLogger) Log(message string) {
	l.counter++
	d := color.New(color.FgHiYellow, color.Bold)
	d.Printf("BASE")
	fmt.Printf("[%04d] ", l.counter)
	fmt.Println(message)
}

func (l *DBLogger) Error(message string) {
	l.counter++
	d := color.New(color.FgRed, color.Bold)
	d.Printf("BASE")
	fmt.Printf("[%04d] ", l.counter)
	fmt.Println(message)
}
