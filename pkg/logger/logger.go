package logger

import (
	"os"

	log "github.com/sirupsen/logrus"
)

var Log = log.WithFields(log.Fields{
	"logName":  "gobarber-golang",
	"logIndex": "message",
})

func SetupLogs() {
	log.SetFormatter(&log.JSONFormatter{})
	log.SetOutput(os.Stdout)
}
