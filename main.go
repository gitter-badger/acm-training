package main

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"github.com/uestc-acm/acm-training/config"
	"github.com/uestc-acm/acm-training/db"
	"github.com/uestc-acm/acm-training/router"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"
)

func main() {
	config.InitCommandLineFlags()
	if config.VersionPresent() {
		fmt.Printf("acm-training version %s\n", config.Version)
		os.Exit(0)
	}
	// Initalizes the database for web hosting, before the HTTP server starts.
	db.Init()

	r := router.Create()
	server := &http.Server{
		Addr:    ":8080",
		Handler: r,
	}

	go func() {
		log.Printf(
			`UESTC ACM/ICPC training backend
Server MODE: %v
HTTP port hosted at: %v
SQLite database served at: %v`,
			gin.Mode(), server.Addr, config.GetDbFile())
		// service connections
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
	}()

	// Wait for interrupt signal to gracefully shutdown the server with
	// a timeout of 5 seconds.
	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt)

	<-quit

	log.Println("Shutdown Server ...")
	db.DB().Close()

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := server.Shutdown(ctx); err != nil {
		log.Fatal("Server Shutdown:", err)
	}
	log.Println("Server exiting")
}
