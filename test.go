package main

import (
	"fmt"
	//"sync"
	"time"
)




func A(ch chan int) {
	fmt.Println("produce!")
	ch <- 1
}

func B(ch chan int) {
	var n int
	n = <- ch
	fmt.Println(n)
}

func main() {
	//var wg sync.WaitGroup
	//wg.Add(1)

	ch := make(chan int,1)
	ch <- 10
	go B(ch)
	go A(ch)
	

	time.Sleep(time.Second)
}