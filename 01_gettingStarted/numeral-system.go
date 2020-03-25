package main

import "fmt"

func main() {
	/*	fmt.Printf("%d - %b \n", 42, 42)
		fmt.Printf("%d \t %b \t %#X \n", 42, 42, 42)*/
	for i := 65; i < 123; i++ {
		fmt.Printf("%03d \t %b \t %x \t %q \n", i, i, i, i)
	}
}
