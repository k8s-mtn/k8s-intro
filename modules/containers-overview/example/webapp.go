package main

import (
	"log"
	"net/http"
	"os"
)

func main() {
	siteVar := os.Getenv("SITE")
	siteDir := "site1"
	if siteVar == "2" {
		siteDir = "site2"
	}

	// Simple static webserver:
	log.Fatal(http.ListenAndServe(":8080", http.FileServer(http.Dir("assets/"+siteDir))))
}
