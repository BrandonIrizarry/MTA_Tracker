package main

import (
	"context"
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/BrandonIrizarry/MTA_Tracker/cmd/stoptable/internal/database"
	"github.com/BrandonIrizarry/MTA_Tracker/cmd/stoptable/internal/onebusaway"
	"github.com/BrandonIrizarry/MTA_Tracker/internal/geturl"
	"github.com/joho/godotenv"
	"github.com/spf13/pflag"

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

type routeList []string

func (rlist *routeList) Set(value string) error {
	if len(*rlist) > 0 {
		return errors.New("interval flag already set")
	}

	for _, route := range strings.Split(value, ",") {
		*rlist = append(*rlist, route)
	}

	return nil
}

func (rlist *routeList) Type() string {
	return "routeList"
}

func (rlist *routeList) String() string {
	return fmt.Sprint(*rlist)
}

// Flags.
var newRoutes routeList
var obsoleteRoutes routeList

type RouteExistsError struct {
	routeID string
}

func (r RouteExistsError) Error() string {
	return fmt.Sprintf("Route '%s' exists", r.routeID)
}

func main() {
	// See which routes we're adding and/or removing.
	pflag.Parse()

	cfg, err := initConfig()

	if err != nil {
		log.Fatal(err)
	}

	for _, shortRoute := range newRoutes {
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
}

func addRoute(cfg config, routeID string) error {
	routeExists, err := cfg.dbQueries.TestRouteExists(context.Background(), routeID)

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

	return nil
}

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

func init() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	godotenv.Load(".env")

	pflag.VarP(&newRoutes, "add", "a", "Add a route")
	pflag.VarP(&obsoleteRoutes, "remove", "r", "Remove a route")
}
