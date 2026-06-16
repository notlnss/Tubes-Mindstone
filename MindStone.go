package main

import "fmt"

const NMAX int = 10000

type tanggal struct {
	hari  int
	bulan int
	tahun int
}

type Pencapaian struct {
	namaTugas   string
	kepentingan int
	skorMood    int
	skorStres   int
	tanggal     tanggal
	selesai     bool
}

type daftarPencapaian [NMAX]Pencapaian

func main() {
	var n int
	var pilihan string
	var A daftarPencapaian
	var B daftarPencapaian
	n = 0
	fmt.Println("==================================================")
	fmt.Println("   M I N D S T O N E  -  Capaian & Mental Track  ")
	fmt.Println("==================================================")
	for {
		B = A
		fmt.Println("\n==================================================")
		fmt.Println("                     M E N U                    ")
		fmt.Println("==================================================")
		fmt.Println("1.  Tambah Data Pencapaian")
		fmt.Println("2.  Ubah Data Pencapaian")
		fmt.Println("3.  Hapus Data Pencapaian")
		fmt.Println("4.  Tampilkan Semua Data")
		fmt.Println("5.  Cari Berdasarkan Nama Tugas (Sequential Search)")
		fmt.Println("6.  Cari Berdasarkan Tanggal (Binary Search)")
		fmt.Println("7.  Urutkan Berdasarkan Kepentingan (Selection Sort)")
		fmt.Println("8.  Urutkan Berdasarkan Tanggal (Selection Sort)")
		fmt.Println("9.  Urutkan Berdasarkan Skor Mood (Insertion Sort)")
		fmt.Println("10. Statistik ")
		fmt.Println("0.  Keluar")
		fmt.Println("==================================================")
		fmt.Print("Pilih (1-10): ")
		fmt.Scan(&pilihan)
		if pilihan == "1" {
			tambahPencapaian(&A, &n)
		} else if pilihan == "2" {
			ubahPencapaian(&A, n)
		} else if pilihan == "3" {
			hapusPencapaian(&A, &n)
		} else if pilihan == "4" {
			cetakSemua(A, n)
		} else if pilihan == "5" {
			cariSequential(&A, n)
		} else if pilihan == "6" {
			cariBinaryTanggal(&B, n)
		} else if pilihan == "7" {
			selectionSortKepentingan(&A, n)
			fmt.Println("Data berhasil diurutkan berdasarkan kepentingan (tinggi ke rendah).")
			cetakSemua(A, n)
		} else if pilihan == "8" {
			selectionSortTanggal(&A, n)
			fmt.Println("Data berhasil diurutkan berdasarkan tanggal.")
			cetakSemua(A, n)
		} else if pilihan == "9" {
			insertionSortMood(&A, n)
			fmt.Println("Data berhasil diurutkan berdasarkan skor mood (rendah ke tinggi).")
			cetakSemua(A, n)
		} else if pilihan == "10" {
			statistik(A, n)
		} else if pilihan == "0" {
			fmt.Println("Terima kasih telah menggunakan MindStone. Sampai jumpa!")
			return
		} else {
			fmt.Println("Pilihan tidak valid. Silakan coba lagi.")
		}
	}
}

// FUNGSI INPUT TANGGAL

func inputTanggal() tanggal {
	var t tanggal
	fmt.Print("Tanggal selesai - tanggal  : ")
	fmt.Scan(&t.hari)
	fmt.Print("                  bulan : ")
	fmt.Scan(&t.bulan)
	fmt.Print("                  tahun : ")
	fmt.Scan(&t.tahun)
	return t
}

func tanggalKeInt(t tanggal) int {
	return t.tahun*10000 + t.bulan*100 + t.hari

}

// FUNGSI PRINT KETERANGAN

func keteranganKepentingan(k int) string {
	var ket string
	if k == 1 {
		ket = "Rendah"
	} else if k == 2 {
		ket = "Sedang"
	} else {
		ket = "Tinggi"
	}
	return ket
}

