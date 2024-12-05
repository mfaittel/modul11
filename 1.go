package main

import "fmt"

func main() {
	jumlahsuara := 0
	suaravalid := 0

	var nomor [21]int

	fmt.Println("0 = selesai")

	for {
		var pilih int
		n, _ := fmt.Scan(&pilih)

		if n == 0 {
			continue
		}

		if pilih == 0 {
			break
		}

		jumlahsuara++

		if pilih >= 1 && pilih <= 20 {
			suaravalid++
			nomor[pilih]++
		}
	}

	fmt.Printf("Suara masuk: %d\n", jumlahsuara)
	fmt.Printf("Suara sah: %d\n", suaravalid)

	for i := 1; i <= 20; i++ {
		if nomor[i] > 0 {
			fmt.Printf("%d:%d\n", i, nomor[i])
		}
	}
}
