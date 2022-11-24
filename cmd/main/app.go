package main

import (
	"github.com/julienschmidt/httprouter"
	"log"
	"net"
	"net/http"
	"rest-api-service/internal/user"
	"time"
)

func start(router *httprouter.Router) {
	// 127.0.0.1 - loopback interface address (localhost)
	// 0.0.0.0 -- listen to all the interfaces
	listener, err := net.Listen("tcp", "0.0.0.0:1234") // why? because we can use in socket and on ip ( when socket - unix )
	if err != nil {
		panic(err)
	}

	// specify the server settings
	// ReadTimeout: maximum duration for reading the entire request, including the body
	// WriteTimeout: maximum duration before timing out writes of the response
	server := &http.Server{
		Handler:      router,
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	log.Println("Server started on port 1234")
	log.Fatalln(server.Serve(listener))
}

func main() {
	// logger init
	log.Println("Starting the application...")
	router := httprouter.New()
	log.Println("Registering the handlers...")
	user.NewHandler().Register(router)
	log.Println("Starting the server...")
	start(router)
}
