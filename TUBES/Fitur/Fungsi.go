package Fitur

import (
	"TUBES/Dataset"
	"fmt"
	"strings"
)

func buatID(total int) int {
	return total + 1
}
func generateIDTransaksi(totalMasuk int) string {
	return fmt.Sprintf("TR%02d", totalMasuk+1)
}
func generateIDTransaksiKeluar(totalKeluar int) string {
	return fmt.Sprintf("TR%02d", totalKeluar+1)
}

func TransaksiBarangMasuk() {
	var totalMasuk = len(Dataset.TransaksiList)
	var id_transaksi string
	var id, jumlah int
	var nama string
	id_transaksi = generateIDTransaksi(totalMasuk)
	fmt.Print("Masukkan ID Barang    : ")
	fmt.Scan(&id)
	fmt.Print("Masukkan Nama Barang  : ")
	fmt.Scan(&nama)
	fmt.Print("Jumlah Barang Masuk   : ")
	fmt.Scan(&jumlah)

	

	for i, barang := range Dataset.BarangList {
		if barang.ID == id && barang.Nama == nama {
			Dataset.BarangList[i].Stok += jumlah
			fmt.Println("Transaksi berhasil.")
			totalMasuk++
			break
		}else{
			fmt.Println("Transaksi gagal Periksa ID barang atau nama barang")
			return
		}
	}
	Dataset.TransaksiList = append(Dataset.TransaksiList, Dataset.Transaksi{
		ID_Transaksi: id_transaksi,
		ID_Barang:    id,
		Nama_barang:  nama,
		Jumlah:       jumlah,
	})
	
	
}

func TransaksiBarangKeluar() {
	
	var totalKeluar = len(Dataset.TransakasiKeluar)
	var id_transaksi string
	var id, jumlah int
	var nama string
	
	id_transaksi = generateIDTransaksiKeluar(totalKeluar)
	fmt.Print("Masukkan ID Barang    : ")
	fmt.Scan(&id)
	fmt.Print("Masukkan Nama Barang  : ")
	fmt.Scan(&nama)
	fmt.Print("Jumlah Barang Masuk   : ")
	fmt.Scan(&jumlah)
	
	
	

	for i, barang := range Dataset.BarangList {
		if barang.ID == id {
			if barang.Stok >= jumlah  && barang.Nama == nama{
				Dataset.BarangList[i].Stok -= jumlah
				fmt.Println("Transaksi berhasil.")
			} else {
				fmt.Println("transakasi gagal (Stok kurang atau nama tidak sesuai dengan ID barang).")
				return
			}
			totalKeluar++
			break
		}else{
			fmt.Println("Barang dengan ID tersebut tidak ditemukan.")
			return
		}
	}
	Dataset.TransakasiKeluar = append(Dataset.TransakasiKeluar, Dataset.Transaksi{
		ID_Transaksi: id_transaksi,
		ID_Barang:    id,
		Nama_barang:  nama,
		Jumlah:       jumlah,
	})
}

func CariBarang() {
	var keyword, kategori string
	fmt.Print("Masukkan kata kunci: ")
	fmt.Scan(&keyword)
	fmt.Print("Masukkan kategori: ")
	fmt.Scan(&kategori)

	fmt.Println("Hasil pencarian:")
	for _, barang := range Dataset.BarangList {
		if (strings.Contains(strings.ToLower(barang.Nama), strings.ToLower(keyword)) ||
			strings.Contains(strings.ToLower(barang.Kategori), strings.ToLower(keyword))) &&
			(kategori == "" || strings.EqualFold(barang.Kategori, kategori)) {
			fmt.Printf("ID: %d, Nama: %s, Stok: %d, Harga: %.2f, Kategori: %s\n",
				barang.ID, barang.Nama, barang.Stok, barang.Harga, barang.Kategori)
		}
	}
}
func CariNamaBarang() {
	var keyword string
	fmt.Print("Masukkan Nama Barang: ")
	fmt.Scan(&keyword)

	fmt.Println("Hasil pencarian:")
	for _, barang := range Dataset.BarangList {
		if (strings.Contains(strings.ToLower(barang.Nama), strings.ToLower(keyword)) ||
			strings.Contains(strings.ToLower(barang.Kategori), strings.ToLower(keyword))) {
			fmt.Printf("ID: %d, Nama: %s, Stok: %d, Harga: %.2f, Kategori: %s\n",
				barang.ID, barang.Nama, barang.Stok, barang.Harga, barang.Kategori)
		}
	}
}
func TambahBarang() {
	var total = len(Dataset.BarangList)
	var id, stok, jum int
	var nama, kategori string
	var harga float64
	fmt.Print("Masukan jumlah barang yang akan ditambahkan : ")
	fmt.Scan(&jum)
	for i := 0; i < jum; i++ {
		id = buatID(total)
		fmt.Print("Nama Barang		: ")
		fmt.Scan(&nama)
		fmt.Print("Stok Barang		: ")
		fmt.Scan(&stok)
		fmt.Print("Harga Barang		: ")
		fmt.Scan(&harga)
		fmt.Print("Kategori Barang		: ")
		fmt.Scan(&kategori)

		Dataset.BarangList = append(Dataset.BarangList, Dataset.Barang{
			ID:       id,
			Nama:     nama,
			Stok:     stok,
			Harga:    harga,
			Kategori: kategori,
		})
		total++
	}

	fmt.Println("Barang berhasil ditambahkan.")
}

