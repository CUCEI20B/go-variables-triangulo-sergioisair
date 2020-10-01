package main

import "fmt"

func main() {
	var a uint64
	var b uint64
	fmt.Scan(&a)
	fmt.Scan(&b)
	output := (a * b) / 2
	fmt.Print(output)
}
