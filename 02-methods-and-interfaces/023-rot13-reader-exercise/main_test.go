package main

import (
	"testing"
)

func TestIndexOfValid(t *testing.T) {
	s := []rune{'a', 'b', 'c'}
	r := 'b'
	expected := 1

	i, err := indexOf(s, r)

	if err != nil {
		t.Error(err)
	}

	if i != expected {
		t.Error(i)
	}
}

func TestIndexOfInvalid(t *testing.T) {
	s := []rune{'a', 'b', 'c'}
	r := 'd'

	_, err := indexOf(s, r)

	if err == nil {
		t.Error(err)
	}
}

func TestGetShiftedChar(t *testing.T) {
	s := []rune{'a', 'b', 'c', 'd', 'e', 'f'}
	table := []struct {
		i rune
		o rune
	}{
		{
			i: 'a',
			o: 'd',
		},
		{
			i: 'e',
			o: 'b',
		},
		{
			i: 'c',
			o: 'f',
		},
		{
			i: 'f',
			o: 'c',
		},
		{
			i: 'd',
			o: 'a',
		},
	}

	for _, test := range table {

		shifted, err := shiftRune(len(s)/2, s, test.i)

		if err != nil {
			t.Error(err)
		}

		if shifted != test.o {
			t.Errorf("actual %v expected %v", shifted, test.o)
		}
	}

}
