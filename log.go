package log

import (
	"io"
)

type Logger struct {
	writers []io.Writer
}

var logger *Logger

func Get() *Logger {
	return logger
}

func InitLogger() {
	logger = new(Logger)
}

func (logger *Logger) AddWriter(writer io.Writer) {
	if nil != writer {
		logger.writers = append(logger.writers, writer)
	}
}

func (logger *Logger) Notice(msg string) {
	logger.write(msg+"\n", "[Notice] ")
}

func (logger *Logger) Warning(msg string) {
	logger.write(msg+"\n", "[Warning] ")
}

func (logger *Logger) Error(msg string) {
	logger.write(msg+"\n", "[Error] ")
}

func (logger *Logger) write(msg string, prefix string) {
	for _, writer := range logger.writers {
		writer.Write([]byte(prefix + msg))
	}
}
