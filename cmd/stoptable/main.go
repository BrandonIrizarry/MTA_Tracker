package main

import (
	"fmt"
	"log"
	"os"

	"github.com/BrandonIrizarry/MTA_Tracker/internal/geturl"
	"github.com/joho/godotenv"
)

const (
	stopsForRouteBaseURL = "https://bustime.mta.info/api/where/stops-for-route/MTA NYCT_%s.json"
)

func main() {
	godotenv.Load(".env")

	apiKey := os.Getenv("API_KEY")

	if apiKey == "" {
		log.Fatal("API_KEY environment variable not set")
	}

	queries := map[string]string{
		"key":              apiKey,
		"version":          "2",
		"includePolylines": "false",
	}

	stopsForRouteBaseURLFilled := fmt.Sprintf(stopsForRouteBaseURL, "M11")

	responseBytes, err := geturl.Call(stopsForRouteBaseURLFilled, queries)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(string(responseBytes))
}
