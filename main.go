package main

import (
	"fmt"

	"github.com/boltx3/GoCyber/port/port.go"
)

func main() {
	fmt.Println("Port Scanning")
	results := port.InitialScan("localhost")
	fmt.Printf(results)
}
