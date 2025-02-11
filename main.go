package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/redis/go-redis/v9"
)

func main() {

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	// Create a basic handler
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello from Render!")
	})

	// Construct the address string: 0.0.0.0:PORT
	addr := fmt.Sprintf("0.0.0.0:%s", port)

	log.Printf("Server starting on %s", addr)

	// Start the server
	if err := http.ListenAndServe(addr, nil); err != nil {
		log.Fatalf("Error starting server: %v", err)
	}

	opt, err := redis.ParseURL("redis://red-cuk6fsij1k6c73d4j520:6379")
	if err != nil {
		panic(err)
	}
	client := redis.NewClient(opt)

	ctx := context.Background()

	err = client.Set(ctx, "foo", "bar", 0).Err()
	if err != nil {
		panic(err)
	}
	_ = client.Set(ctx, "fizz", "buzz", 0).Err()
}
