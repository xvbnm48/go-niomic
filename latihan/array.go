package main

import "fmt"

func main() {
	var buah = [4]string{"apel", "jeruk", "anggur", "mangga"}
	fmt.Println(buah[3])
	fmt.Println("jumlah buah adalah:" , len(buah))
}