package main

import "fmt"

func main() {
	var harga_figure = map[string]int{"nico_yazawa": 100000, "ninynm_figure": 120000, "tsukasa_figure": 150000}
	fmt.Println("harga figure nico yazawa : ", harga_figure["nico_yazawa"])
}