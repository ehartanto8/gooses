package main

import (
	"fmt"
	"time"
)

func printCount(c chan int) {
	num := 0

	for num >= 0 {
		num = <-c // Input
		fmt.Print(num, " ")
	}
}

// Ada 3: num, c, v. Num dapet input dari c. C dapet input dari v
func main() {
	c := make(chan int)
	a := []int{8, 6, 7, 5, 3, 0, 9, -1}
	
	go printCount(c) // Start routine
	
	for _, v := range a {
		c <- v // Inputting v to c
	}

	time.Sleep(time.Millisecond * 1)
	fmt.Println("End of main")
}