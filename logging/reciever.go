package logging

import (
	"os"
	"strings"
	"time"
)

func SendToTelnetReciever(logMe string){
	logFile, _ := os.OpenFile("HoneyPot.log", os.O_CREATE|os.O_APPEND|os.O_RDWR, 0644)
	defer logFile.Close()
	writeBuf := strings.Replace(logMe, "\n", "\n" + "telnet - " + time.Now().Format(time.RFC3339) + " : ", -1)
	logFile.Write([]byte(writeBuf))
}

func SendToSshReciever(logMe string){
	logFile, _ := os.OpenFile("HoneyPot.log", os.O_CREATE|os.O_APPEND|os.O_RDWR, 0644)
	defer logFile.Close()
	writeBuf := logMe
	logFile.Write([]byte(writeBuf))
}

func SendToHttpReciever(logMe string){
	logFile, _ := os.OpenFile("HoneyPot.log", os.O_CREATE|os.O_APPEND|os.O_RDWR, 0644)
	defer logFile.Close()
	writeBuf := logMe
	logFile.Write([]byte(writeBuf))
}