package main

import "fmt"

type MyReader struct{}

func (reader MyReader) Read(bs []byte) (int, error) {
	for i := 0; i < len(bs); i++ {
		bs[i] = 'A'
	}
	return len(bs), nil
}

func main() {
	bs := make([]byte, 8)
	r := MyReader{}
	r.Read(bs)
	for _, b := range bs {
		fmt.Println(string(int(b)))
	}
}

// Initial Working Guess:
// func (reader MyReader) Read(byteSlice []byte) (int, error) {
// 	var numRead int
// 	for i := 0; i < len(byteSlice); i++ {
// 		char := 'A' // rune, not string // <-- https://www.socketloop.com/tutorials/golang-how-to-convert-character-to-ascii-and-back
// 		ascii := int(char)
// 		byteSlice[i] = byte(ascii)
// 		numRead++
// 	}
// 	return numRead, nil
// }
