package LogWriter

import (
	"fmt"
	"os"
)

type LogWriter interface {
	WriteLog(log string) error
}

type ConsoleLogWriter struct {
}

func (c ConsoleLogWriter) WriteLog(log string) error {
	_, err := fmt.Println(log)

	if err != nil {
		return err
	}
	return nil
}

type FileLogWriter struct {
}

func (f FileLogWriter) WriteLog(log string) error {
	fileName := "app.log"
	file, err := os.OpenFile("example.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	data := []byte(log)

	if err == nil {
		file.Write(data)

		fmt.Println("Log has been written to file ", fileName)
		return nil
	}
	fmt.Println("Failed writing to file ", err)
	return err
}
