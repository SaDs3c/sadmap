package main

import (
        "fmt"
        "os"
        "text/tabwriter"

        "github.com/SaDs3c/sadmap/port"
)

func main() {
        fmt.Println("SadMap 0.1 - Powered by sadsec.")

        open := port.ScanPort("tcp", "localhost", 4000)
        fmt.Printf("Port Open: %t\n", open)

        results := port.InitialScan("localhost")
        printTable(results)
}


func printTable(results []port.ScanResult) {
        w := tabwriter.NewWriter(os.Stdout, 0, 0, 1, ' ', tabwriter.Debug)

        fmt.Fprintf(w, "Port\tState\n")
        fmt.Println("------- | ------")
        for _, r := range results {
                fmt.Fprintf(w, "%s\t%s\n", r.Port, r.State)
        }
        w.Flush()
}