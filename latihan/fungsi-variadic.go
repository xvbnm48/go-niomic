package main

import "fmt"

func main() {
	var rata_rata = hitung(1, 2, 3, 4, 5, 6, 7, 8, 9, 10)
	var pesan = fmt.Sprintf("Rata-rata: %.2f", rata_rata)
	fmt.Println(pesan)
}

func hitung(angka ...int)float64{
	var total int = 0
	for _, angka := range angka {
		total += angka
		fmt.Println(angka)
	}
	var avg = float64(total) / float64(len(angka))
	return avg
}