func keteranganMood(s int) string {
	var ket string
	if s == 1 {
		ket = "Sangat Buruk"
	} else if s == 2 {
		ket = "Buruk"
	} else if s == 3 {
		ket = "Cukup"
	} else if s == 4 {
		ket = "Baik"
	} else {
		ket = "Sangat Baik"
	}
	return ket
}

func keteranganStres(s int) string {
	var ket string
	if s == 1 {
		ket = "Sangat Rendah"
	} else if s == 2 {
		ket = "Rendah"
	} else if s == 3 {
		ket = "Cukup"
	} else if s == 4 {
		ket = "Tinggi"
	} else {
		ket = "Sangat Tinggi"
	}
	return ket
}

// FUNGSI INPUT TERVALIDASI

func inputKepentingan() int {
	var k int
	fmt.Println("Tingkat kepentingan:")
	fmt.Println("  1 = Rendah")
	fmt.Println("  2 = Sedang")
	fmt.Println("  3 = Tinggi")
	fmt.Print("Masukkan pilihan : ")
	fmt.Scan(&k)
	for k < 1 || k > 3 {
		fmt.Print("Input tidak valid! Masukkan angka 1-3: ")
		fmt.Scan(&k)
	}
	return k
}

func inputMood() int {
	var s int
	fmt.Println("Skor mood:")
	fmt.Println("  1 = Sangat Buruk")
	fmt.Println("  2 = Buruk")
	fmt.Println("  3 = Cukup")
	fmt.Println("  4 = Baik")
	fmt.Println("  5 = Sangat Baik")
	fmt.Print("Masukkan pilihan : ")
	fmt.Scan(&s)
	for s < 1 || s > 5 {
		fmt.Print("Input tidak valid! Masukkan angka 1-5: ")
		fmt.Scan(&s)
	}
	return s
}

func inputStres() int {
	var s int
	fmt.Println("Skor stres:")
	fmt.Println("  1 = Sangat Rendah")
	fmt.Println("  2 = Rendah")
	fmt.Println("  3 = Cukup")
	fmt.Println("  4 = Tinggi")
	fmt.Println("  5 = Sangat Tinggi")
	fmt.Print("Masukkan pilihan : ")
	fmt.Scan(&s)
	for s < 1 || s > 5 {
		fmt.Print("Input tidak valid! Masukkan angka 1-5: ")
		fmt.Scan(&s)
	}
	return s
}

func inputSelesai() bool {
	var v int
	fmt.Print("Status selesai? (1=Ya, 0=Tidak): ")
	fmt.Scan(&v)
	for v < 0 || v > 1 {
		fmt.Print("Input tidak valid! Masukkan 0 atau 1: ")
		fmt.Scan(&v)
	}
	return v == 1
}

func cariNama(A daftarPencapaian, n int, keyword string) int {
	var found int
	var k int
	found = -1
	k = 0
	for found == -1 && k < n {
		if A[k].namaTugas == keyword {
			found = k
		}
		k = k + 1
	}
	return found
}

// FUNGSI TAMBAH DATA

func tambahPencapaian(A *daftarPencapaian, n *int) {
	var nama string
	var x, i int

	if *n >= NMAX {
		fmt.Println("Kapasitas maksimum tercapai. Tidak dapat menambah data.")
		return
	}

	fmt.Println("\n--- Tambah Pencapaian Baru ---")
	fmt.Print("Banyak Data Ditambah : ")
	fmt.Scan(&x)
	for i = 0; i < x; i = i + 1 {
		fmt.Println("Data Ke", *n+1)
		fmt.Print("Nama tugas       : ")
		fmt.Scan(&nama)
		A[*n].namaTugas = nama
		A[*n].kepentingan = inputKepentingan()
		A[*n].skorMood = inputMood()
		A[*n].skorStres = inputStres()
		A[*n].selesai = inputSelesai()
		A[*n].tanggal = inputTanggal()
		*n = *n + 1
	}
	fmt.Println("Pencapaian berhasil ditambahkan.")
}

