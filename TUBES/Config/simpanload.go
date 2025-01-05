package Config

import (
	"TUBES/Dataset"
	"bufio"
	"encoding/csv"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func LoadData() {
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
func LoadDataTransaksi() {
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
		if len(fields) != 4 {
			continue
		}
		id_Barang, _ := strconv.Atoi(fields[1])
		jumlah_barang, _ := strconv.Atoi(fields[3])
		Dataset.TransaksiList = append(Dataset.TransaksiList, Dataset.Transaksi{
			ID_Transaksi: fields[0],
			ID_Barang:    id_Barang,
			Nama_barang:  fields[2],
			Jumlah:       jumlah_barang,
		})
	}
}
func LoadDataTransaksiKeluar() {
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
		if len(fields) != 4 {
			continue
		}
		
		id_Barang, _ := strconv.Atoi(fields[1])
		jumlah_barang, _ := strconv.Atoi(fields[3])
		Dataset.TransakasiKeluar = append(Dataset.TransakasiKeluar, Dataset.Transaksi{
			ID_Transaksi: fields[0],
			ID_Barang:    id_Barang,
			Nama_barang:  fields[2],
			Jumlah:       jumlah_barang,
		})
	}
}

func Simpan() {
	file, err := os.Create(Dataset.DatabaseFile)
	if err != nil {
		fmt.Println("Error creating file:", err)
		return
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	for _, barang := range Dataset.BarangList {
		writer.Write([]string{
			strconv.Itoa(barang.ID),
			barang.Nama ,
			strconv.Itoa(barang.Stok) ,
			fmt.Sprintf("%.2f", barang.Harga) ,
			barang.Kategori ,
		})
	}
}
func SimpanTransaksi() {
	file, err := os.Create(Dataset.DataTransakasi)
	if err != nil {
		fmt.Println("Error creating file:", err)
		return
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	for _, Transaksi := range Dataset.TransaksiList {
		writer.Write([]string{
			Transaksi.ID_Transaksi,
			strconv.Itoa(Transaksi.ID_Barang),
			Transaksi.Nama_barang,
			strconv.Itoa(Transaksi.Jumlah),
		})
	}
}
func SimpanTransaksiKeluar() {
	file, err := os.Create(Dataset.DataTransakasiKeluar)
	if err != nil {
		fmt.Println("Error creating file:", err)
		return
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	for _, Transaksi := range Dataset.TransakasiKeluar {
		writer.Write([]string{
			Transaksi.ID_Transaksi,
			strconv.Itoa(Transaksi.ID_Barang),
			Transaksi.Nama_barang,
			strconv.Itoa(Transaksi.Jumlah),
		})
	}
}