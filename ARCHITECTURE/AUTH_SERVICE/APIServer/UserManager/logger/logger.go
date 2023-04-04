package logger

import (
	"fmt"
	"github.com/fatih/color"
)

// MEMOLogger - logger for buffer
type MEMOLogger struct {
	Counter int
}

// New - returns an instance of MEMOLogger
func New() *MEMOLogger {
	return &MEMOLogger{}
}

func (l *MEMOLogger) CacheLog(msg string) {
	l.Counter++
	d := color.New(color.FgGreen, color.Bold)
	d.Printf("MEMO")
	fmt.Printf("[%04d] ", l.Counter)
	d1 := color.New(color.FgHiMagenta, color.Bold)
	d1.Printf("CCHE")
	fmt.Printf(" %s \n", msg)
}

func (l *MEMOLogger) DatabaseLog(msg string) {
	l.Counter++
	d := color.New(color.FgGreen, color.Bold)
	d.Printf("MEMO")
	fmt.Printf("[%04d] ", l.Counter)
	d1 := color.New(color.FgHiRed, color.Bold)
	d1.Printf("TODB")
	fmt.Printf(" %s \n", msg)
}

func (l *MEMOLogger) ManagerLog(msg string) {
	d := color.New(color.FgHiGreen, color.Bold)
	d.Printf("MANA")
	fmt.Printf("[0000] ")
	fmt.Printf(" %s \n", msg)
}
