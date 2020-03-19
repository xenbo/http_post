package log

import (
	"flag"
	"fmt"
	"log"
	"os"
	"runtime"
)

var (
	logFileName = flag.String("log", "xxxxxx.log", "Log file name")
)

var DLog *log.Logger

func CreateLog() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	flag.Parse()

	//set logfile Stdout
	logFile, logErr := os.OpenFile(*logFileName, os.O_CREATE|os.O_RDWR|os.O_TRUNC, 0666) //O_APPEND
	if logErr != nil {
		fmt.Println("Fail to find", *logFile, "cServer start Failed")
		os.Exit(1)
	}

	DLog = log.New(logFile,"[Info]",log.Llongfile)
	DLog.SetOutput(logFile)
	DLog.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)

	//debugLog.SetPrefix("[Debug]")
	//write log
	//DLog.Printf("Server abort! Cause:%v \n", "test log file"）
}

