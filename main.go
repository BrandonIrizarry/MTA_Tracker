package main

import (
	"fmt"
	"log"
	"net/url"
	"os"

	"github.com/joho/godotenv"
)

const (
	stopMonitoringBaseURL = "https://bustime.mta.info/api/siri/stop-monitoring.json?&version=2"
)

func main() {
	godotenv.Load(".env")

	apiKey := os.Getenv("API_KEY")

	if apiKey == "" {
		log.Fatal("API_KEY environment variable not set")
	}

	stopMonitoringURL, err := initStopMonitoringURL(apiKey)

	if err != nil {
		log.Fatal(err)
	}

	queryValues := stopMonitoringURL.Query()

	queryValues.Add("LineRef", "M11")
	queryValues.Add("DirectionRef", "1")

	stopMonitoringURL.RawQuery = queryValues.Encode()

	fmt.Println(stopMonitoringURL)
}

func initStopMonitoringURL(apiKey string) (*url.URL, error) {
	apiKeyQuery := fmt.Sprintf("&key=%s", apiKey)
	stopMonitoringURL, err := url.Parse(stopMonitoringBaseURL + apiKeyQuery)

	if err != nil {
		return &url.URL{}, err
	}

	return stopMonitoringURL, nil
}
