package main

import "fmt"

// These fields are populated by by go build flags (see ./Makefile && docker/build.sh)
var (
	Version   string
	BuildDate string
)

// SDKVersion is a function that returns the version of the SDK.
// This function is called from the JS code.
// Parameters:
//   - name: the name of the SDK
//
// Returns:
//   - string: the version of the SDK
//
// Example:
//
//	SDKVersion(); // returns "Cosmos-Checkers-SDK v0.0.1 (2021-09-01)"
func SDKVersion(name string) string {
	return fmt.Sprintf("%s v%s (%s)", name, Version, BuildDate)
}
