package main

import (
	"flag"
	"fmt"
	"os"
	"strings"
	"text/tabwriter"

	"github.com/SaDs3c/sadmap/port"
)

func main() {
	target := flag.String("t", "localhost", "Target host or domain")
	portInput := flag.String("p", "80", "Port to scan (e.g., '80' or '1-25')")
	flag.Parse()

	printBanner("art.txt")

	var ports []int
	if *portInput != "" {
		var err error
		ports, err = parsePortInput(*portInput)
		if err != nil {
			fmt.Println("Error:", err)
			os.Exit(1)
		}
	}
	var results []port.ScanResult
	//for _, p := range ports {
	//      results = append(results, port.ScanPort("tcp", *target, p))
	//      results = append(results, port.ScanPort("udp", *target, p))
	//}

	for _, p := range ports {
		result := port.ScanPort("tcp", *target, p)
		if result.State == "Open" {
			results = append(results, result)
		}
	}
	printTable(results)
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

func parsePortInput(portInput string) ([]int, error) {
	var ports []int

	// By comma
	portList := strings.Split(portInput, ",")
	for _, part := range portList {
		// by hyphen -- range
		if strings.Contains(part, "-") {
			start, end, err := port.ParsePortRange(part)
			if err != nil {
				return nil, err
			}
			// Create list for the range
			for i := start; i <= end; i++ {
				ports = append(ports, i)
			}
		} else {
			// Single port
			portNum, err := port.ParseSinglePort(part)
			if err != nil {
				return nil, err
			}
			ports = append(ports, portNum)
		}
	}

	return ports, nil
}

func printTable(results []port.ScanResult) {
	w := tabwriter.NewWriter(os.Stdout, 0, 0, 1, ' ', tabwriter.Debug)
	fmt.Fprintf(w, "Port\tState\n")
	for _, r := range results {
		fmt.Fprintf(w, "%s\t%s\n", r.Port, r.State)
	}
	w.Flush()
}
