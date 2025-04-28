package golog

import (
	"fmt"
	"os"
	"strings"
	"time"
)

type Level string

const (
	LevelInfo  Level = "INFO"
	LevelWarn  Level = "WARN"
	LevelDebug Level = "DEBUG"
	LevelError Level = "ERROR"
)

var levelColors = map[Level]string{
	LevelInfo:  "\033[32m",
	LevelWarn:  "\033[33m",
	LevelDebug: "\033[36m",
	LevelError: "\033[31m",
}

const resetColor = "\033[0m"

var logFile *os.File

func Init(path string) error {
	var err error
	logFile, err = os.OpenFile(path, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	return err
}

func Close() {
	if logFile != nil {
		logFile.Close()
	}
}

func log(level Level, format string, args ...interface{}) {
	if logFile == nil {
		panic("[glog] Logger not initialized. Call golog.Init(path) first.")
	}

	placeholderCount := strings.Count(format, "{}")
	if placeholderCount != len(args) {
		panic(fmt.Sprintf("[golog] Mismatch: found %d `{}` placeholders but got %d arguments", placeholderCount, len(args)))
	}

	builder := &strings.Builder{}
	argIndex := 0

	for {
		i := strings.Index(format, "{}")
		if i == -1 {
			builder.WriteString(format)
			break
		}
		builder.WriteString(format[:i])
		builder.WriteString(fmt.Sprintf("%v", args[argIndex]))
		format = format[i+2:]
		argIndex++
	}

	message := builder.String()
	timestamp := time.Now().Format("2006-01-02 15:04:05")

	color := levelColors[level]
	consoleLine := fmt.Sprintf("%s [%s%s%s] %s", timestamp, color, level, resetColor, message)
	fmt.Println(consoleLine)

	fileLine := fmt.Sprintf("%s [%s] %s\n", timestamp, level, message)
	_, _ = logFile.WriteString(fileLine)
}

func Info(format string, args ...interface{}) {
	log(LevelInfo, format, args...)
}

func Warn(format string, args ...interface{}) {
	log(LevelWarn, format, args...)
}

func Debug(format string, args ...interface{}) {
	log(LevelDebug, format, args...)
}

func Error(format string, args ...interface{}) {
	log(LevelError, format, args...)
}
