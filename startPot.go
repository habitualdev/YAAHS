package main

import (
	"LogPot/sshd"
	"LogPot/telnetd"
	"LogPot/wapot"
	"time"
)

func main(){
	go telnetd.StartTelnet(2223)
	go wapot.StartWapot(8080)
	go sshd.StartSsh("localtest",2222)
	time.Sleep(1*time.Minute)
}
