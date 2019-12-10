package main

import (
	"fmt"
)

var (
	// Version is ...
	// use `go build -ldflags="-X 'main.Version=v1.0.0'" version.go` to set version in build time
	Version = "development"
)

func main() {
	fmt.Println("Version:\t", Version)
}
