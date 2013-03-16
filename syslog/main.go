package main

import (
	"log/syslog"
	"log"
)

func main() {
	l2, err := syslog.New(syslog.LOG_ERR, "GoExample")
	defer l2.Close()
	if err != nil {
		log.Fatal("error")
	}

    //l2.Emerg("emergency")
	l2.Alert("alert")
    l2.Crit("critical")
    l2.Err("error")
    l2.Warning("warning")
    l2.Notice("notice")
    l2.Info("information")
    l2.Debug("debug")
	l2.Write([]byte("write"))
}
