package sshd

import (
	"LogPot/logging"
	"fmt"
	"io"
	"github.com/gliderlabs/ssh"
	"strconv"
	"time"
)

func StartSsh(serverName string, portNum int) {
	ssh.Handle(func(s ssh.Session) {
		var passwordBuffer = make([]byte, 1)
		var password []byte
		var loopCheck = true
		io.WriteString(s, fmt.Sprintf("Hello %s\n", s.User()) + "\n")
		s.Write([]byte(s.User() + "@" + serverName + "'s password: "))
		for loopCheck {
			s.Read(passwordBuffer)
			if string(passwordBuffer) == "\n" || string(passwordBuffer) == "\r"{loopCheck = false}
				password = append(password, passwordBuffer...)
		}
		logging.SendToSshReciever("ssh - " + time.Now().Format(time.RFC3339) + " : " + s.User() + " : " + s.RemoteAddr().String() + " : " + string(password) + "\n")
	})

	ssh.ListenAndServe(":" + strconv.Itoa(portNum), nil)
}
