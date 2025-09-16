package main

import (
	"log"
	"os"
)


func main(){
	log.SetPrefix("PREFIX: ")
	log.SetFlags(log.Lmsgprefix | log.Ltime | log.Lshortfile)
	log.Println("This is a log message")
	
	_, err := os.Stat("apptest.log")
	if err != nil {
		logFile, err := os.Create("apptest.log")
		if err != nil {
			log.Fatal("Error during creating the file:", err)
		}
		defer logFile.Close()
	} else {
		logFile, _ := os.Open("apptest.log")
		defer logFile.Close()
	}

	logFile , _ = os.Open("apptest.log")


	infoLogger := log.New(logFile, "INFO: ", log.Lshortfile | log.Ldate | log.Ltime)
	// infoLogger.Println("The driver's location is lat1, long1")

	warningLogger := log.New(logFile, "WARNING: ", log.Ldate | log.Ltime | log.Lmsgprefix)
	warningLogger.SetPrefix("Warning Message Prefix: ")
	// infoLogger.Println("Be cautios!!!")

	infoLogger.Println("The driver's location is lat2, long2")
	infoLogger.Println("Be cautios!!!2")
	infoLogger.Println("The driver's location is lat3, long3")
	infoLogger.Println("Be cautios!!!3")

}

