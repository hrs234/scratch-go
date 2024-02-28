package main

import (
	"fmt"
	"math"
)

func main() {
	// Heroes selector
	// the given data of heroes is dataHero
	// where initialStartSelector is initial position of cursor when user open the screen, which is []string{row, column}
	// user selection is an inputed from user for choosing the hero which up is change an row to upper, down change row to down, left is change column to left, and right is change column to right

	dataHero := [3][4]string{
		{"A", "B", "C", "D"},
		{"E", "F", "G", "H"},
		{"I", "J", "K", "L"},
	}
	initialStartSelector := [2]int{1, 2}
	userSelection := []string{"up", "up", "down", "left", "right"}

	var up, down, left, right int
	// count user input
	for _, userSelectEle := range userSelection {
		switch userSelectEle {
		case "up":
			up++
			break
		case "down":
			down++
			break
		case "left":
			left++
			break
		case "right":
			right++
			break
		}
	}

	// horizontal move calculation (up and down) when up is decrease when down is increase
	initialStartSelector[0] = initialStartSelector[0] - up
	if math.Signbit(float64(initialStartSelector[0])) {
		initialStartSelector[0] = 0
	}

	initialStartSelector[0] = initialStartSelector[0] + down
	if initialStartSelector[0] >= len(dataHero)-1 {
		initialStartSelector[0] = len(dataHero) - 1
	}

	// vertical move calculation (left and right) when left is decrease when right is increase
	initialStartSelector[1] = initialStartSelector[1] - left
	if math.Signbit(float64(initialStartSelector[1])) {
		initialStartSelector[1] = 0
	}

	initialStartSelector[1] = initialStartSelector[1] + right
	if initialStartSelector[1] >= len(dataHero[initialStartSelector[0]]) {
		initialStartSelector[1] = len(dataHero[initialStartSelector[0]])
	}

	// render selection
	// fmt.Println(fmt.Sprintf("up: %d, down: %d, left: %d, right: %d", up, down, left, right))
	// fmt.Println(initialStartSelector)
	fmt.Println("Selected Hero is : ", dataHero[initialStartSelector[0]][initialStartSelector[1]])
}
