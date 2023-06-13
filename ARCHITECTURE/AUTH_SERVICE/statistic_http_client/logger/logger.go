package logger

import (
	"fmt"
	"github.com/fatih/color"
)

// STATLogger - logger for statistic_service
type STATLogger struct {
	Counter int
}

// New - returns an instance of MEMOLogger
func New() *STATLogger {
	return &STATLogger{}
}

func (l *STATLogger) Log(msg string) {
	l.Counter++
	d := color.New(color.FgGreen, color.Bold)
	d.Printf("STAT")
	fmt.Printf("[%04d]", l.Counter)
	fmt.Printf(" %s \n", msg)
}
