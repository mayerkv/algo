package main

import "fmt"

func main() {
	for x := 0; x < 25; x++ {
		for y := 0; y < 25; y++ {
			writePoint(x, y)
		}
		fmt.Println()
	}
}

func writePoint(x int, y int) {
	fmt.Printf("%s  ", getPoint(x, y))
}

func getPoint(x int, y int) string {
	if x == y {
		return "#"
	}
	return "."
}
