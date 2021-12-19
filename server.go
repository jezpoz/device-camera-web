package main

import (
	"fmt"
	"log"
	"os"
	"runtime"

	"github.com/joho/godotenv"

	"github.com/jezpoz/device-camera-web/platform/authenticator"
	"github.com/jezpoz/device-camera-web/platform/router"
)

func main() {
	ConfigRuntime()
	StartWorkers()
	StartApp()
}

// ConfigRuntime sets the number of operating system threads.
func ConfigRuntime() {
	nuCPU := runtime.NumCPU()
	runtime.GOMAXPROCS(nuCPU)
	fmt.Printf("Running with %d CPUs\n", nuCPU)
}

// StartWorkers start starsWorker by goroutine.
func StartWorkers() {
	go statsWorker()
}

func StartApp() {
	if err := godotenv.Load(); err != nil {
		log.Fatalf("Failed to load the env vars: %v", err)
	}

	auth, err := authenticator.New()
	if err != nil {
		log.Fatalf("Failed to initialize the authenticator: %v", err)
	}

	rtr := router.New(auth)
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	if err := rtr.Run(":" + port); err != nil {
		log.Panicf("error: %s", err)
	}
}
