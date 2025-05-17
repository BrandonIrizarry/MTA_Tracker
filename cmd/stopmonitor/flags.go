// flags.go: Unmarshal command line arguments given to the stopmonitor
// program.
package main

import (
	"log"
	"os"

	goptparse "github.com/BrandonIrizarry/goptparse/v2"
)

// A struct for holding command-line arguments, properly unmarshalled
// as Go values.
type flagConfig struct {
	filename string
}

// initFlags invokes the goptparse library, populating and returning a
// flagConfig struct, along with any errors.
func initFlags() (flagConfig, error) {
	options := []goptparse.Option{
		{
			Long:  "file",
			Short: 'f',
			Kind:  goptparse.KindRequired,
			Help:  "Filename from which to read saved bus routes.",
		},
	}

	var fconfig flagConfig

	flagValues, _, optErr := goptparse.Parse(options, os.Args)

	if optErr != nil {
		log.Fatal(optErr)
	}

	for _, value := range flagValues {
		switch value.Long {
		case "file":
			fconfig.filename = value.Optarg
		}
	}

	return fconfig, nil
}
