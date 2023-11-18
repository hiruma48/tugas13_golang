package main

import (
	"crypto/sha1"
	"encoding/base64"
	"fmt"
)

func main() {
	var nama_saya = "tri ariyanto"
	var sha = sha1.New()

	var encodestring = base64.StdEncoding.EncodeToString([]byte(nama_saya))
	fmt.Println("nama saya : ", encodestring)
	sha.Write([]byte(nama_saya))
	var enkripsi = sha.Sum(nil)
	var stringenkripsi = fmt.Sprintf("%x", enkripsi)
	fmt.Println("enkripsi adalah : ", stringenkripsi)

}
