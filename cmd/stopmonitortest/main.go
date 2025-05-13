package main

import (
	"errors"
	"fmt"
	"log"
	"os"

	"github.com/BrandonIrizarry/MTA_Tracker/internal/geturl"
	"github.com/joho/godotenv"
)

const (
	stopMonitoringBaseURL = "https://bustime.mta.info/api/siri/stop-monitoring.json"
)

type config struct {
	apiKey string
}

func main() {
	cfg, initConfigErr := initConfig()

	if initConfigErr != nil {
		log.Fatal(initConfigErr)
	}

	queries := map[string]string{
		"key":           cfg.apiKey,
		"version":       "2",
		"MonitoringRef": "401386",
	}

	responseBytes, err := geturl.Call(stopMonitoringBaseURL, queries)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(string(responseBytes))
}

func initConfig() (config, error) {
	apiKey := os.Getenv("API_KEY")

	if apiKey == "" {
		return config{}, errors.New("API_KEY environment variable not set")
	}

	cfg := config{
		apiKey: apiKey,
	}

	return cfg, nil
}

func init() {
	godotenv.Load(".env")
}
