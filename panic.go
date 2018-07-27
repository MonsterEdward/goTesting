package main

import "os"

func main() {
	panic("A problem")

	_, err := os.Create("./tmp")
	if err != nil {
		panic(err)
	}
}
