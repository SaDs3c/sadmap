package main

import (
        "fmt"

        "github.com/SaDs3c/sadmap/port"
)

func main() {
        fmt.Println("SadMap 0.1 - Powered by sadsec.")

        open := port.ScanPort("tcp", "localhost", 4000)
        fmt.Printf("Port Open: %t\n", open)

        results := port.InitialScan("localhost")
        fmt.Println(results)
}