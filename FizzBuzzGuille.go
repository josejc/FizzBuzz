package main

import "fmt"

func main() {
	var cont int
	
	for i:=1; i<=100; i++ {
		cont = 0	
		if (i%3 == 0) {
			cont = 1
		 	fmt.Printf("Fizz")
		}
		if (i%5 == 0) {
			cont = 1
		 	fmt.Printf("Buzz")
		}
		if (cont == 0) {
			fmt.Printf("%d",i)
		}
		fmt.Printf("\n")
	}
}