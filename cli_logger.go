package log

import (
	"fmt"
	"io"
)

type CliLogger struct {
}

func CreateCLILogger() io.Writer {
	return CliLogger{}
}

func (logger CliLogger) Write(msg []byte) (int, error) {
	return fmt.Println(string(msg))
}
