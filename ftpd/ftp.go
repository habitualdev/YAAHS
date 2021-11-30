package ftpd

import (
	"LogPot/logging"
	"fmt"
	"goftp.io/server/v2"
	"goftp.io/server/v2/driver/file"
)

type localLogger server.StdLogger
func (l localLogger) Print(sessionID string, message interface{}) {
	logging.SendToFtpReciever(fmt.Sprintf("%s  %s", sessionID, message))
}
func (l localLogger) Printf(sessionID string, format string, v ...interface{}) {
	logging.SendToFtpReciever(sessionID + fmt.Sprintf(format, v...))
}
func (l localLogger) PrintCommand(sessionID string, command string, params string) {
	if command == "PASS" {
		logging.SendToFtpReciever(fmt.Sprintf("%s > PASS ****", sessionID))
	} else {
		logging.SendToFtpReciever(fmt.Sprintf("%s > %s %s", sessionID, command, params))
	}
}
func (l localLogger) PrintResponse(sessionID string, code int, message string) {
	logging.SendToFtpReciever(fmt.Sprintf("%s < %d %s", sessionID, code, message))
}



func StartFtp(serverName string, portNum int, welcomeMessage string){

	logger := localLogger{}

	driver, _ := file.NewDriver("./dirtypot")

	opts := &server.Options{
		Commands:       nil,
		Driver:         driver,
		Auth:           nil,
		Perm:           nil,
		Name:           serverName,
		Hostname:       "",
		PublicIP:       "",
		PassivePorts:   "",
		Port:           portNum,
		WelcomeMessage: welcomeMessage,
		Logger:         logger,
		RateLimit:      0,
	}
	
	server.NewServer(opts)
}