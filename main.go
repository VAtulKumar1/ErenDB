package main

import (
	"fmt"

	"github.com/VAtulKumar1/ErenDB/chapter1"
)

func main() {
	fmt.Println("Hello from Atul1")
	str := "Hello from go"
	byteArray := []byte(str)
	chapter1.SaveData1("fileDB.txt", byteArray)
}
