package port

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

func ParsePortInput(portInput string) ([]int, error) {
	var ports []int

	// By comma
	portList := strings.Split(portInput, ",")
	for _, part := range portList {
		// by hyphen -- range
		if strings.Contains(part, "-") {
			start, end, err := ParsePortRange(part)
			if err != nil {
				return nil, err
			}
			// Create list for the range
			for i := start; i <= end; i++ {
				ports = append(ports, i)
			}
		} else {
			// Single port
			portNum, err := ParseSinglePort(part)
			if err != nil {
				return nil, err
			}
			ports = append(ports, portNum)
		}
	}

	return ports, nil
}

func ParseSinglePort(portStr string) (int, error) {
	port, err := strconv.Atoi(portStr)
	if err != nil {
		return 0, err
	}

	if port > 65535 {
		return 0, fmt.Errorf("ports specified must be between 0 and 65535 inclusive")
	}

	return port, nil
}

// Helper function to parse port range string (e.g., "1-25")
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

	if start > end {
		return 0, 0, fmt.Errorf("your port range %d-%d is backwards. Did you mean %d-%d?",
			start, end, end, start)
	}

	if end > 65535 {
		return 0, 0, fmt.Errorf("ports specified must be between 0 and 65535 inclusive")
	}

	return start, end, nil
}
