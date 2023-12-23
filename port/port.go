package port

import (
        "errors"
        "net"
        "strconv"
        "strings"
        "time"
)

type ScanResult struct {
        Port  string
        State string
}

func ScanPort(protocol, hostname string, port int) ScanResult {
        result := ScanResult{Port: protocol + "/" + strconv.Itoa(port)}
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

func InitialScan(hostname string) []ScanResult {
        var results []ScanResult

        for i := 1; i <= 1024; i++ {
                results = append(results, ScanPort("tcp", hostname, i))
                results = append(results, ScanPort("udp", hostname, i))
        }

        return results
}

func ScanPorts(hostname string, start, end int) []ScanResult {
        var results []ScanResult

        for i := start; i <= end; i++ {
                results = append(results, ScanPort("tcp", hostname, i))
                results = append(results, ScanPort("udp", hostname, i))
        }

        return results
}
// (like., "1-25")
func ParsePortRange(portRange string) (int, int, error) {
        parts := strings.Split(portRange, "-")
        if len(parts) != 2 {
                return 0, 0, errors.New("invalid port range format")
        }

        start, err := strconv.Atoi(parts[0])
        if err != nil {
                return 0, 0, err
        }

        end, err := strconv.Atoi(parts[1])
        if err != nil {
                return 0, 0, err
        }

        return start, end, nil
}
func ParseSinglePort(portStr string) (int, error) {
        port, err := strconv.Atoi(portStr)
        if err != nil {
                return 0, err
        }

        return port, nil
}