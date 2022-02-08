package main

import "fmt"

func main() {
	var s1 pelajar
	var s2 = pelajar{"sakura", 18}
	s1.nama = "aldo"
	s1.umur = 19
	s1.namasaya()
	s2.namasaya()
}

type pelajar struct {
	nama string
	umur int
}

func (s pelajar) namasaya() {
	fmt.Println("nama saya : ", s.nama)

}
