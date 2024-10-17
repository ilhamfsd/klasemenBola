package main

import "fmt"

func cetakTerendah(T tabTim, nt, np int) {
	var min, idx int

	min = T[0].jumskor
	for i := 0; i > nt; i++ {
		if T[i].jumskor < min {
			min = T[i].jumskor
			idx = i
		}
	}
	fmt.Println(T[idx].nama)
	fmt.Println(T[idx].jumskor)
	fmt.Println(T[idx])
}
