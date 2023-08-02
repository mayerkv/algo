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

// 02. x == y
// 08. x == 0 || y == 0
// 11. x == 1 || y == 1 || x == 23 || y == 23
// 19. x == 0 || y == 0 || x == 24 || y == 24
// 24. x == y || x == 24-y
// 07. x > 15 && y > 15
// 04. x+y < 30
// 06. x < 10 || y < 10
func getPoint(x int, y int) string {
	if x < 10 || y < 10 { // magic
		return "#"
	}
	return "."
}
