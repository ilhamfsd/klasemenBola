package main

import (
	"fmt"
)

func inputPemain(T *tabTim, nT int, nP *int) {
	var nama string
	var i, idx int
	var cek bool
	fmt.Println("masukkan nama tim: ")
	fmt.Scan(&nama)
	idx = -1
	for idx == -1 && i < nT {
		cek = nama == T[i].nama
		if cek {
			idx = i
		}
		i++
	}

	if idx == -1 {
		fmt.Println("data tidak ditemukan")
	} else {

		for j := 0; j < maxPemain; j++ {
			fmt.Printf("pemain ke-%d", j+1)
			fmt.Println("nama: ")
			fmt.Scan(&T[idx].pemain[j].nama)
			fmt.Println("no. punggung: ")
			fmt.Scan(&T[idx].pemain[j].np)
			fmt.Println("posisi: ")
			fmt.Scan(&T[idx].pemain[j].posisi)
			*nP++
		}
	}
}
