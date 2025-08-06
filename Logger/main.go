package main

type Logger interface {
	INFO(msg string) string
	WARN(msg string) string
	ERROR(msg string) string
}

type LogFormatter interface {
	Format(level string, message string) string
}
