package telnetd

import (
	"LogPot/logging"
	"github.com/reiver/go-oi"
	"github.com/reiver/go-telnet"
	"strconv"
	"time"
)

var EchoAndLog telnet.Handler = echoAndLogHandler{}


type echoAndLogHandler struct{}


func (handler echoAndLogHandler) ServeTELNET(ctx telnet.Context, w telnet.Writer, r telnet.Reader) {

	var buffer [1]byte // Seems like the length of the buffer needs to be small, otherwise will have to wait for buffer to fill up.
	p := buffer[:]
	timeStamp :=time.Now().Format(time.RFC3339) + " : "
	logging.SendToTelnetReciever("telnet - " + timeStamp)
	for {
		n, err := r.Read(p)
		if n > 0 {
			logging.SendToTelnetReciever(string(p[:n]))
			oi.LongWrite(w, p[:n])
		}
		if nil != err {
			break
		}
	}
}

func StartTelnet(portNum int) {

	var handler telnet.Handler = echoAndLogHandler{}

	err := telnet.ListenAndServe(":" + strconv.Itoa(portNum), handler)
	if nil != err {

		panic(err)
	}
}