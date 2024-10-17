package main

import "fmt"

func cetakDescending(T tabTim, nt, np int) {
	var descending tabTim
	var temp tim	
	var i, idxMaks, pass int
	for pass < nt {
		idxMaks = pass
		i = pass + 1
		for i < nt {
			if T[idxMaks].jumskor < T[i].jumskor {
				idxMaks = i
			}
			i++
		}
		temp = T[pass]
		descending[pass] = T[idxMaks]
		T[idxMaks] = temp
		pass++
	}
	for j := 0; j < nt; j++ {
		fmt.Println(descending[j])
	}

}
