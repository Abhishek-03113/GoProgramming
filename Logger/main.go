package main

import (
	"Logger/LogWriter" // Adjust the import path as per your project folder structure
	"fmt"
)

func main() {
	// Setup ProcessLogger and MemoryLogger
	processLogger := ProcessLogger{
		Process: Process{
			PID:        123,
			Status:     "Running",
			MemoryUsed: 342.78,
		},
	}

	memoryLogger := MemoryLogger{
		Process: Process{
			PID:        456,
			Status:     "Running",
			MemoryUsed: 892.44,
		},
	}

	// Formatters
	plainFormatter := PlainFormatter{}
	jsonFormatter := JSONFormatter{}

	// Writers
	consoleWriter := LogWriter.ConsoleLogWriter{}
	fileWriter := LogWriter.FileLogWriter{}

	// Log messages using ProcessLogger
	log1 := processLogger.INFO("Service started")
	log2 := processLogger.WARN("Memory usage high")
	log3 := processLogger.ERROR("Crash occurred")

	// Format logs using PlainFormatter
	formatted1 := plainFormatter.Format("INFO", log1)
	formatted2 := plainFormatter.Format("WARN", log2)
	formatted3 := plainFormatter.Format("ERROR", log3)

	// Write logs to console
	_ = consoleWriter.WriteLog(formatted1)
	_ = consoleWriter.WriteLog(formatted2)
	_ = consoleWriter.WriteLog(formatted3)

	// Write logs to file using JSONFormatter
	log4 := memoryLogger.INFO("Another process started")
	log5 := memoryLogger.WARN("Critical memory alert")
	log6 := memoryLogger.ERROR("Kernel panic")

	formatted4 := jsonFormatter.Format("INFO", log4)
	formatted5 := jsonFormatter.Format("WARN", log5)
	formatted6 := jsonFormatter.Format("ERROR", log6)

	_ = fileWriter.WriteLog(formatted4)
	_ = fileWriter.WriteLog(formatted5)
	_ = fileWriter.WriteLog(formatted6)
	_ = fileWriter.WriteLog(formatted6)
	_ = fileWriter.WriteLog(formatted6)
	_ = fileWriter.WriteLog(formatted6)

	fmt.Println("Logs written using ProcessLogger and MemoryLogger.")
}
