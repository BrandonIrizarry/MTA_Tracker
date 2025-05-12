// flags.go: Unmarshal command line arguments.
package main

import (
	"log"
	"os"
	"strings"

	"nullprogram.com/x/optparse"
)

type flagConfig struct {
	newRoutes      []string
	obsoleteRoutes []string
}

func initFlags() flagConfig {
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

	for _, value := range flagValues {
		switch value.Long {
		case "add":
			fconfig.newRoutes = strings.Split(value.Optarg, ",")
		case "remove":
			fconfig.obsoleteRoutes = strings.Split(value.Optarg, ",")
		}
	}

	return fconfig
}
