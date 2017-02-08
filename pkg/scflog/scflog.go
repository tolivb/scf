package scflog

import "log"
import "os"
import "sync"

const (
	DEBUG = iota
	LOG
	ERROR
	NA
)

type Logger interface {
	Msg(format string, v ...interface{})
	Err(format string, v ...interface{})
	Dbg(format string, v ...interface{})
	Level(int)
}

type Log struct {
	error *log.Logger
	info  *log.Logger
	debug *log.Logger
	level int
	mutex *sync.Mutex
}

//Err writes messages to stderr
func (log *Log) Err(format string, v ...interface{}) {
	if log.level > ERROR {
		return
	}

	log.error.Printf(format, v...)
}

//Msg writes messages to stdout
func (log *Log) Msg(format string, v ...interface{}) {
	if log.level > DEBUG {
		return
	}

	log.info.Printf(format, v...)
}

//Msg writes messages to stdout
func (log *Log) Dbg(format string, v ...interface{}) {
	if log.level > DEBUG {
		return
	}

	log.debug.Printf(format, v...)
}

func (log *Log) Level(l int) {
	log.mutex.Lock()
	log.level = l
	log.mutex.Unlock()
}

//New creates new simple logger
func New(level int) *Log {
	log := &Log{
		error: log.New(os.Stdout, "ERROR: ", log.Ldate|log.Ltime|log.Lshortfile),
		info:  log.New(os.Stderr, "INFO: ", log.Ldate|log.Ltime|log.Lshortfile),
		debug: log.New(os.Stdout, "DEBUG: ", log.Ldate|log.Ltime|log.Lshortfile),
		level: level,
		mutex: &sync.Mutex{},
	}

	return log
}
