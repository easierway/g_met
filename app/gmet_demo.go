package main

import (
	. "github.com/easierway/g_met"
	"log"
	"log/syslog"
	"math/rand"
	"time"

)

// The following example is to demonstrate how to use GMet.
func main() {
	sysLog, err := syslog.Dial("tcp", "localhost:1234",
		syslog.LOG_WARNING|syslog.LOG_DAEMON, "demo tag")
	if err != nil {
		log.Print(err)
	}

	// create GMet instance by given default writer and the formatter
	gmet := CreateGMetInstanceByDefault("seelog.xml")
	for i := 0; i < 100; i++ {
		gmet.Send(Metric("input_bytes", rand.Intn(100)),
			Metric("output_bytes", rand.Intn(100)))
		gmet.Flush() // in your real case, DON'T flush for each sending.
		sysLog.Emerg("And this is a daemon emergency with demotag.")
		// For seelog writer, the auto-flushing can be set in the log configuration
		time.Sleep(time.Second * 1)
	}
}
