package middleware

import (
	log "github.com/sirupsen/logrus"
	"net/http"
	"os"
	"time"
)


func LoggerMiddleware(h func(http.ResponseWriter, *http.Request)) func(http.ResponseWriter, *http.Request) {
	logger := log.New()

	logFileName := "sparta.out.log"
	file, err := os.OpenFile(logFileName, os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		panic("Cannot open file `" + logFileName + "`")
	}
	logger.Out = file
	logger.SetFormatter(&log.JSONFormatter{})

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		logger.WithFields(log.Fields{
			"requestTime": time.Now().Format("2006/1/2 15:04:05"),
			"host": r.Host,
			"method": r.Method,
			"path": r.URL.Path,
		}).Info("[sparta-log]")

		h(w, r)
	})
}

