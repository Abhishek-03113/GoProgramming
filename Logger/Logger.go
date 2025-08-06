package main

import "fmt"

type Logger interface {
	INFO(msg string) string
	WARN(msg string) string
	ERROR(msg string) string
}

type LogFormatter interface {
	Format(level string, message string) string
}

type Process struct {
	PID        int
	Status     string
	MemoryUsed float64
}
type ProcessLogger struct {
	Process
}

func (p ProcessLogger) INFO(msg string) string {
	return fmt.Sprintf("Process Id : %d Process Status : %s Process MemoryUsage : %.2f ", p.PID, p.Status, p.MemoryUsed)
}

func (p ProcessLogger) WARN(msg string) string {
	return fmt.Sprintf("The process is consuming high memory PID : %d, Memory Used : %.2f", p.PID, p.MemoryUsed)
}

func (p ProcessLogger) ERROR(msg string) string {
	return fmt.Sprintf("An Error has occured during runtime of process %d", p.PID)
}

type MemoryLogger struct {
	Process
}

func (m MemoryLogger) INFO(msg string) string {
	return fmt.Sprintf("The process: %d, Memory Used : %.2f", m.PID, m.MemoryUsed)

}

func (m MemoryLogger) WARN(msg string) string {
	return fmt.Sprintf("The process is consuming high memory PID : %d, Memory Used : %.2f", m.PID, m.MemoryUsed)

}

func (m MemoryLogger) ERROR(msg string) string {
	return fmt.Sprintf("Out of system memory :_)")
}

type PlainFormatter struct {
}

func (p PlainFormatter) Format(level string, message string) string {
	return fmt.Sprintf("Log Level : %s, Logger Message : %s", level, message)
}

type JSONFormatter struct {
}

func (J JSONFormatter) Format(level string, message string) string {
	return fmt.Sprintf("{Level : %s, Message : {%s}}", level, message)
}
