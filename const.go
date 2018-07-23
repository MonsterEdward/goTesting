package main

import "fmt"
import "math"

const s string = "constant"

func main() {
	fmt.Println(s)

	const n = 5e10
	const m = 2e20 / n
	fmt.Println(m)
	
	fmt.Println(int64(m))

	fmt.Println(math.Sin(n))
}
