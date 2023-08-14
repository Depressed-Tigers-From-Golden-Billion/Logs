package log

import (
	"io"
	"os"
)

type FileLogger struct {
	file *os.File
}

func CreateFileLogger(path string) io.Writer {
	var logger FileLogger

	file, err := os.OpenFile(path, os.O_CREATE|os.O_APPEND, 0640) // Read/Write owner + read group
	if nil != err {
		return nil
	}

	logger.file = file

	return logger
}

func (logger FileLogger) Write(msg []byte) (int, error) {
	return logger.file.Write(msg)
}
