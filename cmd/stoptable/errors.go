package main

import "fmt"

type RouteExistsError struct {
	routeID string
}

func (r RouteExistsError) Error() string {
	return fmt.Sprintf("Route '%s' exists", r.routeID)
}

type RouteDoesNotExistError struct {
	routeID string
}

func (r RouteDoesNotExistError) Error() string {
	return fmt.Sprintf("Route '%s' doesn't exist", r.routeID)
}
