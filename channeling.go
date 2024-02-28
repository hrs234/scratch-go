package main

import (
	"fmt"
	"runtime"
)

func main() {
	runtime.GOMAXPROCS(2)

	// init the channel
	var msg = make(chan string)
	var saySomething = func(say string) {
		// the value delivered into msg channel variable
		msg <- fmt.Sprintf("Hi, %s", say)
	}

	// make a function run asynchronus with goroutine in go keyword
	go saySomething("Budi")
	go saySomething("Sunima")
	go saySomething("jaja")

	// data received from msg channel asynchrounusly
	var msgA = <-msg
	fmt.Println(msgA)

	var msgB = <-msg
	fmt.Println(msgB)

	var msgC = <-msg
	fmt.Println(msgC)

}