// FUNGSI UBAH DATA

func ubahPencapaian(A *daftarPencapaian, n int) {
	var nama string
	var idx int

	fmt.Print("\nMasukkan nama tugas yang ingin diubah: ")
	fmt.Scan(&nama)

	idx = cariNama(*A, n, nama)

	if idx == -1 {
		fmt.Println("Tugas tidak ditemukan.")
		return
	}

	fmt.Printf("Data ditemukan: %s. Masukkan data baru:\n", A[idx].namaTugas)
	fmt.Print("Nama tugas baru  : ")
	fmt.Scan(&A[idx].namaTugas)
	A[idx].kepentingan = inputKepentingan()
	A[idx].skorMood = inputMood()
	A[idx].skorStres = inputStres()
	A[idx].selesai = inputSelesai()
	A[idx].tanggal = inputTanggal()

	fmt.Println("Data berhasil diubah.")
}

// FUNGSI HAPUS DATA

func hapusPencapaian(A *daftarPencapaian, n *int) {
	var nama string
	var idx, i int

	fmt.Print("\nMasukkan nama tugas yang ingin dihapus: ")
	fmt.Scan(&nama)

	idx = cariNama(*A, *n, nama)

	if idx == -1 {
		fmt.Println("Tugas tidak ditemukan.")
		return
	}

	i = idx
	for i <= *n-2 {
		A[i] = A[i+1]
		i = i + 1
	}
	*n = *n - 1
	fmt.Println("Pencapaian berhasil dihapus.")
}

// FUNGSI CETAK SEMUA DATA

func cetakSemua(A daftarPencapaian, n int) {
	var i int
	var statusStr string

	if n == 0 {
		fmt.Println("Belum ada data pencapaian.")
		return
	}

	fmt.Println("\n===== DAFTAR PENCAPAIAN =====")
	i = 0
	for i < n {
		if A[i].selesai {
			statusStr = "Selesai"
		} else {
			statusStr = "Belum Selesai"
		}
		fmt.Printf("\n[%d] Tugas      : %s\n", i+1, A[i].namaTugas)
		fmt.Printf("    Kepentingan: %d (%s)\n", A[i].kepentingan, keteranganKepentingan(A[i].kepentingan))
		fmt.Printf("    Mood       : %d (%s)\n", A[i].skorMood, keteranganMood(A[i].skorMood))
		fmt.Printf("    Stres      : %d (%s)\n", A[i].skorStres, keteranganStres(A[i].skorStres))
		fmt.Printf("    Tanggal    : %02d-%02d-%d\n", A[i].tanggal.hari, A[i].tanggal.bulan, A[i].tanggal.tahun)
		fmt.Printf("    Status     : %s\n", statusStr)
		i = i + 1
	}
}

//SEQUENTIAL SEARCH

func cariSequential(A *daftarPencapaian, n int) {
	var keyword string
	var i int
	var ditemukan bool
	var statusStr string
	ditemukan = false
	fmt.Print("\nMasukkan nama tugas yang dicari: ")
	fmt.Scan(&keyword)
	fmt.Println("\n--- Hasil Pencarian (Sequential Search) ---")
	i = 0
	for i < n {
		if A[i].namaTugas == keyword {
			if A[i].selesai {
				statusStr = "Selesai"
			} else {
				statusStr = "Belum Selesai"
			}
			fmt.Printf("Data Ke      : %d\n", i+1)
			fmt.Printf("Tugas        : %s\n", A[i].namaTugas)
			fmt.Printf("Kepentingan  : %d (%s)\n", A[i].kepentingan, keteranganKepentingan(A[i].kepentingan))
			fmt.Printf("Mood         : %d (%s)\n", A[i].skorMood, keteranganMood(A[i].skorMood))
			fmt.Printf("Stres        : %d (%s)\n", A[i].skorStres, keteranganStres(A[i].skorStres))
			fmt.Printf("Tanggal      : %02d-%02d-%d\n", A[i].tanggal.hari, A[i].tanggal.bulan, A[i].tanggal.tahun)
			fmt.Printf("Status       : %s\n", statusStr)
			ditemukan = true
		}
		i = i + 1
	}
	if !ditemukan {
		fmt.Println("Tugas tidak ditemukan.")
	}
}

