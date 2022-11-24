package logging

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"io"
	"os"
	"path"
	"runtime"
)

// kafka -- info, debug
// file -- error, trace
// stdout -- warning critical
type writerHook struct {
	Writer    []io.Writer
	LogLevels []logrus.Level
}

// Levels - returns the levels
func (hook *writerHook) Levels() []logrus.Level {
	return hook.LogLevels
}

// Fire - calls every time when we log something
func (hook *writerHook) Fire(entry *logrus.Entry) error {
	line, err := entry.String()
	if err != nil {
		return err
	}
	for _, writer := range hook.Writer {
		_, err := fmt.Fprint(writer, line)
		if err != nil {
			return err
		}
	}
	return nil
}

var e *logrus.Entry

type Logger struct {
	*logrus.Entry
}

func GetLogger() *Logger {
	return &Logger{e}
}

func (l *Logger) GetLoggerWithField(k string, v interface{}) *Logger {
	return &Logger{l.WithField(k, v)}
}

func Init() {
	l := logrus.New()
	l.SetReportCaller(true)
	l.SetFormatter(&logrus.TextFormatter{
		FullTimestamp: true,
		CallerPrettyfier: func(frame *runtime.Frame) (function string, file string) {
			filename := path.Base(frame.File)
			return fmt.Sprintf("%s()", frame.Function), fmt.Sprintf("%s:%d", filename, frame.Line)
		},
		DisableColors: true,
	})

	err := os.Mkdir("logs", 0644)
	if err != nil && !os.IsExist(err) {
		panic(err)
	}

	if _, err := os.Stat("logs/all.log"); err != nil {
		_, err := os.Create("logs/all.log")
		if err != nil {
			panic(err)
		}
	}

	file, err := os.OpenFile("logs/all.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		panic(err)
	}

	l.SetOutput(io.Discard) // why? we need to write to file and console

	l.AddHook(&writerHook{
		Writer:    []io.Writer{file, os.Stdout},
		LogLevels: logrus.AllLevels,
	})

	l.SetLevel(logrus.TraceLevel)

	e = logrus.NewEntry(l)
}
