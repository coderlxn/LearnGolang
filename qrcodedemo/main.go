package main

import "fmt"
import "github.com/skip2/go-qrcode"

func main()  {
	//var png []byte
	//png, err := qrcode.Encode("https://example.org", qrcode.Medium, 256)

	err := qrcode.WriteFile("All examples use the qrcode.Medium error Recovery Level and create a fixed 256x256px size QR Code. The last function creates a white on black instead of black on white QR Code.The maximum capacity of a QR Code varies according to the content encoded and the error recovery level. The maximum capacity is 2,953 bytes, 4,296 alphanumeric characters, 7,089 numeric digits, or a combination of these.", qrcode.Medium, 256, "qr.png")
	if err == nil {
		fmt.Println("generate ok")
	} else {
		fmt.Println(err)
	}
}
