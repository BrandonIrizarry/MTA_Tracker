// flags.go: Unmarshal command line arguments.
package main

import (
	"fmt"
	"log"
	"os"
	"strings"

	"nullprogram.com/x/optparse"
)

type flagConfig struct {
	newRoutes      []string
	obsoleteRoutes []string
}

func initFlags() (flagConfig, error) {
	// See which routes we're adding and/or removing.
	options := []optparse.Option{
		{Long: "add", Short: 'a', Kind: optparse.KindRequired},
		{Long: "remove", Short: 'r', Kind: optparse.KindRequired},
	}

	var fconfig flagConfig

	// For now 'rest' is unused. I might use it to query the bus
	// stop id.
	flagValues, _, optErr := optparse.Parse(options, os.Args)

	if optErr != nil {
		log.Fatal(optErr)
	}

	// Check for arguments that are duplicated across flags,
	// e.g. -a M11 -r M11.
	seenFlags := make(map[string]bool)

	for _, value := range flagValues {
		switch value.Long {
		case "add":
			fconfig.newRoutes = strings.Split(value.Optarg, ",")

			for _, route := range fconfig.newRoutes {
				seenFlags[route] = true
			}
		case "remove":
			fconfig.obsoleteRoutes = strings.Split(value.Optarg, ",")

			for _, route := range fconfig.obsoleteRoutes {
				if seenFlags[route] {
					return flagConfig{}, fmt.Errorf("Attempt to both add and remove route: %s", route)
				}
			}
		}
	}

	return fconfig, nil
}
