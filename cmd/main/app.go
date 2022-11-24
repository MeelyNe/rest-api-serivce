package main

import (
	"github.com/julienschmidt/httprouter"
	"net"
	"net/http"
	"os"
	"path"
	"path/filepath"
	"rest-api-service/internal/config"
	"rest-api-service/internal/user"
	"rest-api-service/pkg/logging"
	"time"
)

func start(router *httprouter.Router, logger *logging.Logger, cfg *config.Config) {
	// 127.0.0.1 - loopback interface address (localhost)
	// 0.0.0.0 -- listen to all the interfaces

	var listener net.Listener
	var listenerError error

	if cfg.Listen.Type == "sock" {
		// /path/to/binary(exec)
		// Dir() -- /path/to
		logger.Info("Starting the server on socket...")
		file, err := filepath.Abs(filepath.Dir(os.Args[0]))
		if err != nil {
			logger.Fatal("Error while getting the path to the executable file: ", err)
		}
		path.Join(file, "app.sock")
		logger.Info("Create unix socket...")
		listener, listenerError = net.Listen("unix", path.Join(file, "app.sock"))
		if listenerError != nil {
			logger.Fatal("Error while creating the unix socket: ", listenerError)
		}
		logger.Info("Server listening on socket...")
	} else {
		logger.Info("Starting the server on TCP...")
		logger.Info("Create TCP socket...")
		listener, listenerError = net.Listen("tcp", cfg.GetListenAddress())
		if listenerError != nil {
			logger.Fatal("Error while creating the TCP socket: ", listenerError)
		}
		logger.Info("Server is listening on ", cfg.GetListenAddress())
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
	logger.Info("Read configuration...")
	cfg := config.GetConfig()
	logger.Info("Registering the handlers...")
	user.NewHandler(logger).Register(router)
	logger.Info("Starting the server...")
	start(router, logger, cfg)
}
