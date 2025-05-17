package main

import (
	"bufio"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/BrandonIrizarry/MTA_Tracker/cmd/stopmonitor/internal/siribus"
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

	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		stopID := scanner.Text()
		rawID, found := strings.CutPrefix(stopID, "MTA_")

		if !found {
			log.Fatal("MTA_ stop ID prefix missing")
		}

		responseBytes, callErr := geturl.Call(stopMonitoringBaseURL, map[string]string{
			"key":           cfg.apiKey,
			"version":       "2",
			"MonitoringRef": rawID,
		})

		if callErr != nil {
			log.Fatal(callErr)
		}

		// Marshall the response bytes into a SiriBus struct.
		var siriBusData siribus.SiriBus

		if err := json.Unmarshal(responseBytes, &siriBusData); err != nil {
			log.Fatal(err)
		}

		errorConditionJSON := siriBusData.Siri.ServiceDelivery.StopMonitoringDelivery[0].ErrorCondition

		// FIXME: the error message in the response itself
		// repeats three times. Can we amend this from our
		// end?
		if desc := errorConditionJSON.Description; desc != "" {
			log.Fatalf("Stop-monitoring API call failed: %s", desc)
		}

		// FIXME: for now, only look at the first entry of
		// StopMonitoringDelivery, until I figure out why this
		// is an array of multiple values, and not simply a
		// single value.
		for _, stopVisit := range siriBusData.Siri.ServiceDelivery.StopMonitoringDelivery[0].MonitoredStopVisit {
			lineRef := stopVisit.MonitoredVehicleJourney.LineRef
			destName := stopVisit.MonitoredVehicleJourney.DestinationName[0] // FIXME: why array?
			arrivalProximityText := stopVisit.MonitoredVehicleJourney.MonitoredCall.ArrivalProximityText

			fmt.Printf("%s to %s: %s\n", lineRef, destName, arrivalProximityText)
		}
	}
}

// initConfig encapsulates the code used to define the config struct's
// various fields, such as the API key.
//
// The newly constructed config struct is returned, along with an
// error.
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

// init handles any otherwise non-refactorable administrivia needed by
// the application at large, such as loading '.env'.
func init() {
	godotenv.Load(".env")
}
