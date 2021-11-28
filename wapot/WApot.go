package wapot

import (
	"LogPot/logging"
	"fmt"
	"net/http"
	"strconv"
	"time"
)

func StartWapot(portNum int) {
	console_log := 1
	port := ":" + strconv.Itoa(portNum)
	fmt.Println("wapot v1.0\n\n[+] Listening on port", port)
	http.Handle("/", http.FileServer(http.Dir("./http")))
	if err := http.ListenAndServe(port, logRequest(http.DefaultServeMux, console_log)); err != nil {
		panic(err)
	}

}

func logRequest(handler http.Handler, clog int) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		logMe := fmt.Sprintln("http - %s : %s %s %s \"%s\" \"%s\" \n", time.Now().Format(time.RFC3339),req.RemoteAddr, req.Header.Get("X-Forwarded-For"), req.Method, req.URL.RequestURI(), req.UserAgent())
		logging.SendToHttpReciever(logMe)
		handler.ServeHTTP(w, req)
	})
}




