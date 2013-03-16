package main

import (
	"log"
	"log/syslog"
)

func main() {
	l3, err := syslog.New(syslog.LOG_ERR, "GoExample")
	//l3, err := syslog.Dial("udp", "localhost", syslog.LOG_ERR, "GoExample") // connection to a log daemon
	defer l3.Close()
	if err != nil {
		log.Fatal("error")
	}

	//l3.Emerg("emergency")
	l3.Alert("alert")
	l3.Crit("critical")
	l3.Err("error")
	l3.Warning("warning")
	l3.Notice("notice")
	l3.Info("information")
	l3.Debug("debug")
	l3.Write([]byte("write"))
}
