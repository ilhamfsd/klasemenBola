package main

import (
	"fmt"
)

func pilihMenu(T *tabTim, nP, nT *int) {
	var x string
	var A tabTim
	var nTim, nPemain int
	fmt.Scan(&x)

	for x != "0" {
	for !cekPilihMenu(x) {
		fmt.Println("input tidak valid\nsilahkan masukkan input yang valid!")
		fmt.Scan(&x)
	}
	if x == "1" {
		createTim(&A, &nTim)
		*T = A
		*nT = nTim
 		menu()
		fmt.Scan(&x)
	} else if x == "2" {
		inputTim(&A)
		*T = A
		fmt.Print(T)
		menu()
		fmt.Scan(&x)
	} else if x == "3" {
		inputPemain(&A, nTim, &nPemain)
		*T = A
		*nP = nPemain
		menu()
		fmt.Scan(&x)
	} else if x == "4" {
		cetakTim(*T, *nT, *nP)
		menu()
		fmt.Scan(&x)
	} 
}
}