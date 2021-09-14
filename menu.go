package main

import (
	"errors"
	"fmt"
	"strings"
)

func validation(input string) (bool, error) {
	if strings.TrimSpace(input) == "" {
		return false, errors.New("pesanan tidak boleh kosong! silahkan masukan pesanan anda")
	}

	return true, nil
}

func main() {
	var pesanan []string
	var menu, next string

	fmt.Println("============================")
	fmt.Println("Toko Makanan Indonesia")
	fmt.Println("============================")
	fmt.Println("Tahu \nTempe \nSambal \nNasi \nLele \nAyam")
	fmt.Println("============================")

pilihan:
	fmt.Print("Masukan menu pesanan anda dalam huruf [ eg: nasi] : ")
	fmt.Scanln(&menu)

	if valid, err := validation(menu); valid == false {
		fmt.Println(err)
		goto pilihan
	} else {
		pesanan = append(pesanan, menu)
	}

pesanLagi:
	fmt.Print("Lanjutkan Memesan? [Y/T]: ")
	fmt.Scanln(&next)

	if next == "y" || next == "Y" {
		goto pilihan

	} else if next == "t" || next == "T" {
		for _, pesan := range pesanan {
			fmt.Print("Pesanan Anda : ", pesan, "\n")
		}
	} else {
		goto pesanLagi
	}

}
