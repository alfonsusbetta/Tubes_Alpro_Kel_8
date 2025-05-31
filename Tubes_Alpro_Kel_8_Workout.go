package main

import "fmt"

type Workout struct {
	Jenis string
	Waktu int 
}

const NMAX int = 10 
	      
func main() {
	var daftarWorkout [NMAX]Workout 
	var jumlahWorkout int 
	TampilkanMenu(&daftarWorkout, &jumlahWorkout)
}
	
func TampilkanMenu(daftarWorkout *[NMAX]Workout, jumlahWorkout *int) {
/*IS: Menu belum ditampilkan, program menunggu input pengguna.*/
/*FS: Menu tampil di layar, sistem siap memproses pilihan pengguna.*/
	fmt.Println()
	for {
		fmt.Println("======== Workout Menu ========")
		fmt.Println("1. Tambah Workout")
		fmt.Println("2. Lihat Riwayat Workout")
		fmt.Println("3. Edit Riwayat Workout")
		fmt.Println("4. Urutkan Riwayat Workout")
		fmt.Println("5. Cari Latihan")
		fmt.Println("6. Rekomendasi Workout")
		fmt.Println("7. Tampilkan Laporan")
		fmt.Println("8. Keluar")
		fmt.Println("==============================")
		fmt.Print("Pilih opsi: ")

		var pilihan string
		fmt.Scanln(&pilihan)

		switch pilihan {
		case "1":
			tambahWorkout(daftarWorkout, jumlahWorkout)
		case "2":
			lihatWorkout(*daftarWorkout, *jumlahWorkout)
		case "3":
			editWorkout(daftarWorkout, jumlahWorkout)
		case "4":
			urutkanWorkout(daftarWorkout, *jumlahWorkout)
		case "5":
			cariWorkout(*daftarWorkout, *jumlahWorkout)
		case "6":
			rekomendasiWorkout(*daftarWorkout, *jumlahWorkout)
		case "7":
			tampilkanLaporan(*daftarWorkout, *jumlahWorkout)
		case "8":
			fmt.Println("Thank you🙏🏼and 𝐠𝐨𝐨𝐝𝐛𝐲𝐞👋🏻")
			return 
		}
	}
}

func tambahWorkout(daftarWorkout *[NMAX]Workout, jumlahWorkout *int){
/*IS: jumlahWorkout menunjukkan jumlah workout yang telah tercatat, dan pengguna belum memasukkan data workout baru.*/
/*FS: Sebuah workout baru dengan jenis dan waktu tertentu ditambahkan ke array daftarWorkout, dan jumlahWorkout bertambah satu.*/
	if *jumlahWorkout >= len(daftarWorkout) {
		fmt.Println("Daftar workout sudah penuh. Tidak bisa menambah workout lagi.\n")
		return
	}

	fmt.Println("======== Tambah Workout ========")
	fmt.Println("1. Push-up")
	fmt.Println("2. Sit-up")
	fmt.Println("3. Plank")
	fmt.Println("4. Squat")
	fmt.Println("5. Lari di tempat")
	fmt.Print("Pilih jenis workout (1-5): ")

	var pilihan int
	fmt.Scanln(&pilihan)

	var jenis string
	switch pilihan {
	case 1:
		jenis = "Push-up"
	case 2:
		jenis = "Sit-up"
	case 3:
		jenis = "Plank"
	case 4:
		jenis = "Squat"
	case 5:
		jenis = "Lari_di_tempat"
	default:
		fmt.Println("Pilihan tidak valid.\n")
		return
	}
    var waktu int
	fmt.Print("Masukkan waktu selama workout (menit): ")
	fmt.Scanln(&waktu)

	daftarWorkout[*jumlahWorkout] = Workout{Jenis: jenis, Waktu: waktu}
	*jumlahWorkout++ 
	fmt.Println("Workout berhasil ditambahkan!\n")
}

func lihatWorkout(daftarWorkout [NMAX]Workout, jumlahWorkout int) {
/*IS: Variabel daftarWorkout berisi sejumlah data workout (bisa kosong atau terisi), dan belum ada informasi riwayat workout yang dicetak ke layar.*/
/*FS: Jika jumlahWorkout > 0, seluruh data workout ditampilkan ke layar secara terformat. Jika tidak, muncul pesan bahwa belum ada workout yang tercatat.*/
	fmt.Println()
	if jumlahWorkout == 0 {
		fmt.Println("Belum ada workout yang tercatat.\n")
		return
	}
	fmt.Println("======== Riwayat Workout ========")
	for i := 0; i < jumlahWorkout; i++ {
		fmt.Printf("%d.  %-15s |  %d menit\n", i+1, daftarWorkout[i].Jenis, daftarWorkout[i].Waktu)
	}
	fmt.Println()
}
	
