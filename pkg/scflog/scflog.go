package scflog

import "log"
import "os"

const (
	DEBUG = iota
	INFO
	ERROR
	NA
)

type Logger interface {
	Info(format string, v ...interface{})
	Error(format string, v ...interface{})
	Debug(format string, v ...interface{})
}

type Log struct {
	error *log.Logger
	info  *log.Logger
	debug *log.Logger
	level int
}

//Err writes messages to stderr
func (log *Log) Error(format string, v ...interface{}) {
	if log.level > ERROR {
		return
	}

	log.error.Printf(format, v...)
}

//Info writes messages to stdout
func (log *Log) Info(format string, v ...interface{}) {
	if log.level > INFO {
		return
	}

	log.info.Printf(format, v...)
}

//Debug writes messages to stdout
func (log *Log) Debug(format string, v ...interface{}) {
	if log.level > DEBUG {
		return
	}

	log.debug.Printf(format, v...)
}

//New creates new simple logger
func New(level int) *Log {
	log := &Log{
		error: log.New(os.Stdout, "ERROR: ", log.Ldate|log.Ltime|log.Lshortfile),
		info:  log.New(os.Stderr, "INFO: ", log.Ldate|log.Ltime|log.Lshortfile),
		debug: log.New(os.Stdout, "DEBUG: ", log.Ldate|log.Ltime|log.Lshortfile),
		level: level,
	}

	return log
}
