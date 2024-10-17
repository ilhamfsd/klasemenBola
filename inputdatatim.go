package main

import "fmt"

func inputTim(T *tabTim) {

	var nama string
	var i, idx int
	fmt.Scan(&nama)
	idx = -1
	for idx == -1 && i < maxTim {
		if nama == T[i].nama {
			idx = i
		}
		i++
	}

	if idx == -1 {
		fmt.Println("nama tim tidak ditemukan")
	} else {
		fmt.Println("masukkan jumlah gol yang dicetak: ")
		fmt.Scan(&T[idx].jumGol)
		fmt.Println("masukkan jumlah kebobolan: ")
		fmt.Scan(&T[idx].jumKebobolan)
		fmt.Println("masukkan jumlah menang: ")
		fmt.Scan(&T[idx].jumKemenangan)
		fmt.Println("masukkan jumlah kalah: ")
		fmt.Scan(&T[idx].jumKekalahan)
		fmt.Println("masukkan jumlah seri: ")
		fmt.Scan(&T[idx].jumSeri)
		fmt.Println("Masukkan data pemain: ")
		T[idx].jumskor = (T[idx].jumKemenangan * 3) + (T[idx].jumSeri * 1)
	}
}
