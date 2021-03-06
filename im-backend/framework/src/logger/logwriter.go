package logger

type LogWriter interface {
	Open(logPath string) error
	WriteString(s string) (n int, err error)
	Write([]byte) (n int, err error)
	Close()
	Flush()
}