func UbahBarang() {
	var id int
	fmt.Print("Masukkan ID Barang yang akan diubah: ")
	fmt.Scan(&id)

	for i, barang := range Dataset.BarangList {
		if barang.ID == id {
			fmt.Print("Nama Barang (sebelumnya: ", barang.Nama, ")		 : ")
			fmt.Scan(&Dataset.BarangList[i].Nama)
			fmt.Print("Stok Barang (sebelumnya: ", barang.Stok, ")		 : ")
			fmt.Scan(&Dataset.BarangList[i].Stok)
			fmt.Print("Harga Barang (sebelumnya: ", barang.Harga, ")	 : ")
			fmt.Scan(&Dataset.BarangList[i].Harga)
			fmt.Print("Kategori Barang (sebelumnya: ", barang.Kategori, "): ")
			fmt.Scan(&Dataset.BarangList[i].Kategori)
			fmt.Println("Barang berhasil diubah.")
			return
		}
	}

	fmt.Println("Barang dengan ID tersebut tidak ditemukan.")
}

func HapusBarang() {
	var total = len(Dataset.BarangList)
	var id int
	fmt.Print("Masukkan ID Barang yang akan dihapus: ")
	fmt.Scan(&id)
	found := false
	for i, barang := range Dataset.BarangList {
		if barang.ID == id {
			Dataset.BarangList = append(Dataset.BarangList[:i], Dataset.BarangList[i+1:]...)
			found = true
			break
		}
	}
	if found {
		for i := id - 1; i < len(Dataset.BarangList); i++ {
			Dataset.BarangList[i].ID--
			
		}
		total--
		fmt.Println("Barang berhasil dihapus.")
	}else{
		fmt.Println("Barang dengan ID tersebut tidak ditemukan.")
	}
}

func TampilkanBarang() {
	for{
	var pilihan int
	fmt.Println("----------------------------------")
	fmt.Println("| 1. Tampilkan semua barang      |")
	fmt.Println("| 2. Tampilkan berdasarkan stok  |")
	fmt.Println("| 3. Tampilkan Transaksi Masuk   |")
	fmt.Println("| 4. Tampilkan Transaksi Keluar  |")
	fmt.Println("| 5. Kembali                     |")
	fmt.Println("----------------------------------")
	fmt.Print("Pilih opsi: ")
	fmt.Scan(&pilihan)
	if pilihan == 5 {
		break
	}
	switch pilihan {
	case 1:
		fmt.Println("---------------------------------------------------------------")
		fmt.Printf("|%-5s|%-20s|%-10s|%-10s|%-20s\n", "ID", "Nama", "Stok", "Harga", "Kategori")
		fmt.Println("---------------------------------------------------------------")
		for _, barang := range Dataset.BarangList {
			fmt.Printf("|%-5d|%-20s|%-10d|%-10.2f|%-20s\n",
				barang.ID, barang.Nama, barang.Stok, barang.Harga, barang.Kategori)
			fmt.Println("---------------------------------------------------------------")
		}

	case 2:
		fmt.Println("Barang terurut berdasarkan stok:")
		fmt.Println("---------------------------------------------------------------")
		fmt.Printf("|%-5s|%-20s|%-10s|%-10s|%-20s\n", "ID", "Nama", "Stok", "Harga", "Kategori")
		fmt.Println("---------------------------------------------------------------")
		for i := 0; i < len(Dataset.BarangList); i++ {
			for j := i + 1; j < len(Dataset.BarangList); j++ {
				if Dataset.BarangList[i].Stok > Dataset.BarangList[j].Stok {
					Dataset.BarangList[i], Dataset.BarangList[j] = Dataset.BarangList[j], Dataset.BarangList[i]
				}
			}
		}
		
		for _, barang := range Dataset.BarangList {
			fmt.Printf("|%-5d|%-20s|%-10d|%-10.2f|%-20s\n",
				barang.ID, barang.Nama, barang.Stok, barang.Harga, barang.Kategori)
		}
	case 3:
		fmt.Println("---------------------------------------------------------------")
		fmt.Printf("|%-15s|%-15s|%-20s|%-10s\n", "ID Transaksi", "ID Barang", "Nama Barang", "Jumlah")
		fmt.Println("---------------------------------------------------------------")
		for _, transakasi := range Dataset.TransaksiList {
			fmt.Printf("|%-15s|%-15d|%-20s|%-10d\n",
				transakasi.ID_Transaksi, transakasi.ID_Barang, transakasi.Nama_barang, transakasi.Jumlah)
		}
	case 4:
		fmt.Println("---------------------------------------------------------------")
		fmt.Printf("|%-15s|%-15s|%-20s|%-10s\n", "ID Transaksi", "ID Barang", "Nama Barang", "Jumlah")
		fmt.Println("---------------------------------------------------------------")
		for _, transakasi := range Dataset.TransakasiKeluar {
			fmt.Printf("|%-15s|%-15d|%-20s|%-10d\n",
				transakasi.ID_Transaksi, transakasi.ID_Barang, transakasi.Nama_barang, transakasi.Jumlah)
		}
	default:
		fmt.Println("Pilihan tidak valid.")
	}
}}

func UrutID(){
	for i := 0; i < len(Dataset.BarangList); i++ {
		for j := i + 1; j < len(Dataset.BarangList); j++ {
			if Dataset.BarangList[i].ID > Dataset.BarangList[j].ID {
				Dataset.BarangList[i], Dataset.BarangList[j] = Dataset.BarangList[j], Dataset.BarangList[i]
			}
		}
	}
}
