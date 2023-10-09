package main

import (
	"context"
	configpackage "dc-nearshore/cmd/api/config"
	"net/http"
	"os"
	"time"

	"github.com/rs/zerolog/log"
)

func main() {

	ds, err := configpackage.InitDS()

	if err != nil {
		log.Error().Str("error", "error InitDS")
		os.Exit(1)
	}

	router, errConfig := configpackage.SetUp(ds)

	if errConfig != nil {
		log.Error().Str("error", "Failure to setup")
		os.Exit(1)
	}

	log.Print("Starting server ...")

	srv := &http.Server{
		Addr:    ":8080",
		Handler: router,
	}

	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Printf("Failed to initialize server: %v", err)
		}
	}()

	// Wait for kill signal of channel
	quit := make(chan os.Signal)

	<-quit

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// Shutdown server
	if err := srv.Shutdown(ctx); err != nil {
		log.Printf("Server forced to shutdown: %v", err)
	}
}
