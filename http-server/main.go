package main

import (
	"context"
	"fmt"
	"http-server/controller"
	"http-server/db"
	"http-server/server"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	// creating a new server instance
	s := server.NewServer(8080)

	// creating a database instance (similar to connecting an external database like postgres, mongodb etc..)
	db := db.NewDatabase()

	// creating a handler instance which takes in a db instance
	// implements all the handler function along with db operations
	handler := controller.NewController(db)

	// defining middlewares -> wrapping our handler function with extra functions so
	// it can be executed in the same manner
	s.Use(server.LoggingMiddleware) // handler -> LoggingMiddleware(Handler)
	s.Use(server.RecoveryMiddleware) // handler -> RecoveryMiddleware(LoggingMiddleware(Handler))

	// Registering all the routes with there respective handler function
	s.Router.GET("/", handler.Home)
	s.Router.GET("/time", handler.GetTime)
	s.Router.GET("/users", handler.UserHandler)
	s.Router.POST("/user", handler.CreateUserHandler)

	// Starting the server in a go-routine so the main thread can listen for terminal requests (Ctrl+C)
	// create a channel so when the main threads recieves the signal it notifies the goroutine to shutdown
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt, syscall.SIGTERM)

	go func(){
		// Starting the server
		if err := s.Run(); err != nil {
			fmt.Printf("Error running the server%v\n", err)
		}
	}()

	fmt.Println("Server is running on localhost:8080")
	
	// Blocking the main thread until interupt signal is received
	<-stop

	// creating a context with 5 second timeout, after context timeouts we forcefully shut down the server
	ctx, cancel := context.WithTimeout(context.Background(), 5 * time.Second)
	defer cancel() // Release resources when done IMP

	if err := s.ShutDown(ctx); err != nil {
		log.Fatalf("Forcefully shutting down the server due to %v", err.Error())
	}

	fmt.Println("Server gracefully stoopped")
}
