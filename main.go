package main

import (
	"fmt"
)

func main() {
	drawItAll(10)
}

func drawItAll(n int) {
	height, width := n, ((2 * n) - 1)
	for i := 0; i < height; i++ {
		drawRow(i, width)
	}
}

func drawRow(i, width int) {
	startDots := (width / 2) - i
	endDots := (width / 2) + i
	for i := 0; i < width; i++ {
		if i >= startDots && i <= endDots {
			fmt.Print("*")
		} else {
			fmt.Print("-")
		}
	}
	fmt.Print("\n")
}
