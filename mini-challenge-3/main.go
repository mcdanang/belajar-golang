package main

import (
	"fmt"
	"os"
)

type Person struct {
	ID        string
	nama      string
	alamat    string
	pekerjaan string
	alasan    string
}

func (p Person) getData() string {
	return fmt.Sprintf("ID : %s\nnama : %s\nalamat : %s\npekerjaan : %s\nalasan : %s\n", p.ID, p.nama, p.alamat, p.pekerjaan, p.alasan)
}

func getInitialData() []Person {
	return []Person{
		{"0", "Fitri", "Jl. Lorem", "Backend", "Alasan Fitri"},
		{"1", "Danang", "Jl. Ipsum", "Frontend", "Alasan Danang"},
		{"2", "Anggraini", "Jl. Dolor", "Fullstack", "Alasan Anggraini"},
	}
}

func main() {
	if len(os.Args) > 1 {

		// Initialize Data
		people := getInitialData()
		arg := os.Args[1]
		var result string

		for _, val := range people {
			if val.ID == arg || val.nama == arg {
				result = val.getData()
				break
			}
		}

		if result == "" {
			result = fmt.Sprintf("Tidak ditemukan data dengan ID / nama yang sesuai dengan argumen\n")
		}

		fmt.Printf("%s", result)

	} else {
		fmt.Println("Tolong masukan nama atau nomor absen")
		fmt.Println("Contoh: 'go run main.go Fitri' atau 'go run main.go 2'")
	}
}
