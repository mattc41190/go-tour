package main

import (
	"fmt"
)

func main() {
	board := [][]string{
		{"-", "-", "-"},
		{"-", "-", "-"},
		{"-", "-", "-"},
	}

	// players take turns
	board[0][1] = "X"
	board[1][0] = "O"
	board[0][0] = "X"
	board[1][1] = "O"
	board[0][2] = "X" // X wins: Top - Across

	for i := 0; i < len(board); i++ {
		fmt.Println(board[i])
	}

}
