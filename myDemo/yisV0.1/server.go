package main

import (
	"bytes"
	"encoding/binary"
	"fmt"
)

func main() {
	dataBuff := bytes.NewBuffer([]byte{})

	a := 11
	err := binary.Write(dataBuff, binary.LittleEndian, uint32(a))
	if err != nil {
		fmt.Println("write1 error", err)
		return
	}

	b := 99
	err = binary.Write(dataBuff, binary.LittleEndian, uint32(b))
	if err != nil {
		fmt.Println("write2 error")
		return
	}

	data := []byte{'1', '2', '3', '4', '5', '6'}
	err = binary.Write(dataBuff, binary.LittleEndian, data)
	if err != nil {
		fmt.Println("write3 error")
		return
	}

	sli := dataBuff.Bytes()

	fmt.Printf("%v", sli)
}
