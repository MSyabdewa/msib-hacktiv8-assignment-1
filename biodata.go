package main

import (
	"encoding/json"
	"fmt"
	"log"
	"msib-hacktiv8-assignment-1/model"
	"os"
	"path/filepath"
	"strings"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Usage: go run biodata.go <student_code>")
		return
	}

	jsonPath := filepath.Join("data", "participants.json")
	data, err := os.ReadFile(jsonPath)
	if err != nil {
		log.Fatal(err)
	}

	var students model.Students
	if err := json.Unmarshal(data, &students); err != nil {
		log.Fatal(err)
	}

	studentCode := strings.ToLower(os.Args[1]) // Mengonversi input pengguna menjadi huruf kecil
	found := false

	for _, student := range students.Students {
		if strings.ToLower(student.Code) == studentCode { // Mengonversi kode siswa dari JSON menjadi huruf kecil untuk perbandingan
			found = true
			fmt.Printf("ID: %s\n", student.Id)
			fmt.Printf("Code: %s\n", student.Code)
			fmt.Printf("Nama: %s\n", student.Name)
			fmt.Printf("Alamat: %s\n", student.Address)
			fmt.Printf("Pekerjaan: %s\n", student.Occupation)
			fmt.Printf("Alasan: %s\n", student.Reason)
			break
		}
	}

	if !found {
		fmt.Println("No student data was found")
	}
}
