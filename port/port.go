package port

import (
	"fmt"
	"net"
	"strconv"
	"sync"
	"time"
)

type ScanResult struct {
	Hostname string
	Port     int
	State    string
}

// Given a hostname, protocol, and port it scans and returns if port is open or closed
func ScanPort(protocol, hostname string, port int) ScanResult {
	result := ScanResult{Hostname: hostname, Port: port}
	address := hostname + ":" + strconv.Itoa(port)
	conn, err := net.DialTimeout(protocol, address, 60*time.Second)

	if err != nil {
		result.State = "Closed"
		return result
	}
	defer conn.Close()
	result.State = "Open"
	return result
}

// Given a hostname, returns the ports open from 1-1024
func InitialScan(hostname string) []ScanResult {
	var wg sync.WaitGroup
	var results []ScanResult
	for i := 0; i <= 1024; i++ {
		wg.Add(1)
		results = append(results, ScanPort("tcp", hostname, i))
		// fmt.Println("Done scanning ", i)
		wg.Done()
		fmt.Println(results)
	}
	wg.Wait()
	return results
}
