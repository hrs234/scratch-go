package main

import "fmt"

func main() {
	text := "abcd"
	if len(text) >= 4 {
		fmt.Printf(reverse(text))
	} else {
		fmt.Printf(text)
	}
}

func reverse(t string) (result string) {
	for _, v := range t {
		result = string(v) + result
	}
	return
}
