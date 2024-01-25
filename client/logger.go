package main

import (
	"log"
	"os"
)

var (
	ErrorLogger *log.Logger
	InfoLoggger *log.Logger
)

func InitLoggers() {
	InfoLoggger = log.New(os.Stdout, "[INFO] ", log.Ldate|log.Ltime|log.Lshortfile)
	ErrorLogger = log.New(os.Stdout, "[ERROR] ", log.Ldate|log.Ltime|log.Lshortfile)
}
