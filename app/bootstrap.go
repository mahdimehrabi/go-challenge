package app

import (
	"challange/app/controller"
	"challange/app/infrastracture"
	"challange/app/repository"
	"challange/app/routes"
	"challange/app/services"
	"context"
	"go.uber.org/fx"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"
)

var BootstrapModule = fx.Options(
	infrastracture.Module,
	repository.Module,
	services.Module,
	controller.Module,
	routes.Module,
	fx.Invoke(Bootstrap),
)

func Bootstrap(
	lifecycle fx.Lifecycle,
	logger infrastracture.SegmentLogger,
	segmentRoutes routes.SegmentRoutes,
	db infrastracture.PgxDB,
) {
	port := os.Getenv("ServerPort")

	// create a new serve mux and register routes
	sm := http.NewServeMux()
	segmentRoutes.AddRoutes(sm)

	// create a new server
	s := http.Server{
		Addr:         ":" + port,        // configure the bind address
		Handler:      sm,                // set the default handler
		ErrorLog:     logger.LG,         // set the logger for the server
		ReadTimeout:  5 * time.Second,   // max time to read request from the client
		WriteTimeout: 10 * time.Second,  // max time to write response to the client
		IdleTimeout:  120 * time.Second, // max time for connections using TCP Keep-Alive
	}

	// start the server
	go func() {
		logger.LG.Println("Starting server on port " + port)

		err := s.ListenAndServe()
		if err != nil {
			logger.LG.Printf("Error starting server: %s\n", err)
			os.Exit(1)
		}
	}()

	// trap sigterm or interupt and gracefully shutdown the server
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	signal.Notify(c, os.Kill)

	// Block until a signal is received.
	sig := <-c
	log.Println("Got signal:", sig)

	// gracefully shutdown the server, waiting max 30 seconds for current operations to complete
	ctx, _ := context.WithTimeout(context.Background(), 30*time.Second)
	s.Shutdown(ctx)

}
