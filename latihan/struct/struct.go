package main

import "fmt"

func main() {
	var sakura pelajar
	sakura.nama = "sakura"
	sakura.umur = 18
	sakura.kelas = 10
	fmt.Println("nama sakura : ", sakura.nama)
}

type pelajar struct {
	nama  string
	kelas int
	umur  int
}
