package controller

import (
	"encoding/json"
	"fmt"
	"net/http"

	"mahasiswa/models"
)

func TambahMahasiswaHandler(w http.ResponseWriter, r *http.Request) {
	var mahasiswa models.Mahasiswa
	err := json.NewDecoder(r.Body).Decode(&mahasiswa)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	models.DataMahasiswa = append(models.DataMahasiswa, mahasiswa)
	fmt.Fprintf(w, "Data mahasiswa berhasil ditambahkan")
}

func TampilkanDataMahasiswaHandler(w http.ResponseWriter, r *http.Request) {
	if len(models.DataMahasiswa) == 0 {
		fmt.Fprintf(w, "Data mahasiswa kosong")
		return
	}

	fmt.Fprintf(w, "Data Mahasiswa:\n")
	for _, mahasiswa := range models.DataMahasiswa {
		fmt.Fprintf(w, "Nama: %s\nNIM: %s\nAlamat: %s\n\n", mahasiswa.Nama, mahasiswa.NIM, mahasiswa.Alamat)
	}
}
