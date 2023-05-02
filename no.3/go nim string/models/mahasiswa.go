package models

type Mahasiswa struct {
	Nama   string `json:"nama"`
	NIM    string `json:"nim"`
	Alamat string `json:"alamat"`
}

var DataMahasiswa []Mahasiswa
