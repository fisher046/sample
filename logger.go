package main

import (
	"io"
	"io/ioutil"
	"log"
	"os"
)

// Logger classified logs by log level
type Logger struct {
	Debug *log.Logger
	Info  *log.Logger
	Warn  *log.Logger
	Error *log.Logger
}

const (
	envLogPath  = "LOG_PATH"
	envLogLevel = "LOG_LEVEL"
)

var logger *Logger

func init() {
	logPath := os.Getenv(envLogPath)
	logLevel := os.Getenv(envLogLevel)

	var f io.Writer
	var err error
	if logPath != "" {
		f, err = os.OpenFile(logPath, os.O_CREATE|os.O_RDWR|os.O_APPEND, 0666)
		if err != nil {
			panic(err)
		}
	} else {
		f = os.Stdout
	}

	if logLevel == "Error" {
		logger = &Logger{
			Debug: log.New(ioutil.Discard, "DEBUG: ",
				log.Ldate|log.Ltime|log.Lshortfile),
			Info: log.New(ioutil.Discard, "INFO: ",
				log.Ldate|log.Ltime|log.Lshortfile),
			Warn: log.New(ioutil.Discard, "WARNING: ",
				log.Ldate|log.Ltime|log.Lshortfile),
			Error: log.New(f, "Error: ", log.Ldate|log.Ltime|log.Lshortfile),
		}
	} else if logLevel == "Warning" {
		logger = &Logger{
			Debug: log.New(ioutil.Discard, "DEBUG: ",
				log.Ldate|log.Ltime|log.Lshortfile),
			Info: log.New(ioutil.Discard, "INFO: ",
				log.Ldate|log.Ltime|log.Lshortfile),
			Warn:  log.New(f, "WARNING: ", log.Ldate|log.Ltime|log.Lshortfile),
			Error: log.New(f, "Error: ", log.Ldate|log.Ltime|log.Lshortfile),
		}
	} else if logLevel == "Info" {
		logger = &Logger{
			Debug: log.New(ioutil.Discard, "DEBUG: ",
				log.Ldate|log.Ltime|log.Lshortfile),
			Info:  log.New(f, "INFO: ", log.Ldate|log.Ltime|log.Lshortfile),
			Warn:  log.New(f, "WARNING: ", log.Ldate|log.Ltime|log.Lshortfile),
			Error: log.New(f, "Error: ", log.Ldate|log.Ltime|log.Lshortfile),
		}
	} else if logLevel == "Debug" {
		logger = &Logger{
			Debug: log.New(f, "DEBUG: ", log.Ldate|log.Ltime|log.Lshortfile),
			Info:  log.New(f, "INFO: ", log.Ldate|log.Ltime|log.Lshortfile),
			Warn:  log.New(f, "WARNING: ", log.Ldate|log.Ltime|log.Lshortfile),
			Error: log.New(f, "Error: ", log.Ldate|log.Ltime|log.Lshortfile),
		}
	} else {
		logger = &Logger{
			Debug: log.New(ioutil.Discard, "DEBUG: ",
				log.Ldate|log.Ltime|log.Lshortfile),
			Info: log.New(f, "INFO: ",
				log.Ldate|log.Ltime|log.Lshortfile),
			Warn: log.New(f, "WARNING: ",
				log.Ldate|log.Ltime|log.Lshortfile),
			Error: log.New(f, "Error: ",
				log.Ldate|log.Ltime|log.Lshortfile),
		}
	}
}
