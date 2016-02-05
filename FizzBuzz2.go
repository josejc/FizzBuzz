package main

import "fmt"

func main() {
	for i:=1; i<=100; i++ {
		switch i {
			case (i%3 == 0) && (i%5 == 0):
				fmt.Println("FizzBuzz")
		 	case (i%3 == 0):
		 		fmt.Println("Fizz")
		 	case (i%5 == 0):
		 		fmt.Println("Buzz")
			default:
				fmt.Println(i)
		}
	}
}

// ERROR
prog.go:8: invalid case i % 3 == 0 && i % 5 == 0 in switch on i (mismatched types bool and int)
prog.go:10: invalid case i % 3 == 0 in switch on i (mismatched types bool and int)
prog.go:12: invalid case i % 5 == 0 in switch on i (mismatched types bool and int)