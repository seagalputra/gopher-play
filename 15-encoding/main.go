package main

import (
	"encoding/base64"
	"fmt"
)

func main() {

	var data = "john wick"

	var encodedString = base64.StdEncoding.EncodeToString([]byte(data))
	fmt.Println("encoded:", encodedString)

	decodedByte, _ := base64.StdEncoding.DecodeString(encodedString)
	decodedString := string(decodedByte)
	fmt.Println("decoded:", decodedString)

	encoded := make([]byte, base64.StdEncoding.EncodedLen(len(data)))
	base64.StdEncoding.Encode(encoded, []byte(data))
	str := string(encoded)
	fmt.Println("encoded:", str)

	decoded := make([]byte, base64.StdEncoding.DecodedLen(len(encoded)))
	_, err := base64.StdEncoding.Decode(decoded, encoded)
	if err != nil {
		fmt.Println(err.Error())
	}
	decodedStr := string(decoded)
	fmt.Println("decoded:", decodedStr)

	url := "https://kalipare.com/"
	encodedUrl := base64.URLEncoding.EncodeToString([]byte(url))
	fmt.Println(encodedUrl)

	decodeUrl, _ := base64.URLEncoding.DecodeString(encodedUrl)
	fmt.Println(string(decodeUrl))
}
