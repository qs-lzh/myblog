package logger

import (
	"log"
	"net/http"
	"os"
)

type Logger struct {
	InfoLog *log.Logger
	ErrLog  *log.Logger
}

func NewLogger() *Logger {
	return &Logger{
		InfoLog: log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime),
		ErrLog:  log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile),
	}
}

func (l *Logger) Info(msg string) {
	l.InfoLog.Println(msg)
}

func (l *Logger) LogRequest(r *http.Request) {
	l.InfoLog.Printf("Request: \"%s %s\" from %s\n", r.Method, r.URL, r.RemoteAddr)
}

func (l *Logger) LogServerStart(port string) {
	l.InfoLog.Printf("Server started on part %s\n", port)
}

func (l *Logger) LogPageRender(page string) {
	l.InfoLog.Printf("Render: page \"%s\" rendered successfully\n", page)
}

func (l *Logger) LogDBModify(action, table string) {
	l.InfoLog.Printf("DB: %s action on \"%s\"\n", action, table)
}

func (l *Logger) Error(msg string) {
	l.ErrLog.Println(msg)
}
