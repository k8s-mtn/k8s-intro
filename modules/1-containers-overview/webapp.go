package main

import (
	"fmt"
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
	fmt.Println("Started example webapp")

	// Simple static webserver:
	log.Fatal(http.ListenAndServe(":8080", http.FileServer(http.Dir("assets/"+siteDir))))
}