// BINARY SEARCH

func cariBinaryTanggal(B *daftarPencapaian, n int) {
	var left, right, mid, target, found int
	var statusStr string

	selectionSortTanggal(B, n)
	target = tanggalKeInt(inputTanggal())

	left = 0
	right = n - 1
	found = -1

	for left <= right && found == -1 {
		mid = (left + right) / 2
		if tanggalKeInt(B[mid].tanggal) == target {
			found = mid
		} else if target < tanggalKeInt(B[mid].tanggal) {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}

	if found == -1 {
		fmt.Println("Tidak ada pencapaian pada tanggal tersebut.")
		return
	}

	if B[found].selesai {
		statusStr = "Selesai"
	} else {
		statusStr = "Belum Selesai"
	}
	fmt.Printf("Data Ditemukan !!")
	fmt.Printf("\nTugas      : %s\n", B[found].namaTugas)
	fmt.Printf("Kepentingan: %d (%s)\n", B[found].kepentingan, keteranganKepentingan(B[found].kepentingan))
	fmt.Printf("Mood       : %d (%s)\n", B[found].skorMood, keteranganMood(B[found].skorMood))
	fmt.Printf("Stres      : %d (%s)\n", B[found].skorStres, keteranganStres(B[found].skorStres))
	fmt.Printf("Tanggal    : %02d-%02d-%d\n", B[found].tanggal.hari, B[found].tanggal.bulan, B[found].tanggal.tahun)
	fmt.Printf("Status     : %s\n", statusStr)
}

// SELECTION SORT KEPENTINGAN
func selectionSortKepentingan(A *daftarPencapaian, n int) {
	var pass, j, idxMax int
	var temp Pencapaian
	pass = 0
	//DESCEND
	for pass < n-1 {
		idxMax = pass
		j = pass + 1
		for j < n {
			if A[j].kepentingan > A[idxMax].kepentingan {
				idxMax = j
			}
			j = j + 1
		}
		temp = A[pass]
		A[pass] = A[idxMax]
		A[idxMax] = temp
		pass = pass + 1
	}
}

// INSERtION SORT MOOD
func insertionSortMood(A *daftarPencapaian, n int) {
	var i, j int
	var temp Pencapaian
	//DESCEND
	i = 1
	for i < n {
		temp = A[i]
		j = i - 1
		for j >= 0 && A[j].skorMood > temp.skorMood {
			A[j+1] = A[j]
			j = j - 1
		}
		A[j+1] = temp
		i = i + 1
	}
}

// SELECTION SORT TANGGAL

func selectionSortTanggal(A *daftarPencapaian, n int) {
	var pass, j, idxMin int
	var temp Pencapaian
	//ASCEND
	pass = 0
	for pass < n-1 {
		idxMin = pass
		j = pass + 1
		for j < n {
			if tanggalKeInt(A[j].tanggal) < tanggalKeInt(A[idxMin].tanggal) {
				idxMin = j
			}
			j = j + 1
		}
		temp = A[pass]
		A[pass] = A[idxMin]
		A[idxMin] = temp
		pass = pass + 1
	}
}

// STATISTIK

func statistik(A daftarPencapaian, n int) {
	var pilihan int

	fmt.Println("\n===== STATISTIK =====")
	fmt.Println("1. Statistik per Minggu")
	fmt.Println("2. Statistik per Bulan")
	fmt.Print("Pilih opsi: ")
	fmt.Scan(&pilihan)

	if pilihan == 1 {
		statistikPerMinggu(A, n)
	} else if pilihan == 2 {
		statistikPerBulan(A, n)
	} else {
		fmt.Println("Pilihan tidak valid.")
	}
}

func statistikPerMinggu(A daftarPencapaian, n int) {
	var mingguInput, bulanInput, tahunInput int
	var i, hari int
	var totalStres, jumlahData, jumlahSelesai int
	var rataStres float64
	var persentaseSelesai float64

	fmt.Print("\nMasukkan minggu ke- (1-5): ")
	fmt.Scan(&mingguInput)
	fmt.Print("Masukkan bulan (1-12)   : ")
	fmt.Scan(&bulanInput)
	fmt.Print("Masukkan tahun          : ")
	fmt.Scan(&tahunInput)

	totalStres = 0
	jumlahData = 0
	jumlahSelesai = 0

	i = 0
	for i < n {
		if A[i].tanggal.bulan == bulanInput && A[i].tanggal.tahun == tahunInput {
			hari = A[i].tanggal.hari
			if hari >= (mingguInput-1)*7+1 && hari <= mingguInput*7 {
				totalStres = totalStres + A[i].skorStres
				jumlahData = jumlahData + 1
				if A[i].selesai {
					jumlahSelesai = jumlahSelesai + 1
				}
			}
		}
		i = i + 1
	}

	fmt.Printf("\n===== STATISTIK MINGGU KE-%d (%02d/%d) =====\n", mingguInput, bulanInput, tahunInput)
	if jumlahData == 0 {
		fmt.Println("Tidak ada data pada minggu tersebut.")
	} else {
		rataStres = float64(totalStres) / float64(jumlahData)
		persentaseSelesai = float64(jumlahSelesai) / float64(jumlahData) * 100.0
		fmt.Printf("Total tugas tercatat       : %d\n", jumlahData)
		fmt.Printf("Tugas berhasil diselesaikan: %d\n", jumlahSelesai)
		fmt.Printf("Persentase target selesai  : %.2f%%\n", persentaseSelesai)
		fmt.Printf("Rata-rata tingkat stres    : %.2f / 5 (%s)\n", rataStres, keteranganStres(int(rataStres+0.5)))
	}
}

func statistikPerBulan(A daftarPencapaian, n int) {
	var bulanInput, tahunInput int
	var i int
	var totalStres, jumlahData, jumlahSelesai int
	var rataStres float64
	var persentaseSelesai float64

	fmt.Print("\nMasukkan bulan (1-12)   : ")
	fmt.Scan(&bulanInput)
	fmt.Print("Masukkan tahun          : ")
	fmt.Scan(&tahunInput)

	totalStres = 0
	jumlahData = 0
	jumlahSelesai = 0

	i = 0
	for i < n {
		if A[i].tanggal.bulan == bulanInput && A[i].tanggal.tahun == tahunInput {
			totalStres = totalStres + A[i].skorStres
			jumlahData = jumlahData + 1
			if A[i].selesai {
				jumlahSelesai = jumlahSelesai + 1
			}
		}
		i = i + 1
	}

	fmt.Printf("\n===== STATISTIK BULAN %02d/%d =====\n", bulanInput, tahunInput)
	if jumlahData == 0 {
		fmt.Println("Tidak ada data pada bulan tersebut.")
	} else {
		rataStres = float64(totalStres) / float64(jumlahData)
		persentaseSelesai = float64(jumlahSelesai) / float64(jumlahData) * 100.0
		fmt.Printf("Total tugas tercatat       : %d\n", jumlahData)
		fmt.Printf("Tugas berhasil diselesaikan: %d\n", jumlahSelesai)
		fmt.Printf("Persentase target selesai  : %.2f%%\n", persentaseSelesai)
		fmt.Printf("Rata-rata tingkat stres    : %.2f / 5 (%s)\n", rataStres, keteranganStres(int(rataStres+0.5)))
	}
}
