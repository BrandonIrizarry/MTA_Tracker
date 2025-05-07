package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"os"
	"time"

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
	queryValues.Add("MonitoringRef", "401386")

	stopMonitoringURL.RawQuery = queryValues.Encode()

	jsonDump, err := callURL(*stopMonitoringURL)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(string(jsonDump))
}

func callURL(url url.URL) ([]byte, error) {
	client := &http.Client{
		Timeout: 5 * time.Second,
	}

	req, err := http.NewRequest("GET", url.String(), nil)

	if err != nil {
		return nil, err
	}

	res, err := client.Do(req)

	if err != nil {
		return nil, err
	}

	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)

	if err != nil {
		return nil, err
	}

	return body, nil
}

func initStopMonitoringURL(apiKey string) (*url.URL, error) {
	apiKeyQuery := fmt.Sprintf("&key=%s", apiKey)
	stopMonitoringURL, err := url.Parse(stopMonitoringBaseURL + apiKeyQuery)

	if err != nil {
		return &url.URL{}, err
	}

	return stopMonitoringURL, nil
}
