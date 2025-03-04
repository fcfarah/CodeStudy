package main

import (
	"fmt"
	"strconv"
	"strings"
)

// ipToInt converts an IPv4 address to an integer.
func ipToInt(ip string) int {
	parts := strings.Split(ip, ".")
	p0, _ := strconv.Atoi(parts[0])
	p1, _ := strconv.Atoi(parts[1])
	p2, _ := strconv.Atoi(parts[2])
	p3, _ := strconv.Atoi(parts[3])
	return (p0 << 24) + (p1 << 16) + (p2 << 8) + p3
}

// countIPAddresses calculates the number of IP addresses between two IPv4 addresses.
func countIPAddresses(start, end string) int {
	startInt := ipToInt(start)
	endInt := ipToInt(end)
	return endInt - startInt
}

func main() {
	entradaStart := "10.0.0.0"
	entradaEnd := "10.0.0.50"
	saida := countIPAddresses(entradaStart, entradaEnd)
	fmt.Println(saida) // Output: 50

	entradaStart2 := "10.0.0.0"
	entradaEnd2 := "10.0.1.0"
	saida2 := countIPAddresses(entradaStart2, entradaEnd2)
	fmt.Println(saida2) // Output: 256

	entradaStart3 := "20.0.0.10"
	entradaEnd3 := "20.0.1.0"
	saida3 := countIPAddresses(entradaStart3, entradaEnd3)
	fmt.Println(saida3) // Output: 246
}
