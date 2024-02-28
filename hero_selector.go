package main

import (
	"fmt"
)

func renderSelection(initialStartSelector []int, dataHero [][]string) string {
	hero := ""

	for i := initialStartSelector[0]; i < len(dataHero); i++ {
		for n := initialStartSelector[1]; n < len(dataHero[i]); n++ {
			fmt.Println(dataHero[i][n])
			hero = dataHero[i][n]
		}
	}

	return hero
}

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
	userSelection := []string{"up", "up", "down"}
	finalSelectedRes := ""

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

	// render selection

	fmt.Println(dataHero)
	fmt.Println(initialStartSelector)
	fmt.Println(fmt.Sprintf("up: %d, down: %d, left: %d, right: %d", up, down, left, right))
	fmt.Println("Selected Hero is : ", finalSelectedRes)
}
