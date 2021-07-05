package main

import (
	"io"
	"log"
	server2 "morrah77.org/location_tracker/server"
	"morrah77.org/location_tracker/storage"
	"net/http"
	"os"
)

var (
	addressToListen string = `:8080`
	certfile, keyfile string
	logStream io.Writer = os.Stdout
	logger *log.Logger
)


func init() {
	addressFromEnvVar := os.Getenv(`HISTORY_SERVER_LISTEN_ADDR`)
	if addressFromEnvVar != `` {
		// TODO validate the address
		addressToListen = addressFromEnvVar
	}
	logger = log.New(logStream, `location_tracker`, log.Flags())
	logger.Printf(`Listening on: %v\n`, addressToListen)
	certfile = `certs/location-tracker-dev.crt`
	keyfile = `certs/location-tracker-dev.key`

}

func main()  {
	var server *http.Server = &http.Server{
		Addr:              addressToListen,
		Handler:           server2.NewServer(storage.NewStorage(), `/location/`, logger),
		TLSConfig:         nil,
		ReadTimeout:       0,
		ReadHeaderTimeout: 0,
		WriteTimeout:      0,
		IdleTimeout:       0,
		MaxHeaderBytes:    0,
		TLSNextProto:      nil,
		ConnState:         nil,
		ErrorLog:          nil,
		BaseContext:       nil,
		ConnContext:       nil,
	}
	logger.Fatal(server.ListenAndServeTLS(certfile, keyfile))
}