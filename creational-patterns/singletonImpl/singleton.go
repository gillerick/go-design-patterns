package main

import (
	"fmt"
	"sync"
)

type MyLogger struct {
	loglevel int
}

//Log a message using the logger
func (l *MyLogger) Log(s string) {
	fmt.Println(l.loglevel, ":", s)
}

//SetLogLevel sets the log level of the logger
func (l *MyLogger) SetLogLevel(level int) {
	l.loglevel = level
}

//The logger instance
var logger *MyLogger

//ToDo: Use the sync package to enforce goroutine safety
//This ensure that code runs once and is only available to a single go-routine
var once sync.Once

//ToDo: the getLoggerInstance function provides global access to the logger class instance
func getLoggerInstance() *MyLogger {
	once.Do(func() {
		if logger == nil {
			fmt.Println("Creating Logger Instance")
			logger = &MyLogger{}
		}
	})

	fmt.Println("Returning Logger Instance")
	return logger
}
