package main

import (
	"fmt"
)

func numberCanDevided(limit int, sharedNumA int, sharedNumB int) []int {
	selectedA := []int{}
	selectedB := []int{}
	selected := []int{}
	for a := 1; a <= limit; a++ {
		if sharedNumA%a == 0 {
			selectedA = append(selectedA, a)
		}
	}

	for b := 1; b <= limit; b++ {
		if sharedNumB%b == 0 {
			selectedB = append(selectedB, b)
		}
	}

	for _, c := range selectedA {
		for _, d := range selectedB {
			if c == d {
				selected = append(selected, c)
			}
		}
	}

	return selected
}

func main() {
	n := 24
	m := 12
	limit := m

	if m <= n {
		limit = n
	}

	fmt.Println(len(numberCanDevided(limit, n, m)))
}
