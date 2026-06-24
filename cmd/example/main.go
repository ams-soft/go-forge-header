package main

import (
	header "github.com/ams-soft/go-forge-header"
)

func main() {
	header.Print(nil, "example system", "0.0.1", "prod", false, new(8080))

	hdr := header.Buffer(nil, "example system", "0.0.1", "prod", false, new(8080))
	for _, line := range hdr {
		println(line)
	}
}
