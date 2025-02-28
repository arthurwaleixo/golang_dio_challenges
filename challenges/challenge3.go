package main

import (
	"fmt"
	"time"
)

func goPing(word chan string, done chan bool) {
	for {
		select {
		case word <- "ping":
		case <-done:
			return
		}
	}
}

func goPong(word chan string, done chan bool) {
	for {
		select {
		case word <- "pong":
		case <-done:
			return
		}
	}
}

func printWord(word chan string, done chan bool) {
	for {
		select {
		case message := <-word:
			fmt.Println(message)
			time.Sleep(1 * time.Second)
		case <-done:
			return
		}
	}
}

func main() {
	word := make(chan string)
	done := make(chan bool)

	go goPing(word, done)
	go goPong(word, done)
	go printWord(word, done)

	fmt.Println("Press Enter to exit")
	fmt.Scanln()

	close(done)

	time.Sleep(1 * time.Second)
	fmt.Println("Exiting...")
}
