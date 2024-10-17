package main

import "fmt"

func cetakTertinggi(T tabTim, nt, np int) {
	var maks, idx int

	for i := 0; i < nt; i++ {
		if T[i].jumskor > maks {
			maks = T[i].jumskor
			idx = i
		}
	}
	fmt.Println(T[idx].nama)
	fmt.Println(T[idx].jumskor)
	fmt.Println(T[idx])
}