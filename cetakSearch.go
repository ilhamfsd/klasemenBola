package main

import "fmt"

func cetakSearch(T tabTim, nt, np int) {

	var nama string
	var i, idx int
	fmt.Scan(&nama)
	idx = -1
	for idx == -1 && i < nt {
		if nama == T[i].nama {
			idx = i
		}
		i++
	}

	if idx == -1 {
		fmt.Println("Tidak ada data yang cocok dengan pencarian anda")
	} else {
		fmt.Printf("Data Tim %s: ", T[idx].nama)
		fmt.Println("jumlah gol yang dicetak: ", T[idx].jumGol)
		fmt.Println("jumlah kebobolan: ", T[idx].jumKebobolan)
		fmt.Println("jumlah menang: ", T[idx].jumKemenangan)
		fmt.Println("jumlah kalah: ", T[idx].jumKekalahan)
		fmt.Println("jumlah seri: ", T[idx].jumSeri)
	}
}
