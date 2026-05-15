package main

import (
	header "github.com/ams-soft/go-forge-header"
)

func main() {
	header.Print(nil, "example system", "0.0.1", "prod", false)
}
