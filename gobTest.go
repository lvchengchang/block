package main

import (
	"bytes"
	"encoding/gob"
	"fmt"
)

type Person struct {
	Name    string
	Age     uint
	Address string
}

func main() {
	var chang Person
	chang.Name = "chang"
	chang.Age = 22
	chang.Address = "anhui"

	var buffer bytes.Buffer

	// define buffer
	encode := gob.NewEncoder(&buffer)
	encode.Encode(chang)

	// define decoder buffer

	decode := gob.NewDecoder(bytes.NewBuffer(buffer.Bytes()))

	var test Person
	decode.Decode(&test)
	fmt.Println(test.Address)
}
