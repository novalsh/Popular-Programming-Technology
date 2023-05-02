package main

import (
	"fmt"
	"log"
	"net/http"

	controller "mahasiswa/controllers"
)

func main() {
	http.HandleFunc("/nama", controller.TambahMahasiswaHandler)
	http.HandleFunc("/semuadata", controller.TampilkanDataMahasiswaHandler)
	fmt.Println("Server berjalan pada port 1000")
	log.Fatal(http.ListenAndServe(":1000", nil))

}
