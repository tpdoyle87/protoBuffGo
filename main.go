package main

import (
	"encoding/binary"
	"fmt"
	"os"
)

func main() {
	test()
}
func encode(num uint64) []byte {
	var b []byte
	for num > 0 {
		p := num & 0x7f
		num >>= 7
		if num > 0 {
			p |= 0x80
		}
		b = append(b, byte(p))
	}
	return b
}

func decode(b []byte) uint64 {
	var num uint64
	var offset int
	for i := 0; i < len(b); i++ {
		x := 0xff & b[i]
		num |= uint64(x) << offset
		offset += 7
	}
	return num
}

func test() {
	var files = []string{"150.uint64", "1.uint64", "maxint.uint64"}
	fmt.Println("**** Starting Test ****")
	for _, file := range files {
		f, err := os.Open("varint/" + file)
		if err != nil {
			fmt.Println("Error reading file: ", err)
		}

		buffer := make([]byte, 8)
		_, err = f.Read(buffer)
		if err != nil {
			fmt.Println("Error reading file: ", err)
		}
		err = f.Close()
		if err != nil {
			fmt.Println("Error closing file: ", err)
		}
		number := binary.BigEndian.Uint64(buffer)
		encoded := encode(number)
		fmt.Printf("Encoded value: %x\n", encoded)
		decoded := decode(encoded)
		fmt.Printf("Decoded value: %d\n", decoded)
		fmt.Println("Roundtrip test: ", number == decoded)
		fmt.Println("---------------------------------")

	}
	fmt.Println("**** Test Complete ****")

}
