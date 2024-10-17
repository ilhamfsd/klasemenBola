package main

import "fmt"

func cetakAscending(T tabTim, nt, np int) {
	var ascending tabTim
	var temp tim
	var i, idxMin, pass int
	for pass < nt {
		idxMin = pass
		i = pass + 1
		for i < nt {
			if T[idxMin].jumskor > T[i].jumskor {
				idxMin = i
			}
			i++

		}
		temp = T[pass]
		ascending[pass] = T[idxMin]
		T[idxMin] = temp
		pass++
	}
	for j := 0; j < nt; j++ {
		fmt.Println(ascending[j])
	}
}
