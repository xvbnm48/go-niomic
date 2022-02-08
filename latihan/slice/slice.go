package main

import "fmt"

func main() {
	var kendaraan = []string{"mobil", "motor", "sepeda", "bus"}
	kendaraan = append(kendaraan, "truk")
	fmt.Println(kendaraan[2])
	fmt.Println(len(kendaraan))
	fmt.Println(kendaraan)

}