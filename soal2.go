package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	// Membaca input dari user
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Masukkan nama (minimal 3 kata): ")
	input, _ := reader.ReadString('\n')

	// Menghapus spasi di awal/akhir dan memisahkan kata-kata menjadi slice
	input = strings.TrimSpace(input)
	words := strings.Split(input, " ")

	// Memastikan input memiliki minimal 3 kata
	if len(words) < 3 {
		fmt.Println("Input harus terdiri dari minimal 3 kata.")
		return
	}

	// Membalik urutan kata
	for i, j := 0, len(words)-1; i < j; i, j = i+1, j-1 {
		words[i], words[j] = words[j], words[i]
	}

	// Mencetak setiap kata dalam baris baru
	for _, word := range words {
		fmt.Println(word)
	}
}
