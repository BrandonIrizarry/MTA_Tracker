package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/BrandonIrizarry/MTA_Tracker/cmd/stoptable/internal/onebusaway"
	"github.com/BrandonIrizarry/MTA_Tracker/internal/geturl"
	"github.com/joho/godotenv"
)

const (
	// stopsForRouteBaseURL is the endpoint for fetching static
	// data about a given bus route. Here, its given as a
	// formattable string, so that the actual bus route is
	// configurable. For example, if the user queries the M11
	// route, the last component of the URL becomes 'MTA
	// NYCT_M11'.
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

	// Marshall the response bytes into a OneBusAway struct.
	var data onebusaway.OneBusAway

	if err := json.Unmarshal(responseBytes, &data); err != nil {
		log.Fatal(err)
	}

	fmt.Println(data)
}
