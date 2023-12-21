package main

import (
	"fmt"

	"github.com/Boltx3/cybersecuritysuite/port"
)

func main() {
	fmt.Println("Port Scanning")
	results := port.InitialScan("localhost")
	fmt.Printf(results)
}
