package main

import (
	"TUBES/Dataset"
	"bufio"
	"encoding/csv"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	
	loadData()
	loadDataTransaksi()
	loadDataTransaksiKeluar()

	
	for {
		fmt.Println("\nAplikasi Inventori Barang")
		fmt.Println("1. Tambah Barang")
		fmt.Println("2. Ubah Barang")
		fmt.Println("3. Hapus Barang")
		fmt.Println("4. Transaksi ")
		fmt.Println("5. Cari Barang")
		fmt.Println("6. Tampilkan Barang")
		fmt.Println("7. Keluar")
		fmt.Print("Pilih opsi: ")

		var pilihan int
		fmt.Scan(&pilihan)

		switch pilihan {
		case 1:
			tambahBarang()
		case 2:
			ubahBarang()
		case 3:
			hapusBarang()
		case 4:
			fmt.Println("1. Transaksi Barang Masuk")
			fmt.Println("2. Transaksi Barang Keluar")
			var tran int
			fmt.Print("masukan opsi : ")
			fmt.Scan(&tran)
			switch tran{
			case 1 : 
				transaksiBarangMasuk()			
			case 2 :
				transaksiBarangKeluar()
			}		
		case 5:
			cariBarang()
		case 6:
			tampilkanBarang()
		case 7:
			simpanData()
			simpanDataTransaksi()
			simpanDataTransaksiKeluar()
			fmt.Println("Data disimpan. Keluar...")
			return
		default:
			fmt.Println("Pilihan tidak valid. Silakan coba lagi.")
		}
	}
}

func loadData() {
	file, err := os.Open(Dataset.DatabaseFile)
	if err != nil {
		fmt.Println("File database tidak ditemukan, membuat file baru...")
		file, _ = os.Create(Dataset.DatabaseFile)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		fields := strings.Split(line, ",")
		if len(fields) != 5 {
			continue
		}
		id, _ := strconv.Atoi(fields[0])
		stok, _ := strconv.Atoi(fields[2])
		harga, _ := strconv.ParseFloat(fields[3], 64)
		Dataset.BarangList = append(Dataset.BarangList, Dataset.Barang{
			ID:       id,
			Nama:     fields[1],
			Stok:     stok,
			Harga:    harga,
			Kategori: fields[4],
		})
	}
}
func loadDataTransaksi() {
	file, err := os.Open(Dataset.DataTransakasi)
	if err != nil {
		fmt.Println("File database tidak ditemukan, membuat file baru...")
		file, _ = os.Create(Dataset.DataTransakasi)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		fields := strings.Split(line, ",")
		if len(fields) != 5 {
			continue
		}
		id_transaksi , _ := strconv.Atoi(fields[0])
		id_Barang, _ := strconv.Atoi(fields[0])
		jumlah_barang, _ := strconv.Atoi(fields[2])
		Dataset.TransaksiList = append(Dataset.TransaksiList, Dataset.Transaksi{
			ID_Transaksi:       id_transaksi,
			ID_Barang :      	id_Barang,
			Nama_barang:     	fields[1],
			Jumlah:     		jumlah_barang,
		})
	}
}
func loadDataTransaksiKeluar() {
	file, err := os.Open(Dataset.DataTransakasiKeluar)
	if err != nil {
		fmt.Println("File database tidak ditemukan, membuat file baru...")
		file, _ = os.Create(Dataset.DataTransakasiKeluar)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		fields := strings.Split(line, ",")
		if len(fields) != 5 {
			continue
		}
		id_transaksi , _ := strconv.Atoi(fields[0])
		id_Barang, _ := strconv.Atoi(fields[0])
		jumlah_barang, _ := strconv.Atoi(fields[2])
		Dataset.TransakasiKeluar = append(Dataset.TransakasiKeluar, Dataset.Transaksi{
			ID_Transaksi:       id_transaksi,
			ID_Barang :      	id_Barang,
			Nama_barang:     	fields[1],
			Jumlah:     		jumlah_barang,
		})
	}
}

