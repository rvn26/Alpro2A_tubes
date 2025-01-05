package Dataset


type Barang struct {
	ID       int
	Nama     string
	Stok     int
	Harga    float64
	Kategori string
}
type Transaksi struct{
	ID_Transaksi string
	ID_Barang int
	Nama_barang string
	Jumlah int
}

var BarangList []Barang
var TransaksiList []Transaksi
var TransakasiKeluar []Transaksi
var DatabaseFile = "database.txt"
var DataTransakasi = "Transaksi.txt"
var DataTransakasiKeluar = "Transaksi_Keluar.txt"