package LogWriter

type LogWriter interface {
	Write(log string) error
}
