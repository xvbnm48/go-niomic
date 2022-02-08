package main

import "fmt"

func main() {
	// a := 1
	// b := 2

	// if a == b {
	// 	println("a == b")
	// } else {
	// 	println("a tidak sama dengan b")
	// }

	nilai := 70
	switch nilai {
		case 50:
			fmt.Println("Nilai anda di bawah rata rata")
		case 60:
			fmt.Println("Nilai anda di rata rata")
		case 70:
			fmt.Println("Nilai anda di atas rata rata")
		case 80:
			fmt.Println("Nilai anda sangat bagus")	
	}

}