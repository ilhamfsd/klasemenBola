package main

import "fmt"

func cetakTim(T tabTim, nt, np int) {
	var x string

	if CekTimKosong(T) {
		fmt.Println("Anda belum membuat tim.\n Silahkan buat tim terlebih dahulu!")
	} else {
		menuCetakTim()
		fmt.Scan(&x)
		for !cekPilihMenuCetak(x) {
			menuCetakTim()
			fmt.Println("Masukkan inputan yang valid!")
			fmt.Scan(&x)
		}

		if x == "1" {
			cetakTertinggi(T, nt, np)
		} else if x == "2" {
			cetakTerendah(T, nt, np)
		} else if x == "3" {
			cetakSearch(T, nt, np)
		} else if x == "4" {
			cetakDescending(T, nt, np)
		} else if x == "5" {
			cetakAscending(T, nt, np)
		}
	}
}
