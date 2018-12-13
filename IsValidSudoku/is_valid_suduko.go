package main

import (
	"log"
)

func isValidSudoku(board [][]byte) bool {
	//row
	for i := 0; i < 9; i ++ {
		resM := make(map[byte]string, 0)
		index := 0
		for j := 0; j < 9; j ++ {
			if board[i][j] != '.' {
				index ++
				resM[board[i][j]] = string(board[j][i])
			}
		}
		if len(resM) != index {
			return false
		}
	}
	//list
	for i := 0; i < 9; i ++ {
		resM := make(map[string]string, 0)
		index := 0
		for j := 0; j < 9; j ++ {
			if board[j][i] != '.' {
				index ++
				resM[string(board[j][i])] = string(board[j][i])
			}
		}
		if len(resM) != index {
			return false
		}
	}
	//block
	for i := 0; i < 3; i ++ {
		for j := 0; j < 3; j ++ {
			resM := make(map[byte]string, 0)
			index := 0
			for k := 0; k < 3; k ++ {
				for l := 0; l < 3; l ++ {
					v := i*3 + k
					b := j*3 + l
					if board[v][b] != '.' {
						index ++
						resM[board[v][b]] = string(board[j][i])
					}
				}
			}
			if len(resM) != index {
				return false
			}
		}
	}
	return true
}

func main() {
	ss := [][]byte{
		{'5', '3', '.', '.', '7', '.', '.', '.', '.'},
		{'6', '.', '.', '1', '9', '5', '.', '.', '.'},
		{'.', '9', '8', '.', '.', '.', '.', '6', '.'},
		{'8', '.', '.', '.', '6', '.', '.', '.', '3'},
		{'4', '.', '.', '8', '.', '3', '.', '.', '1'},
		{'7', '.', '.', '.', '2', '.', '.', '.', '6'},
		{'.', '6', '.', '.', '.', '.', '2', '8', '.'},
		{'.', '.', '.', '4', '1', '9', '.', '.', '5'},
		{'.', '.', '.', '.', '8', '.', '.', '7', '9'}}
	log.Println(isValidSudoku(ss))
}
