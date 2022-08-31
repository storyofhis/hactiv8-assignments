package main

import (
	"fmt"
	"os"
)

type Biodata struct {
	ID        string
	Nama      string
	Alamat    string
	Pekerjaan string
	Alasan    string
}

var allBiodata = []Biodata{
	{ID: "B001", Nama: "Elbert Siht", Alamat: "New York", Pekerjaan: "Photographer", Alasan: "Karena Mau Pindah Job"},
	{ID: "B002", Nama: "Michael Jackson", Alamat: "Alaska", Pekerjaan: "Chef", Alasan: "Karena Mau Pindah Job"},
	{ID: "B003", Nama: "Nial Horan", Alamat: "NewCastle", Pekerjaan: "Football Player", Alasan: "Karena Mau Pindah Job"},
	{ID: "B004", Nama: "Mark Zuckerburg", Alamat: "Jakarta", Pekerjaan: "Software Engineer", Alasan: "Mendalami Karir"},
	{ID: "B005", Nama: "James Richs", Alamat: "London", Pekerjaan: "Data Engineer", Alasan: "Mau Nyoba di karir baru"},
	{ID: "B006", Nama: "Ritch Carlton", Alamat: "Wasington", Pekerjaan: "Data Scientist", Alasan: "Mau Nyoba software engineer"},
	{ID: "B007", Nama: "Taylor Switf", Alamat: "Banyuwangi", Pekerjaan: "IOS Developerr", Alasan: "Backend Developer"},
	{ID: "B008", Nama: "Nicky Minaj", Alamat: "Malang", Pekerjaan: "Mahasiswa", Alasan: "Mau Ikut Studi Independen"},
	{ID: "B009", Nama: "Justin Mieber", Alamat: "MojoKerto", Pekerjaan: "Frontend Engineer", Alasan: "Mau Fullstack"},
	{ID: "A001", Nama: "Billie Elish", Alamat: "Ngawi", Pekerjaan: "Bussiness Inteligent", Alasan: "Mau ngikutin Perkembangan zaman"},
	{ID: "A002", Nama: "Lukas Graham", Alamat: "Blitar", Pekerjaan: "Feature Engineer", Alasan: "Penasaran ajasih"},
	{ID: "A003", Nama: "Wayne Rooney", Alamat: "Madiun", Pekerjaan: "Data Analyst", Alasan: "Diajak Temen"},
	{ID: "A004", Nama: "Cristiano Ronaldo", Alamat: "Kediri", Pekerjaan: "DevOps Developer", Alasan: "Gak mo kalah sama messi"},
	{ID: "A005", Nama: "Lionel Messi", Alamat: "Bali", Pekerjaan: "Cloud Engineer", Alasan: "Gak mo kalah sama ronaldo"},
	{ID: "A006", Nama: "Neymar Jr", Alamat: "Padang", Pekerjaan: "Flutter Engineer", Alasan: "Ikut-ikutan Ronaldo sama messi"},
}

func disPlay() {
	for _, each := range allBiodata {
		fmt.Printf("ID: %v, Nama : %v, Alamat: %v, Pekerjaan: %v, Alasan: %v\n", each.ID, each.Nama, each.Alamat, each.Pekerjaan, each.Alasan)
	}
}

func main() {
	var input string
	if os.Args[1] == "all" {
		disPlay()
	} else if os.Args[1] == "id" {
		fmt.Println("cari id berapa : ")
		fmt.Scanf("%v\n", &input)
		for _, each := range allBiodata {
			if input == each.ID {
				fmt.Printf("ID: %v, Nama : %v, Alamat: %v, Pekerjaan: %v, Alasan: %v\n", each.ID, each.Nama, each.Alamat, each.Pekerjaan, each.Alasan)
			}
		}
	}
}