func editWorkout(daftarWorkout *[NMAX]Workout, jumlahWorkout *int) {
/*IS: Data Workout sudah tersedia, pengguna belum memilih workout untuk diedit atau dihapus.*/
/*FS: Data workout berhasil diubah atau dihapus, array diperbarui.*/
	fmt.Println("======== Edit Riwayat Workout ========")
	if *jumlahWorkout == 0 {
		fmt.Println("Belum ada workout yang tercatat.\n")
		return
	}

	lihatWorkout(*daftarWorkout, *jumlahWorkout)
	fmt.Print("Pilih nomor workout yang ingin diedit: ")
	var nomor int
	fmt.Scanln(&nomor)

	if nomor < 1 || nomor > *jumlahWorkout {
		fmt.Println("Nomor tidak valid.\n")
		return
	}

	var index int = nomor - 1

	fmt.Println("1. Ubah Jenis Workout")
	fmt.Println("2. Ubah Waktu Workout")
	fmt.Println("3. Hapus Workout")
	fmt.Print("Pilih opsi yang ingin diubah (1-3): ")
	var opsi int
	fmt.Scanln(&opsi)

	var pilihan int
	var waktu int

	switch opsi {
	case 1:
		fmt.Println("Pilih jenis workout baru:")
		fmt.Println("1. Push-up")
		fmt.Println("2. Sit-up")
		fmt.Println("3. Plank")
		fmt.Println("4. Squat")
		fmt.Println("5. Lari di tempat")
		fmt.Print("Pilih jenis workout (1-5): ")
		fmt.Scanln(&pilihan)

		switch pilihan {
		case 1:
			daftarWorkout[index].Jenis = "Push-up"
		case 2:
			daftarWorkout[index].Jenis = "Sit-up"
		case 3:
			daftarWorkout[index].Jenis = "Plank"
		case 4:
			daftarWorkout[index].Jenis = "Squat"
		case 5:
			daftarWorkout[index].Jenis = "Lari_di_tempat"
		default:
			fmt.Println("Pilihan tidak valid.\n")
			return
		}
		fmt.Println("Jenis workout berhasil diubah!\n")

	case 2:
		fmt.Print("Masukkan waktu baru selama workout (menit): ")
		fmt.Scanln(&waktu)
		daftarWorkout[index].Waktu = waktu
		fmt.Println("Waktu workout berhasil diubah!\n")

	case 3:
		for i := index; i < *jumlahWorkout-1; i++ {
			(*daftarWorkout)[i] = (*daftarWorkout)[i+1]
		}
		*jumlahWorkout--
		fmt.Println("Workout berhasil dihapus!\n")

	default:
		fmt.Println("Pilihan tidak valid.\n")
	}
}



func urutkanWorkout(daftarWorkout *[NMAX]Workout, jumlahWorkout int) {
/*IS: Data belum terurut, pengguna belum memilih metode pengurutan.*/
/*FS: Data workout diurutkan sesuai pilihan pengguna dan ditampilkan.*/
	var pilihan int
	var n int = jumlahWorkout
	fmt.Print("Pilih Opsi 1(Abjad) atau 2(Waktu):")
	fmt.Scanln(&pilihan)
		switch pilihan {
			case 1: 
				metodeAbjad(daftarWorkout, jumlahWorkout)
				fmt.Println("======== Riwayat Workout ========")
				for i:= 0; i < n; i++{
					fmt.Printf("%d.  %-15s |  %d menit\n", i+1, daftarWorkout[i].Jenis, daftarWorkout[i].Waktu)
				}
				fmt.Println()
			case 2:
				metodeWaktu(daftarWorkout, jumlahWorkout)
				fmt.Println("======== Riwayat Workout ========")
				for i:= 0; i < n; i++{
					fmt.Printf("%d.  %-15s |  %d menit\n", i+1, daftarWorkout[i].Jenis, daftarWorkout[i].Waktu)
				}
				fmt.Println()
		}
}

func metodeAbjad(daftarWorkout *[NMAX]Workout, jumlahWorkout int) {
/*IS: Data Workout belum diurutkan berdasarkan nama.*/
/*FS: Data Workout diurutkan menurun berdasarkan abjad.*/
	var i, idx, pass int
	var n int = jumlahWorkout
	var temp Workout
	pass = 1 
	for pass < n {
		idx = pass - 1
		i = pass 
		for i < n {
			if daftarWorkout[i].Jenis > daftarWorkout[idx].Jenis {
				idx = i
			}
			i++
		}
		temp = daftarWorkout[pass-1]
		daftarWorkout[pass-1] = daftarWorkout[idx]
		daftarWorkout[idx] = temp
		pass++
	}
}

