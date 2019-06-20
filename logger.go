package golog

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"strings"
	"sync"
	"time"
)

const (
	LevelDebug = iota
	LevelInfo
	LevelNotice
	LevelWarning
	LevelError
	LevelPanic
)

var levelMessage = map[int]string{
	LevelDebug:   "debug",
	LevelInfo:    "info",
	LevelNotice:  "notice",
	LevelWarning: "warning",
	LevelError:   "error",
	LevelPanic:   "panic",
}

type Logger struct {
	output io.Writer
	mu     sync.Mutex
	level  int
	prefix string
}

func New(out io.Writer) *Logger {
	l := &Logger{
		output: out,
		level:  LevelDebug,
	}
	return l
}

var stdLog = New(os.Stderr)

func (l *Logger) Output(level int, s string) {
	if level < l.level {
		return
	}
	outMsg := l.formatOutput(level, s)
	l.output.Write(outMsg)
}

func (l *Logger) formatOutput(level int, s string) []byte {
	levelMsg, _ := levelMessage[level]
	var buf bytes.Buffer
	buf.WriteByte('[')
	buf.WriteString(time.Now().Format("2006-01-02 15:04:05"))
	buf.WriteString("] ")
	if l.prefix != "" {
		buf.WriteString(l.prefix)
		buf.WriteString(" ")
	}
	buf.WriteString(strings.ToUpper(levelMsg))
	buf.WriteString(": ")
	buf.WriteString(s)
	if len(s) == 0 || s[len(s)-1] != '\n' {
		buf.WriteByte('\n')
	}
	return buf.Bytes()
}

func (l *Logger) SetPrefix(prefix string) {
	l.prefix = prefix
}

func (l *Logger) SetOutput(w io.Writer) {
	l.mu.Lock()
	defer l.mu.Unlock()
	l.output = w
}

func (l *Logger) SetLevel(level int) {
	l.mu.Lock()
	defer l.mu.Unlock()
	l.level = level
}

func (l *Logger) GetOutput() io.Writer {
	return l.output
}

func (l *Logger) Debug(s string, v ...interface{}) {
	l.Output(LevelDebug, fmt.Sprintf(s, v...))
}

func (l *Logger) Info(s string, v ...interface{}) {
	l.Output(LevelInfo, fmt.Sprintf(s, v...))
}

func (l *Logger) Notice(s string, v ...interface{}) {
	l.Output(LevelNotice, fmt.Sprintf(s, v...))
}

func (l *Logger) Warning(s string, v ...interface{}) {
	l.Output(LevelWarning, fmt.Sprintf(s, v...))
}

func (l *Logger) Error(s string, v ...interface{}) {
	l.Output(LevelError, fmt.Sprintf(s, v...))
}

func (l *Logger) Panic(s string, v ...interface{}) {
	l.Output(LevelPanic, fmt.Sprintf(s, v...))
	panic(s)
}

func SetOutput(w io.Writer) {
	stdLog.SetOutput(w)
}

func GetOutput() io.Writer {
	return stdLog.GetOutput()
}

func SetPrefix(prefix string) {
	stdLog.SetPrefix(prefix)
}

func Debug(s string, v ...interface{}) {
	stdLog.Debug(s, v...)
}

func Info(s string, v ...interface{}) {
	stdLog.Info(s, v...)
}

func Notice(s string, v ...interface{}) {
	stdLog.Notice(s, v...)
}

func Warning(s string, v ...interface{}) {
	stdLog.Warning(s, v...)
}

func Error(s string, v ...interface{}) {
	stdLog.Error(s, v...)
}

func Panic(s string, v ...interface{}) {
	stdLog.Panic(s, v...)
}
