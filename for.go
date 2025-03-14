package main

import "fmt"

func main() {
	// basic for loop
	i := 1
	for i <= 3 {
		fmt.Println(i)
		i = i + 1
	}

	// classic initial condition after for loop
	for j := 0; j < 3; j++ {
		fmt.Println(j)
	}

	// range over integer
	for i := range 3 {
		fmt.Println(i)
	}

	// break or return to exit loop
	for {
		fmt.Println("infinite loop, break to exit")
		break
	}

	for n := range 6 {
		if n%2 == 0 {
			continue
		}
		fmt.Println(n)
	}
}
