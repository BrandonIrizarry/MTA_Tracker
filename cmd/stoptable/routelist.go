// routelist.go: Define command-line flags for describing routes to be
// added, and routes to be deleted.
//
// Since these are given in a bespoke manner on the command line as a
// comma-separated list, we must define them eventually as VarP, and
// therefore we need to make them satisfy the required Value interface
// defined in the pflags package. This version of Value adds a third
// method - Type - not found in the flags package.
package main

import (
	"errors"
	"fmt"
	"strings"
)

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
