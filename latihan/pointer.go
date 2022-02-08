package main

import "fmt"

func main() {
	nomor := 10
	var alamat_nomor *int = &nomor

	fmt.Println("nilai nomor", nomor)
	fmt.Println("alamat nomor", alamat_nomor)
}
