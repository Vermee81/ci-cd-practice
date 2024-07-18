package main

import "fmt"

var version string

func EvenOrOdd(number int) string {
	if number%2 == 0 {
		return "even"
	} else {
		return "odd"
	}
}

func main() {
	fmt.Printf("2 is %s number\n", EvenOrOdd(2))
	fmt.Printf("3 is %s number\n", EvenOrOdd(3))
    fmt.Printf("version %s \n", version)
}
