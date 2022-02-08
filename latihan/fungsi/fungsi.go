package main

import "fmt"

func main() {
	var x1, x2 = tambah(2, 3)
	fmt.Println(x1)
	fmt.Println(x2)
}

func tampilkan_pesan() string {
	return "sakura endo sangat cantik sekali"
}

func tambah(x int, y int) (int,int) {
	var z = x * y
	var a = x + y
	return z, a
}
