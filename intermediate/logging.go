package main

import (
	"log"
	"os"
)

func checkError(err error){
	if err != nil{
		panic(err)
	}
}


func main(){
	// log.Println("This is the first log message")
	// log.SetFlags(log.Lshortfile |log.Lmicroseconds |log.Lmsgprefix)
	// log.SetPrefix("This is a Prefix. ")
	// log.Println("This is the second log message")
	
	logFile, err := os.OpenFile("app.log", os.O_CREATE | os.O_APPEND | os.O_WRONLY, 0666)
	if err != nil {
		log.Fatal("Error while trying to create the file:", err)
	}

	infoLogger = log.New(logFile, "INFO: ", log.Ltime)
	warLogger = log.New(logFile, "WARNING: ", log.Ldate | log.Lshortfile | log.Lmsgprefix)

	infoLogger.Println("The username is incorrect")
	warLogger.Println("Be cautious!!!")

	defer logFile.Close()

	// More advanced libraries for logging: zap, logrus
	

}

var (
	infoLogger = log.New(os.Stdout, "INFO: ", log.Ltime)
	warLogger = log.New(os.Stdout, "WARNING: ", log.Ldate | log.Lshortfile | log.Lmsgprefix)
)