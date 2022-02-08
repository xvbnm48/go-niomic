package main

import "fmt"

func main() {
	var mahasiswa = map[string]int{"aldo": 182, "yosep": 178}
	tampil_mahasiswa(mahasiswa["aldo"], mahasiswa["yosep"])
}

func tampil_mahasiswa(x int, a int) (int, int) {
	fmt.Println("tinggi aldo : ", x, "cm")
	fmt.Println("tinggi yosep : ", a, "cm")
	return x, a
}