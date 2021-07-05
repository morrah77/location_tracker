package server

import (
	"log"
	"morrah77.org/location_tracker/common"
	"morrah77.org/location_tracker/server/handlers"
	"net/http"
)

type Server struct {
	common.Storage
	mux *http.ServeMux
	pathPrefix string
	logger *log.Logger
}

func(s *Server) ServeHTTP(rw http.ResponseWriter, req *http.Request)  {
	s.mux.ServeHTTP(rw, req)
}

func NewServer(storage common.Storage, pathPrefix string, logger *log.Logger) http.Handler {
	var mux *http.ServeMux = http.NewServeMux();
	mux.Handle(pathPrefix, handlers.NewLocationHandler(storage, pathPrefix, logger));
	server := &Server{
		storage,
		mux,
		pathPrefix,
		logger,
	}
	return server;
}

