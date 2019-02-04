package main

import (
	"errors"
	"fmt"
	"io"
	"os"
	"strings"
)

var lower = []rune{
	'a', 'b', 'c', 'd', 'e', 'f', 'g', 'h', 'i', 'j', 'k', 'l', 'm', 'n', 'o', 'p', 'q', 'r', 's', 't', 'u', 'v', 'w', 'x', 'y', 'z',
}

var upper = []rune{
	'A', 'B', 'C', 'D', 'E', 'F', 'G', 'H', 'I', 'J', 'K', 'L', 'M', 'N', 'O', 'P', 'Q', 'R', 'S', 'T', 'U', 'V', 'W', 'X', 'Y', 'Z',
}

type rot13Reader struct {
	s io.Reader
}

func main() {
	// "Lbh penpxrq gur pbqr!" == You cracked the code!
	s := strings.NewReader("Uryyb Obqvr, lbh ovt ornhgvshy zhg")
	r := rot13Reader{s}
	io.Copy(os.Stdout, &r)
	fmt.Println("")
}

func (rr *rot13Reader) Read(bs []byte) (int, error) {
	numRead, err := rr.s.Read(bs) // bs can be read into because slices are just pointers to arrays
	if err == io.EOF {
		return 0, err
	}
	for i := 0; i < numRead; i++ {
		bs[i] = byte(rotate(rune(bs[i])))
	}
	return numRead, err
}

func rotate(letter rune) rune {
	if 'A' <= letter && letter <= 'Z' {
		shiftedRune, err := shiftRune(len(upper)/2, upper, letter)
		if err != nil {
			panic(err)
		}
		return shiftedRune
	}
	if 'a' <= letter && letter <= 'z' {
		shiftedRune, err := shiftRune(len(upper)/2, lower, letter)
		if err != nil {
			panic(err)
		}
		return shiftedRune
	}
	return letter
}

func indexOf(s []rune, r rune) (int, error) {
	for i := 0; i < len(s); i++ {
		if s[i] == r {
			return i, nil
		}
	}
	return 0, errors.New("no matching rune found")
}

func shiftRune(places int, source []rune, r rune) (rune, error) {
	length := len(source)
	index, err := indexOf(source, r)

	if err != nil {
		return 0, err
	}

	if index+places >= length {
		overflow := index + places - length
		return source[overflow], nil
	}

	return source[index+places], nil
}
