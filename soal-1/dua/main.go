package main

import (
	"fmt"
	"math"
)

func calculateChange(total, bayar int) (bool, int, int, map[int]int) {
	if bayar < total {
		return false, 0, 0, nil
	}

	kembalian := bayar - total

	roundedChange := int(math.Floor(float64(kembalian)/100) * 100)

	jenisUang := []int{100000, 50000, 20000, 10000, 5000, 2000, 1000, 500, 200, 100}
	pecahan := make(map[int]int)

	for _, jenis := range jenisUang {
		if roundedChange >= jenis {
			pecahan[jenis] = roundedChange / jenis
			roundedChange = roundedChange % jenis
		}
	}

	return true, kembalian, int(math.Floor(float64(kembalian)/100) * 100), pecahan
}

func main() {
	var totalBelanja, bayar int
	fmt.Print("Total belanja: Rp ")
	fmt.Scan(&totalBelanja)
	fmt.Print("Uang yang dibayarkan: Rp ")
	fmt.Scan(&bayar)

	valid, kembalian, roundedChange, pecahan := calculateChange(totalBelanja, bayar)

	if !valid {
		fmt.Println("False, kurang bayar")
		return
	}

	fmt.Printf("Kembalian yang harus diberikan kasir: %d,\ndibulatkan menjadi %d\n", kembalian, roundedChange)
	fmt.Println("Pecahan uang:")
	for denom, count := range pecahan {
		if denom >= 1000 {
			fmt.Printf("%d lembar %d\n", count, denom)
		} else {
			fmt.Printf("%d koin %d\n", count, denom)
		}

	}
}
