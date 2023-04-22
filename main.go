package main

import (
	"fmt"
	"io"
	"os"
	"strings"
)

func main() {
	var url, token = Credentials()
	print(url, token)
}

func Credentials() (string, string) {
	var url = ""
	var token = ""

	file, err := os.Open("creds.txt")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer file.Close()

	data := make([]byte, 256)

	for {
		n, err := file.Read(data)

		if err == io.EOF { // если конец файла
			break // выходим из цикла
		}

		if strings.Contains(string(data[:n]), "url") {
			url = string(data[:n])
		}

		if strings.Contains(string(data[:n]), "token") {
			token = string(data[:n])
		}
	}

	return url, token
}
