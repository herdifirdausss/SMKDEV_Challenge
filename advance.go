package main

import (
	"fmt"
)

// buat function untuk menuliskan bilangan prima antara input awal dan akhir
func bilanganprima(awal, akhir int) {
	//iniialisasi variabel
	var bilangan int
	var prima bool
	//perulangan untuk mengecek bilangan prima dari input awal dan akhir
	for bilangan = awal; bilangan <= akhir; bilangan++ {
		//inisialisasi boolean, jika tidak bisa dibagi i make true (default)
		prima = true
		//perulangan untuk membagi bilangan sampai i/2, jika bisa dibagi maka bukan bilangan prima (false)
		for i := 2; i < bilangan/2; i++ {
			if bilangan%i == 0 {
				prima = false
				break
			}
		}
		//tuliskan bilangan tersebut jika prime = true
		if prima {
			fmt.Printf("%v\n", bilangan)
		}
	}
}

func main() {
	//inisialisasi variabel
	var awal, akhir int
	fmt.Println("# range")
	//semua input dianggap angka, sehingga tidak diperlukan validasi
	//memasukkan input
	fmt.Print("start = ")
	fmt.Scan(&awal)
	fmt.Print("end = ")
	fmt.Scan(&akhir)
	//menuliskan keterangan
	fmt.Printf("Prime numbers between %v and %v are :\n", awal, akhir)
	//menuliskan hasil
	bilanganprima(awal, akhir)
	//Test1
	//start = 25, end = 50
	// hasil : 29,31,37,41,43,47
}
