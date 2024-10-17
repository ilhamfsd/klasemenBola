package main

import "fmt"

func createTim(T *tabTim, nT *int) {
		fmt.Println("masukkan nama tim yang ingin anda buat: ")
		fmt.Scan(&T[*nT].nama)
		*nT += 1
}

/*
type tim struct {
	nama                                                       string
	jumKemenangan, jumKekalahan, jumSeri, jumGol, jumKebobolan int
	pemain                                                     tabPemain
}

type pemain struct {
	nama, posisi, np string
}
*/