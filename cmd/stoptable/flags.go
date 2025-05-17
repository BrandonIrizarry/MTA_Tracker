// flags.go: Unmarshal command line arguments given to the stoptable
// program.
package main

import (
	"fmt"
	"log"
	"os"
	"strings"

	goptparse "github.com/BrandonIrizarry/goptparse/v2"
)

// flagConfig holds the arguments passed from the command line,
// including rest/trailing arguments.
type flagConfig struct {
	newRoutes      []string
	obsoleteRoutes []string

	// This is derived from the command-line trailing arguments.
	query string
}

// initFlags handles the required optparse initialization steps.
//
// The only thing missing here is an easy 'help' option,
// unfortunately.
func initFlags() (flagConfig, error) {
	// See which routes we're adding and/or removing.
	options := []goptparse.Option{
		{
			Long:  "add",
			Short: 'a',
			Kind:  goptparse.KindRequired,
			Help: `Use a comma-separated list, no spaces, to add bus routes to the database.
                               For example: -a M11,M12`,
		},

		{
			Long:  "remove",
			Short: 'r',
			Kind:  goptparse.KindRequired,
			Help: `Use a comma-separated list, no spaces, to remove bus routes from the database.
                               For example: -r M11,M12`,
		},
	}

	var fconfig flagConfig

	// 'rest' is used to query the bus stop id.
	flagValues, rest, optErr := goptparse.Parse(options, os.Args)

	if optErr != nil {
		log.Fatal(optErr)
	}

	// Define a map used for checking for arguments that are
	// duplicated across flags and therefore result in an
	// ambiguous command, e.g. -a M11 -r M11.
	seenFlags := make(map[string]bool)

	// Note that arguments to -a and -r are given as
	// comma-delimited values, no spaces.
	for _, value := range flagValues {
		switch value.Long {
		case "add":
			fconfig.newRoutes = strings.Split(value.Optarg, ",")

			// Record that we saw this flag.
			for _, route := range fconfig.newRoutes {
				seenFlags[route] = true
			}
		case "remove":
			fconfig.obsoleteRoutes = strings.Split(value.Optarg, ",")

			// Now check for any duplicated flags.
			for _, route := range fconfig.obsoleteRoutes {
				if seenFlags[route] {
					return flagConfig{}, fmt.Errorf("Attempt to both add and remove route: %s", route)
				}
			}
		}
	}

	// For now, we accept only one query argument, though nothing
	// yet prohibits this from changing in the future.
	if len(rest) > 1 {
		return flagConfig{}, fmt.Errorf("stoptable accepts only one query argument")
	} else if len(rest) == 1 {
		fconfig.query = rest[0]
	}

	return fconfig, nil
}
