package main

import "errors"
import "fmt"

func f1(arg int) (int, error) {
	if arg == 42 {
		return -1, errors.New("Cna't work with 42")
	}
	return arg + 3, nil
}

type argError struct {
	arg int
	prob string
}

func (e *argError) Error() string {
	return fmt.Sprintf("%d - %s", e.arg, e.prob)
}

func f2(arg int) (int, error) {
	if arg == 42 {
		return -1, &argError{arg, "Can't work with it"}
	}
	return arg + 3, nil
}

func main() {
	for _, i := 
}
