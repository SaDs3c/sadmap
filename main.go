package main

import (
	"flag"
	"fmt"
	"net"
	"os"
	"strconv"
	"sync"
	"text/tabwriter"
	"time"

	"github.com/SaDs3c/sadmap/port"
)

type ScanResult struct {
	Port  string
	State string
}

func main() {
	target := flag.String("t", "localhost", "Target host or domain")
	portInput := flag.String("p", "80", "Port to scan (e.g., '80' or '1-25')")
	flag.Parse()

	printBanner("art.txt")

	var ports []int
	if *portInput != "" {
		var err error
		ports, err = port.ParsePortInput(*portInput)
		if err != nil {
			fmt.Println("Error:", err)
			os.Exit(1)
		}
	}
	var results = make([]ScanResult, 0)

	start := time.Now()
	wg := &sync.WaitGroup{}
	for _, p := range ports {
		wg.Add(1)
		go func(protocol string, pt int) {
			defer wg.Done()
			result := ScanResult{Port: protocol + "/" + strconv.Itoa(pt)}
			address := *target + ":" + strconv.Itoa(pt)
			conn, err := net.DialTimeout(protocol, address, 1000*time.Millisecond)
			if err == nil { // not interested in closed ports or unknow ports :)
				result.State = "Open"
				results = append(results, result)
				conn.Close()
			}
		}("tcp", p)
	}

	wg.Wait()
	elapsed := time.Since(start)
	printTable(results)
	fmt.Println("SadMap done: in ", elapsed.Seconds(), "seconds")
	fmt.Println("Scaning as Ended.")
}

func printBanner(fileName string) {
	// Read art from file
	art, err := os.ReadFile(fileName)
	if err != nil {
		fmt.Println("Error reading ASCII art:", err)
		return
	}
	fmt.Println(string(art))
}

func printTable(results []ScanResult) {
	w := tabwriter.NewWriter(os.Stdout, 0, 0, 1, ' ', tabwriter.Debug)
	fmt.Fprintf(w, "Port\tState\n")
	for _, r := range results {
		fmt.Fprintf(w, "%s\t%s\n", r.Port, r.State)
	}
	w.Flush()
}
