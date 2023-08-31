package main

import (
	"fmt"
)

// buat function untuk menghitung jumlah tip
func tips(tagihan float32) float32 {
	if tagihan >= 50 && tagihan <= 300 {
		return tagihan * 0.15
	} else {
		return tagihan * 0.2
	}
}

func main() {
	//misalkan input tagihan bisa berupa float juga
	var tagihan float32
	fmt.Print("Jumlah Tagihan: ")
	//semua input dianggap angka, sehingga tidak diperlukan validasi
	fmt.Scan(&tagihan)
	//jika tagihan 0 atau negatif, maka
	if tagihan < 1 {
		fmt.Println("Input tidak valid")
		return
	}
	//Jika input sudah valid, maka hitung nilai tip berdasarkan function tips yang sudah dibuat
	tip := tips(tagihan)
	//Tampilkan hasil
	fmt.Printf("Tagihan: %.2f\n", tagihan)
	fmt.Printf("Tip: %.2f\n", tip)
	fmt.Printf("Total Nilai: %.2f", tagihan+tip)
	//Test1
	//tagihan = 275, tip = 41.25, total nilai = 316.25

	//Test2
	//tagihan = 40, tip = 8, total nilai = 48

	//Test3
	//tagihan = 430, tip = 86, total nilai = 516
}