func simpanData() {
	file, _ := os.Create(Dataset.DatabaseFile)
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	for _, barang := range Dataset.BarangList {
		writer.Write([]string{
			strconv.Itoa(barang.ID),
			"Nama Barang    : "+ barang.Nama + "\n",
			"Jumlah Stok    : " + strconv.Itoa(barang.Stok)+"\n",
			"Harga barang   :" + fmt.Sprintf("%.2f", barang.Harga)+"\n",
			"Kategori Barang:" + barang.Kategori+ "\n",
		})
	}
}
func simpanDataTransaksi() {
	file, _ := os.Create(Dataset.DataTransakasi)
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	for _, Transaksi := range Dataset.TransaksiList {
		writer.Write([]string{
			"ID Transaksi  : " + strconv.Itoa(Transaksi.ID_Transaksi) ,
			"ID Barang     : " + strconv.Itoa(Transaksi.ID_Barang) ,
			"Nama Barang   : " + Transaksi.Nama_barang ,
			"Jumlah Barang : " + strconv.Itoa(Transaksi.Jumlah) ,
		})
	}
}
func simpanDataTransaksiKeluar() {
	file, _ := os.Create(Dataset.DataTransakasiKeluar)
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	for _, Transaksi := range Dataset.TransakasiKeluar {
		writer.Write([]string{
			"ID Transaksi  : " + strconv.Itoa(Transaksi.ID_Transaksi) ,
			"ID Barang     : " + strconv.Itoa(Transaksi.ID_Barang) ,
			"Nama Barang   : " + Transaksi.Nama_barang ,
			"Jumlah Barang : " + strconv.Itoa(Transaksi.Jumlah) ,
		})
	}
}

func transaksiBarangMasuk() {
	var id_transaksi,id, jumlah int
	var nama string
	fmt.Print("Masukkan ID Transakasi: ")
	fmt.Scan(&id_transaksi)
	fmt.Print("Masukkan ID Barang    : ")
	fmt.Scan(&id)
	fmt.Print("Masukkan Nama Barang  : ")
	fmt.Scan(&nama)
	fmt.Print("Jumlah Barang Masuk   : ")
	fmt.Scan(&jumlah)

	Dataset.TransaksiList = append(Dataset.TransaksiList, Dataset.Transaksi{
		ID_Transaksi:       id_transaksi,
		ID_Barang:     	  	id,
		Nama_barang:     	nama,
		Jumlah :     		jumlah,
	})

	for i, barang := range Dataset.BarangList {
		if barang.ID == id {
			Dataset.BarangList[i].Stok += jumlah
			fmt.Println("Transaksi berhasil.")
			return
		}
	}
	

	fmt.Println("Barang dengan ID tersebut tidak ditemukan.")
}

func transaksiBarangKeluar() {
	var id_transaksi,id, jumlah int
	var nama string
	fmt.Print("Masukkan ID Transakasi: ")
	fmt.Scan(&id_transaksi)
	fmt.Print("Masukkan ID Barang    : ")
	fmt.Scan(&id)
	fmt.Print("Masukkan Nama Barang  : ")
	fmt.Scan(&nama)
	fmt.Print("Jumlah Barang Masuk   : ")
	fmt.Scan(&jumlah)

	Dataset.TransakasiKeluar = append(Dataset.TransakasiKeluar, Dataset.Transaksi{
		ID_Transaksi:       id_transaksi,
		ID_Barang:     	  	id,
		Nama_barang:     	nama,
		Jumlah :     		jumlah,
	})

	for i, barang := range Dataset.BarangList {
		if barang.ID == id {
			if barang.Stok >= jumlah {
				Dataset.BarangList[i].Stok -= jumlah
				fmt.Println("Transaksi berhasil.")
			} else {
				fmt.Println("Stok tidak mencukupi.")
			}
			return
		}
	}

	fmt.Println("Barang dengan ID tersebut tidak ditemukan.")
}

func cariBarang() {
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
	fmt.Println("3. Tampilkan Transaksi Masuk")
	fmt.Println("4. Tampilkan Transaksi Keluar")
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
	case 3 :
		for _, transakasi := range Dataset.TransaksiList {
			fmt.Printf("ID Transaksi : %d, ID barang : %d, Nama Barang : %s, Jumlah : %d \n",
				transakasi.ID_Transaksi, transakasi.ID_Barang, transakasi.Nama_barang, transakasi.Jumlah)
		}
	case 4 :
		for _, transakasi := range Dataset.TransakasiKeluar {
			fmt.Printf("ID Transaksi : %d, ID barang : %d, Nama Barang : %s, Jumlah : %d \n",
				transakasi.ID_Transaksi, transakasi.ID_Barang, transakasi.Nama_barang, transakasi.Jumlah)}
	default:
		fmt.Println("Pilihan tidak valid.")
	}
}
