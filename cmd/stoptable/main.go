package main

import (
	"context"
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"os"

	"github.com/BrandonIrizarry/MTA_Tracker/cmd/stoptable/internal/database"
	"github.com/BrandonIrizarry/MTA_Tracker/cmd/stoptable/internal/onebusaway"
	"github.com/BrandonIrizarry/MTA_Tracker/internal/geturl"
	"github.com/joho/godotenv"

	_ "github.com/mattn/go-sqlite3"
)

const (
	// stopsForRouteBaseURL is the endpoint for fetching static
	// data about a given bus route. Here, its given as a
	// formattable string, so that the actual bus route is
	// configurable. For example, if the user queries the M11
	// route, the last component of the URL becomes 'MTA
	// NYCT_M11'.
	stopsForRouteBaseURL = "https://bustime.mta.info/api/where/stops-for-route/%s.json"

	// tableDBName is the name of the SQLite table where the
	// id-name associations will be stored and queried from.
	tableDBName = "./cmd/stoptable/stoptable.db"
)

type config struct {
	apiKey    string
	dbQueries *database.Queries
}

func main() {
	cfg, err := initConfig()

	if err != nil {
		log.Fatal(err)
	}

	fconfig, initFlagErr := initFlags()

	if initFlagErr != nil {
		log.Fatal(initFlagErr)
	}

	for _, shortRoute := range fconfig.newRoutes {
		route := fmt.Sprintf("MTA NYCT_%s", shortRoute)
		err := addRoute(cfg, route)

		if err != nil {
			if errRouteExists, ok := err.(*RouteExistsError); ok {
				log.Printf("%v; skipping", errRouteExists)
			} else {
				log.Fatal(err)
			}
		}
	}

	for _, shortRoute := range fconfig.obsoleteRoutes {
		route := fmt.Sprintf("MTA NYCT_%s", shortRoute)
		err := removeRoute(cfg, route)

		if err != nil {
			if errRouteDoesNotExist, ok := err.(*RouteDoesNotExistError); ok {
				log.Printf("%v; skipping", errRouteDoesNotExist)
			} else {
				log.Fatal(err)
			}
		}
	}

	// Handle the query, if it exists.
	if fconfig.query != "" {
		results, err := cfg.dbQueries.QueryStopsBySubstring(context.Background(), fconfig.query)

		if err != nil {
			log.Fatal(err)
		}

		for _, result := range results {
			fmt.Println(result)
		}
	}
}

// addRoute adds stops for the given routeID to the database file. It
// queries the appropriate endpoint to get this information.
//
// cfg is needed to perform database queries, such as the actual
// data-inclusion operation.
//
// Attempting to add a route that's already been added to the database
// is an error.
func addRoute(cfg config, routeID string) error {
	routeExists, err := cfg.dbQueries.TestRouteExists(context.Background(), routeID)

	if err != nil {
		return err
	}

	if routeExists == "1" {
		return &RouteExistsError{routeID: routeID}
	}

	stopsForRouteBaseURLFilled := fmt.Sprintf(stopsForRouteBaseURL, routeID)

	responseBytes, err := geturl.Call(stopsForRouteBaseURLFilled, map[string]string{
		"key":              cfg.apiKey,
		"version":          "2",
		"includePolylines": "false",
	})

	if err != nil {
		return err
	}

	// Marshall the response bytes into a OneBusAway struct.
	var data onebusaway.OneBusAway

	if err := json.Unmarshal(responseBytes, &data); err != nil {
		return err
	}

	if code := data.Code; code != 200 {
		return fmt.Errorf("OBA response reported '%d' response code, aborting", code)
	}

	for _, stop := range data.Data.References.Stops {
		stopParams := database.CreateStopParams{
			StopID:  stop.ID,
			Name:    stop.Name,
			RouteID: routeID,
		}

		if err := cfg.dbQueries.CreateStop(context.Background(), stopParams); err != nil {
			log.Print("Database error")
			return err
		}
	}

	// Add this log message to confirm that the application did
	// something.
	log.Printf("Added route '%s'", routeID)

	return nil
}

// removeRoute removes the stops associated with the given routeID.
//
// Note that it doesn't need to make a network request.
//
// Attempting to remove a route that's no longer (or was never)
// present in the database is an error.
func removeRoute(cfg config, routeID string) error {
	routeExists, err := cfg.dbQueries.TestRouteExists(context.Background(), routeID)

	if err != nil {
		return err
	}

	if routeExists == "0" {
		return &RouteDoesNotExistError{routeID: routeID}
	}

	if err := cfg.dbQueries.ClearStopsByRoute(context.Background(), routeID); err != nil {
		log.Print("Database error")
		return err
	}

	// Add this log message to confirm that the application did
	// something.
	log.Printf("Removed route '%s'", routeID)

	return nil
}

// initConfig encapsulates the code used to define the config struct's
// various fields, such as the API key and database-query handle.
//
// The newly constructed config struct is returned, along with an
// error.
func initConfig() (config, error) {
	apiKey := os.Getenv("API_KEY")

	if apiKey == "" {
		return config{}, errors.New("API_KEY environment variable not set")
	}

	db, err := sql.Open("sqlite3", tableDBName)

	if err != nil {
		return config{}, fmt.Errorf("Couldn't open database file '%s'\n", tableDBName)
	}

	dbQueries := database.New(db)

	cfg := config{
		apiKey:    apiKey,
		dbQueries: dbQueries,
	}

	return cfg, nil
}

// init handles any otherwise non-refactorable administrivia needed by
// the application at large, such as loading '.env'.
func init() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	godotenv.Load(".env")
}
