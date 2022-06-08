package main

import "fmt"

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

//ToDo: the getLoggerInstance function provides global access to the logger class instance
func getLoggerInstance() *MyLogger {
	if logger == nil{
		fmt.Println("Creating Logger Instance")
		logger = &MyLogger{}
	}
	fmt.Println("Returning Logger Instance")
	return logger
}