func metodeWaktu(daftarWorkout *[NMAX]Workout, jumlahWorkout int) {
/*IS: Data Workout belum diurutkan berdasarkan waktu.*/
/*FS: Data Workout diurutkan menaik berdasarkan durasi.*/
	var i, pass int
	var temp Workout
	pass = 1
	for pass <= jumlahWorkout-1 {
		i = pass
		temp = (*daftarWorkout)[pass]
		for i > 0 && temp.Waktu < (*daftarWorkout)[i-1].Waktu {
			(*daftarWorkout)[i] = (*daftarWorkout)[i-1]
			i--
		}
		(*daftarWorkout)[i] = temp
		pass++
	}
}

func cariWorkout(daftarWorkout [NMAX]Workout, jumlahWorkout int) {
/*IS: Data workout tersedia, pengguna belum melakukan pencarian.*/
/*FS: Jika data ditemukan, workout dengan jenis atau waktu tertentu ditampilkan. Jika tidak, pesan bahwa data tidak ditemukan muncul.*/
	fmt.Println("Cari Workout dengan format apa?")
	fmt.Println("1. Sequential search")
	fmt.Println("2. Binary search")
	fmt.Println("3. Nilai Ekstrim")
	var metode int 
	fmt.Print("Pilih opsi: ")
	fmt.Scanln(&metode)
		switch metode {
			case 1: 	
			var jenis string
			fmt.Print("Pilih Workout (abjad): ")
			fmt.Scanln(&jenis)

			ditemukan := false
			for i := 0; i < jumlahWorkout; i++ {
				if daftarWorkout[i].Jenis == jenis {
					if !ditemukan {
						fmt.Println("\nWorkout ditemukan:")
					}
					fmt.Println("======== Riwayat Workout ========")
					fmt.Printf("%d.  %-15s |  %d menit\n", i+1, daftarWorkout[i].Jenis, daftarWorkout[i].Waktu)
					ditemukan = true
				}
			}
			if !ditemukan {
				fmt.Printf("Workout dengan jenis %-15s tidak ditemukan.\n", jenis)
				fmt.Println()
			}
			fmt.Println()
			case 2:
				metodeWaktu(&daftarWorkout, jumlahWorkout)
				var waktu int
				fmt.Print("Masukkan waktu workout yang ingin dicari (menit): ")
				fmt.Scanln(&waktu)
				kiri := 0
				kanan := jumlahWorkout - 1
				ditemukan := false

				for kiri <= kanan {
					tengah := (kiri + kanan) / 2
					if daftarWorkout[tengah].Waktu == waktu {
						fmt.Println("\nWorkout ditemukan:")
						fmt.Println("======== Riwayat Workout ========")
						fmt.Printf("%d.  %-15s |  %d menit\n", tengah+1, daftarWorkout[tengah].Jenis, daftarWorkout[tengah].Waktu)
						ditemukan = true
						fmt.Println()
						return
					} else if daftarWorkout[tengah].Waktu < waktu {
						kiri = tengah + 1
					} else {
						kanan = tengah - 1
					}
				}
				if !ditemukan {
					fmt.Printf("Workout dengan waktu %d menit tidak ditemukan.\n", waktu)
				}
				fmt.Println()
			case 3:
				var opsi int 
				fmt.Println("1. Waktu terbesar")
				fmt.Println("2. Waktu terkecil")
				fmt.Print("Pilih opsi: ")
				fmt.Scanln(&opsi)
				switch opsi {
				case 1 : 
					var max int 
					var k int 
					var idx int 
					max = daftarWorkout[0].Waktu
					k = 1 
					for k < jumlahWorkout {
						if max < daftarWorkout[k].Waktu {
						   max = daftarWorkout[k].Waktu 
						   idx = k
						}
						k++
					}
					fmt.Println("\nWorkout ditemukan:")
					fmt.Println("======== Riwayat Workout ========")
					fmt.Printf("%d.  %-15s |  %d menit\n", idx+1, daftarWorkout[idx].Jenis, max)
					fmt.Println()
				case 2: 
					var min int 
					var k int 
					var idx int 
					min = daftarWorkout[0].Waktu
					k = 1 
					for k < jumlahWorkout {
						if min > daftarWorkout[k].Waktu {
						   min = daftarWorkout[k].Waktu 
						   idx = k 
						}
						k++ 
					}
					fmt.Println("\nWorkout ditemukan:")
					fmt.Println("======== Riwayat Workout ========")
					fmt.Printf("%d.  %-15s |  %d menit\n", idx+1, daftarWorkout[idx].Jenis, min)
					fmt.Println()
			default:
				fmt.Println("Metode tidak valid.")
		}
	}
}


