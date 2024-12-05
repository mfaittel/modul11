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

	ketua := 0
	wakil := 0
	for i := 1; i <= 20; i++ {
		if nomor[i] > nomor[ketua] || (nomor[i] == nomor[ketua] && ketua > i) {
			wakil = ketua
			ketua = i
		} else if nomor[i] > nomor[wakil] || (nomor[i] == nomor[wakil] && wakil > i) {
			wakil = i
		}
	}

	fmt.Printf("Suara masuk: %d\n", jumlahsuara)
	fmt.Printf("Suara sah: %d\n", suaravalid)

	if suaravalid > 0 {
		fmt.Printf("Ketua RT: %d\n", ketua)
		if wakil != 0 {
			fmt.Printf("Wakil Ketua RT: %d\n", wakil)
		}
	}
}
