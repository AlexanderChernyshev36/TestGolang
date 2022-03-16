package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan int)
    go say("hello world",ch)
	fmt.Println("1")
	fmt.Println("1")
	fmt.Println("1")
	fmt.Println("1")
	fmt.Println("1")
	fmt.Println(<-ch)
}

func say(word string, ch chan int) {
	time.Sleep(2 * time.Second)
	fmt.Println(word)
	ch <- 9
}