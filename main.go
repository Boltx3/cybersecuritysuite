package main

import (
	"fmt"
<<<<<<< Updated upstream

	"github.com/boltx3/cybersecuritysuite/port"
=======
>>>>>>> Stashed changes
)

func main() {
	var hostname string
	fmt.Println("Port Scanning")
<<<<<<< Updated upstream
	fmt.Println("Enter the hostname you want to scan")
	fmt.Scanln(&hostname)
	results := port.InitialScan(hostname)
	results2 := port.ScanPort("TCP", hostname, 80)
	fmt.Println(results)
	fmt.Println(results2)
=======
	open := scanPort("tcp", "192.168.1.250", 9000)
	fmt.Printf("Port Open: %t\n", open)
>>>>>>> Stashed changes
}
