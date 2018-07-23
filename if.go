package main

import "fmt"

func main() {
	if 7%2 == 0 {
		fmt.Println("7 is even\n")	
	} else {
		fmt.Println("7 is odd\n")
	}

	if 8%4 == 0 {
		fmt.Println("8 is divisible by 4\n")
	}

	if num := 9; num < 0 {
		fmt.Println(num, " is negative\n")
	} else if num < 10 {
		fmt.Println(num, " has 1 digit\n")
	} else {
		fmt.Println(num, " has multiple digits\n")
	}
}