func tampilkanLaporan(daftarWorkout [NMAX]Workout, jumlahWorkout int) {
/*IS: Data workout sudah ada tapi belum dihitung kalorinya.*/
/*FS: Estimasi kalori per workout dan total kalori ditampilkan.*/
	lihatWorkout(daftarWorkout, jumlahWorkout)
	HitungKalori(daftarWorkout)
	fmt.Println()
}

func HitungKalori(daftarWorkout [NMAX]Workout) {
/*IS: Data workout tersedia, kalori belum dihitung.*/
/*FS: Kalori tiap workout dihitung dan dijumlahkan.*/
	kaloriPushUp := 0
    kaloriSitUp := 0
    kaloriPlank := 0
    kaloriSquat := 0
    kaloriLariDiTempat := 0
	totalKalori := 0 
	
	adaPushUp := false
	adaSitUp := false
	adaPlank := false
	adaSquat := false
	adaLari := false

    for i := 0; i < len(daftarWorkout); i++ {
        switch daftarWorkout[i].Jenis {
        case "Push-up":
            kaloriPushUp += 7 * daftarWorkout[i].Waktu
			totalKalori += kaloriPushUp
			adaPushUp = true
        case "Sit-up":
            kaloriSitUp += 5 * daftarWorkout[i].Waktu
			totalKalori += kaloriSitUp
			adaSitUp = true
        case "Plank":
            kaloriPlank += 4 * daftarWorkout[i].Waktu
			totalKalori += kaloriPlank
			adaPlank = true
        case "Squat":
            kaloriSquat += 8 * daftarWorkout[i].Waktu
			totalKalori += kaloriSquat
			adaSquat = true
        case "Lari_di_tempat":
            kaloriLariDiTempat += 10 * daftarWorkout[i].Waktu
			totalKalori += kaloriLariDiTempat
			adaLari = true
        }
    }
	
	
    fmt.Println("Total Kalori yang terbakar:")
		if adaPushUp {
			fmt.Printf("Push-up: %d kalori\n", kaloriPushUp)
		}
		if adaSitUp {
			fmt.Printf("Sit-up: %d kalori\n", kaloriSitUp)
		}
		if adaPlank {
			fmt.Printf("Plank: %d kalori\n", kaloriPlank)
		}
		if adaSquat {
			fmt.Printf("Squat: %d kalori\n", kaloriSquat)
		}
		if adaLari {
			fmt.Printf("Lari di tempat: %d kalori\n", kaloriLariDiTempat)
		}

	fmt.Printf("Total kalori seluruhnya: %d kalori\n", totalKalori)
}

func rekomendasiWorkout(daftarWorkout [NMAX]Workout, jumlahWorkout int) {
/*IS: Tersedia data workout, tetapi belum ada analisis rekomendasi.*/
/*FS: Sistem memberi rekomendasi berdasarkan frekuensi jenis workout.*/
	fmt.Println()
	if jumlahWorkout == 0 {
		fmt.Println("Belum ada workout yang tercatat.\n")
		return
	} 
	jmlPushUp := 0
	jmlSitUp := 0
	jmlPlank := 0
	jmlSquad := 0
	jmlLari_di_tempat := 0
	
	for i:= 0; i < len(daftarWorkout); i++ {
		if daftarWorkout[i].Jenis == "Push-up" {
			jmlPushUp++
		}else if daftarWorkout[i].Jenis == "Sit-up" {
			jmlSitUp++ 
		}else if daftarWorkout[i].Jenis == "Plank" {
			jmlPlank++
		}else if daftarWorkout[i].Jenis == "Squat" {
			jmlSquad++ 
		}else if daftarWorkout[i].Jenis == "Lari_di_tempat" {
			jmlLari_di_tempat++ 
		}
	}
	fmt.Println("Rekomendasi Workout dari kami:")
	if jmlPushUp >= 3 && jmlSitUp >= 3{
		fmt.Println("Cobain Wall Sit deh, karena kamu sudah sering Push-up dan Sit-up")
	}
	if jmlPlank >= 3 && jmlSquad >= 3 {
		fmt.Println("Cobain Stationary Lunge deh, karena kamu sudah sering Plank dan Squat ")
	}
	if jmlLari_di_tempat >= 3 {
		fmt.Println("Cobain Cardio intensif deh, karena kamu sudah sering Lari di tempat di tempat")
	}
if (jmlPushUp < 2 && jmlPlank < 2 ) || (jmlSitUp < 2 && jmlSquad < 2) && jmlLari_di_tempat < 2 {
		fmt.Println("Coba fokus di satu atau dua Workout dulu")
	}
	fmt.Println()
}
