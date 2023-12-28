package main

import (
	"fmt"

	"github.com/boltx3/cybersecuritysuite/port"
)

func main() {
	var hostname string
	// var portnum int
	fmt.Println("Port Scanning")
	fmt.Println("Enter the hostname you want to scan")
	fmt.Scanln(&hostname)
	results := port.InitialScan(hostname)
	// fmt.Println("Enter the port you want to scan")
	// fmt.Scanln(&portnum)
	// results2 := port.ScanPort("TCP", hostname, portnum)
	// fmt.Println(results2)
	fmt.Println(results)
}
