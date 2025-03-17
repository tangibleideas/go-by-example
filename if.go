package main

import "fmt"

func main() {
	// if statement
	if 9%3 == 0 {
		fmt.Println("9 is divisible by 3")
	}

	// if-else statement
	if 9%4 == 0 {
		fmt.Println("9 is divisible by 4")
	} else {
		fmt.Println("9 is not divisible by 4")
	}

	// else-if statement
	// statement can precede conditions; any variables declared in this statement are available in all branches
	if num := 9; num < 0 {
		fmt.Println(num, "is negative")
	} else if num < 10 {
		fmt.Println(num, "has 1 digit")
	} else {
		fmt.Println(num, "has multiple digits")
	}

	//parentheses are not required around conditions in Go
	// only braces are required
	// there is no ternary if in Go

}
