package main

import "fmt"

func main() {
	var buah = []string{"apel", "jeruk", "anggur", "melon"}
	buah = append(buah, "pepaya")
	fmt.Println("jumlah Element: " , len(buah))
	for i := 0; i < len(buah); i++ {
		fmt.Println("Element ke- : ", i, " ", buah[i])
	}
}