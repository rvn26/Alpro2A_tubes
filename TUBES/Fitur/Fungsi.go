package Fitur

import (
	"fmt"
	"TUBES/Dataset"
)

func tambahBarang() {
	var id, stok, jum int
	var nama, kategori string
	var harga float64
	fmt.Print("Masukan jumlah barang yang akan ditambahkan : ")
	fmt.Scan(&jum)
	for i := 0; i < jum; i++ {
		fmt.Print("ID Barang: ")
		fmt.Scan(&id)
		fmt.Print("Nama Barang: ")
		fmt.Scan(&nama)
		fmt.Print("Stok Barang: ")
		fmt.Scan(&stok)
		fmt.Print("Harga Barang: ")
		fmt.Scan(&harga)
		fmt.Print("Kategori Barang: ")
		fmt.Scan(&kategori)

		Dataset.BarangList = append(Dataset.BarangList, Dataset.Barang{
			ID:       id,
			Nama:     nama,
			Stok:     stok,
			Harga:    harga,
			Kategori: kategori,
		})
	}

	fmt.Println("Barang berhasil ditambahkan.")
}

func ubahBarang() {
	var id int
	fmt.Print("Masukkan ID Barang yang akan diubah: ")
	fmt.Scan(&id)

	for i, barang := range Dataset.BarangList {
		if barang.ID == id {
			fmt.Print("Nama Barang (sebelumnya: ", barang.Nama, "): ")
			fmt.Scan(&Dataset.BarangList[i].Nama)
			fmt.Print("Stok Barang (sebelumnya: ", barang.Stok, "): ")
			fmt.Scan(&Dataset.BarangList[i].Stok)
			fmt.Print("Harga Barang (sebelumnya: ", barang.Harga, "): ")
			fmt.Scan(&Dataset.BarangList[i].Harga)
			fmt.Print("Kategori Barang (sebelumnya: ", barang.Kategori, "): ")
			fmt.Scan(&Dataset.BarangList[i].Kategori)
			fmt.Println("Barang berhasil diubah.")
			return
		}
	}

	fmt.Println("Barang dengan ID tersebut tidak ditemukan.")
}

func hapusBarang() {
	var id int
	fmt.Print("Masukkan ID Barang yang akan dihapus: ")
	fmt.Scan(&id)

	for i, barang := range Dataset.BarangList {
		if barang.ID == id {
			Dataset.BarangList = append(Dataset.BarangList[:i], Dataset.BarangList[i+1:]...)
			fmt.Println("Barang berhasil dihapus.")
			return
		}
	}

	fmt.Println("Barang dengan ID tersebut tidak ditemukan.")
}
func tampilkanBarang() {
	var pilihan int
	fmt.Println("1. Tampilkan semua barang")
	fmt.Println("2. Tampilkan barang terurut berdasarkan stok")
	fmt.Print("Pilih opsi: ")
	fmt.Scan(&pilihan)

	switch pilihan {
	case 1:
		for _, barang := range Dataset.BarangList {
			fmt.Printf("ID: %d, Nama: %s, Stok: %d, Harga: %.2f, Kategori: %s\n",
				barang.ID, barang.Nama, barang.Stok, barang.Harga, barang.Kategori)
		}
	case 2:
		for i := 0; i < len(Dataset.BarangList); i++ {
			for j := i + 1; j < len(Dataset.BarangList); j++ {
				if Dataset.BarangList[i].Stok > Dataset.BarangList[j].Stok {
					Dataset.BarangList[i], Dataset.BarangList[j] = Dataset.BarangList[j], Dataset.BarangList[i]
				}
			}
		}
		fmt.Println("Barang terurut berdasarkan stok:")
		for _, barang := range Dataset.BarangList {
			fmt.Printf("ID: %d, Nama: %s, Stok: %d, Harga: %.2f, Kategori: %s\n",
				barang.ID, barang.Nama, barang.Stok, barang.Harga, barang.Kategori)
		}
	default:
		fmt.Println("Pilihan tidak valid.")
	}
}
