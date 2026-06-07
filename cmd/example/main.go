package main

import (
	header "github.com/ams-soft/go-forge-header"
)

func main() {
	header.Print(nil, "example system", "0.0.1", "prod", false)

	hdr := header.Buffer(nil, "example system", "0.0.1", "prod", false)
	for _, line := range hdr {
		println(line)
	}
}
