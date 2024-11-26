package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/zackarysantana/ember/trunk/internal/api"
)

func main() {
	port := ":8080"
	if p := os.Getenv("PORT"); p != "" {
		port = ":" + p
	}

	fmt.Println("MONGO", os.Getenv("MONGODB_URL"))

	fmt.Printf("Running on %s\n", port)
	if err := http.ListenAndServe(port, api.New()); err != nil {
		panic(err)
	}
}
