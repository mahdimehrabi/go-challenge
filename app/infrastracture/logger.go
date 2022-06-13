package infrastracture

import (
	"log"
	"os"
)

type SegmentLogger struct {
	LG *log.Logger
}

func (l *SegmentLogger) Error(err string) {
	l.LG.Print(err)
}

func NewLogger() *SegmentLogger {
	lg := log.New(os.Stdout, "segment ", log.LstdFlags)
	return &SegmentLogger{LG: lg}
}
