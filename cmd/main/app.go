package main

import (
	"github.com/julienschmidt/httprouter"
	"net"
	"net/http"
	"rest-api-service/internal/user"
	"rest-api-service/pkg/logging"
	"time"
)

func start(router *httprouter.Router, logger *logging.Logger) {
	// 127.0.0.1 - loopback interface address (localhost)
	// 0.0.0.0 -- listen to all the interfaces
	logger.Info("Listening to port 1234")
	listener, err := net.Listen("tcp", "0.0.0.0:1234") // why? because we can use in socket and on ip ( when socket - unix )
	if err != nil {
		panic(err)
	}

	// specify the server settings
	// ReadTimeout: maximum duration for reading the entire request, including the body
	// WriteTimeout: maximum duration before timing out writes of the response
	logger.Info("Setting the server settings...")
	server := &http.Server{
		Handler:      router,
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	logger.Info("Starting the server...")
	logger.Panic(server.Serve(listener))
}

// Idea: make an own struct logger with own interface and solve with multiple logger
func main() {
	logging.Init()
	logger := logging.GetLogger()
	logger.Info("Starting the application...")
	router := httprouter.New()
	logger.Info("Registering the handlers...")
	user.NewHandler(logger).Register(router)
	logger.Info("Starting the server...")
	start(router, logger)
}
