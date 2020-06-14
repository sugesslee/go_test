package main

import (
	uuid "github.com/satori/go.uuid"
	log "github.com/sirupsen/logrus"
	"gostudy/log/hook"
	"os"
)

func initLog() {
	// Log as JSON instead of the default ASCII formatter.
	log.SetFormatter(&log.JSONFormatter{})

	// Output to stdout instead of the default stderr
	// Can be any io.Writer, see below for File example
	log.SetOutput(os.Stdout)
	//设置output,默认为stderr,可以为任何io.Writer，比如文件*os.File
	// file, _ := os.OpenFile("./log.log", os.O_CREATE|os.O_WRONLY, 0666)
	// log.SetOutput(file)
	// Only log the warning severity or above.
	log.SetLevel(log.WarnLevel)
	v1 := uuid.NewV1()
	log.AddHook(hook.NewTraceIdHook(v1.String() + " "))
}

//type Hook interface {
//	Levels() []Level
//	Fire(*Entry) error
//}
func main() {
	initLog()
	log.WithFields(log.Fields{
		"animal": "walrus",
	}).Info("A walrus appears")

	// A common pattern is to re-use fields between logging statements by re-using
	// the logrus.Entry returned from WithFields()
	contextLogger := log.WithFields(log.Fields{
		"common": "this is a common field",
		"other":  "I also should be logged always",
	})

	contextLogger.Info("I'll be logged with common and other field")
	contextLogger.Info("Me too")

	log.Trace("Something very low level.")
	log.Debug("Useful debugging information.")
	log.Info("Something noteworthy happened!")
	log.Warn("You should probably take a look at this.")
	log.Error("Something failed but I'm not quitting.")
	// Calls os.Exit(1) after logging
	log.Fatal("Bye.")
	// Calls panic() after logging
	log.Panic("I'm bailing.")
}
