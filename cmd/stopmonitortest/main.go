package main

import (
	"fmt"
	"log"
	"os"

	"github.com/BrandonIrizarry/MTA_Tracker/internal/geturl"
	"github.com/joho/godotenv"
)

const (
	stopMonitoringBaseURL = "https://bustime.mta.info/api/siri/stop-monitoring.json"
)

func main() {
	godotenv.Load(".env")

	apiKey := os.Getenv("API_KEY")

	if apiKey == "" {
		log.Fatal("API_KEY environment variable not set")
	}

	queries := map[string]string{
		"key":           apiKey,
		"version":       "2",
		"LineRef":       "M11",
		"MonitoringRef": "401386",
	}

	responseBytes, err := geturl.Call(stopMonitoringBaseURL, queries)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(string(responseBytes))
}
