package main

import (
	"fmt"
	"os"
)

// Struct untuk menyimpan informasi mahasiswa
type Mahasiswa struct {
	Username  string   // Username untuk login
	Password  string   // Password untuk login
	Saldo     float64  // Saldo awal
	Transaksi []string // Riwayat transaksi
}

// Slice untuk menyimpan daftar pengguna yang terdaftar di ATM
var users = []Mahasiswa{
	{"Kiarra", "2406357072", 3500000, []string{}},
	{"user2", "23456", 3500000, []string{}},
}

// Fungsi untuk memverifikasi username dan password pengguna
func verifikasi(username, password string) (*Mahasiswa, bool) {
	for i, user := range users {
		// Jika username dan password cocok, kembalikan data pengguna
		if user.Username == username && user.Password == password {
			return &users[i], true
		}
	}
	// Jika tidak cocok, kembalikan nil dan false
	return nil, false
}

// Fungsi untuk menampilkan informasi akun pengguna
func lihatAkun(user *Mahasiswa) {
	fmt.Println("=== Informasi Akun ===")
	fmt.Println("Username:", user.Username)
	fmt.Printf("Saldo Saat Ini: %.2f\n", user.Saldo)
}

// Fungsi untuk menampilkan saldo pengguna
func lihatSaldo(user *Mahasiswa) {
	fmt.Printf("Saldo Anda Saat Ini: %.2f\n", user.Saldo)
}

// Fungsi untuk menambah saldo pada akun pengguna
func tambahSaldo(user *Mahasiswa) {
	var jumlah float64
	fmt.Println("Masukkan jumlah saldo yang ingin ditambahkan:")
	fmt.Scanln(&jumlah)
	if jumlah > 0 {
		user.Saldo += jumlah
		user.Transaksi = append(user.Transaksi, fmt.Sprintf("Menambahkan saldo: %.2f", jumlah))
		fmt.Println("Saldo berhasil ditambahkan.")
	} else {
		fmt.Println("Jumlah tidak valid.")
	}
}

// Fungsi untuk menarik saldo dari akun pengguna
func tarikSaldo(user *Mahasiswa) {
	var jumlah float64
	fmt.Println("Masukkan jumlah saldo yang ingin ditarik:")
	fmt.Scanln(&jumlah)
	if jumlah > 0 && jumlah <= user.Saldo {
		user.Saldo -= jumlah
		user.Transaksi = append(user.Transaksi, fmt.Sprintf("Menarik saldo: %.2f", jumlah))
		fmt.Println("Saldo berhasil ditarik.")
	} else if jumlah > user.Saldo {
		fmt.Println("Saldo tidak mencukupi.")
	} else {
		fmt.Println("Jumlah tidak valid.")
	}
}

// Fungsi untuk menampilkan riwayat transaksi pengguna
func historyTransaksi(user *Mahasiswa) {
	fmt.Println("=== Histori Transaksi ===")
	if len(user.Transaksi) == 0 {
		fmt.Println("Tidak ada transaksi.")
	} else {
		for _, transaksi := range user.Transaksi {
			fmt.Println(transaksi)
		}
	}
}

// Fungsi utama untuk menjalankan program
func main() {
	var username, password string

	fmt.Println("=== Selamat Datang di ATM ===")
	// Input username dan password dari pengguna
	fmt.Print("Masukkan Username: ")
	fmt.Scanln(&username)
	fmt.Print("Masukkan Password: ")
	fmt.Scanln(&password)

	// Verifikasi username dan password
	user, valid := verifikasi(username, password)
	if !valid {
		fmt.Println("Username atau password salah.")
		os.Exit(1) // Keluar dari program jika verifikasi gagal
	}

	// Loop utama untuk menampilkan menu ATM
	for {
		fmt.Println("\n=== Menu ATM ===")
		fmt.Println("1. Lihat Informasi Akun")
		fmt.Println("2. Lihat Saldo")
		fmt.Println("3. Tambahkan Saldo")
		fmt.Println("4. Tarik Saldo")
		fmt.Println("5. Histori Transaksi")
		fmt.Println("6. Keluar")

		var pilihan int
		fmt.Print("Pilih menu: ")
		fmt.Scanln(&pilihan)

		// Pilihan menu berdasarkan input pengguna
		switch pilihan {
		case 1:
			lihatAkun(user)
		case 2:
			lihatSaldo(user)
		case 3:
			tambahSaldo(user)
		case 4:
			tarikSaldo(user)
		case 5:
			historyTransaksi(user)
		case 6:
			fmt.Println("Terima kasih telah menggunakan ATM!! Tolong jangan boros")
			os.Exit(0) // Keluar dari program
		default:
			fmt.Println("Pilihan tidak valid.")
		}
	}
}