package main

import (
  "time"
  log "github.com/sirupsen/logrus"
)


type UTCFormatter struct {
  log.Formatter
}

func (u UTCFormatter) Format(e *log.Entry) ([]byte, error) {
  e.Time = e.Time.UTC()
  return u.Formatter.Format(e)
}


func main() {
  // log.SetFormatter(&log.JSONFormatter{})
  log.SetFormatter(UTCFormatter{&log.JSONFormatter{}})
  log.SetLevel(log.InfoLevel)
  timestamp := time.Now().UnixNano()
  log.WithFields(log.Fields{
    "event_time": float64(timestamp)/ 1e9,
    "module": "main",
  }).Info("A walrus appears")
}
