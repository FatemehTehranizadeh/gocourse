package main

import (
	"os"
	"time"

	"github.com/sirupsen/logrus"
)

func main() {
	username := "Reera"
	unixTime := time.Now().Unix()
	location := "Tehran"
	log := logrus.New()

	logFile, err := os.OpenFile("logruslogs.log", os.O_CREATE | os.O_APPEND | os.O_WRONLY, 0666)
	if err != nil {
		log.Fatalln("There is an error in creating the file: ", err)
	}
	defer logFile.Close()
	log.Out = logFile


	log.SetFormatter(&logrus.JSONFormatter{})
	log.WithFields(logrus.Fields{
		"username": username,
		"time":     unixTime,
	}).Warningln("This is a warning with fields")
	log.Info("This is an info log")
	log.Error("This is an error log")

	log.WithFields(logrus.Fields{
		"username": username,
		"time":     unixTime,
		"location": location,
	}).Infoln("This is an info log with fields")
}
