package main

import "fmt"

func main() {
	var i int
	//for loop
	for i = 1; i <= 10; i++ {
		println("sakura endo")
	}
	// while
	a := 0
	for a < 10 {
		a++
		println("sakura miyawaki")
	}

	x := 0
	for {
		fmt.Println("angka: " , x)
		x++
		if x == 10 {
			break
		}
	}
}