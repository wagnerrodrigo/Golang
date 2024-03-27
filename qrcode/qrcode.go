package main

import (
	"os"

	qrcode "github.com/skip2/go-qrcode"
)

func main() {
	data := os.Args[1]

	qrcode, _ := qrcode.Encode(data, qrcode.Highest, 256)

	file, _ := os.Create("qrcode.png")
	defer file.Close()

	file.Write(qrcode)

}
