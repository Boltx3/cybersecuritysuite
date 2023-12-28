package main

import (
	"fmt"

	"github.com/boltx3/cybersecuritysuite/port"
)

func main() {
	var hostname string
	fmt.Println("Port Scanning")
	fmt.Println("Enter the hostname you want to scan")
	fmt.Scanln(&hostname)
	results := port.InitialScan(hostname)
	results2 := port.ScanPort("TCP", hostname, 80)
	fmt.Println(results)
	fmt.Println(results2)
}
