package main

import (
	"fmt"
	"os"

	"piscine"

	"github.com/01-edu/z01"
)

func main() {
	// Basic size we're working with (9x9)
	const MAX int = 9

	// Quit the program if we don't have the correct number of arguments
	if len(os.Args) != MAX+1 {
		fmt.Println("Error")
		return
	}

	// Check if the given arguments are valid for Sudoku
	for index, argument := range os.Args {
		if index != 0 && !piscine.ArgIsValid(argument) {
			fmt.Println("Error")
			return
		}
	}

	// Creation of the Sudoku 9x9 matrix
	sudoku := make([][]rune, MAX)
	copy := make([][]rune, MAX)
	/*
		Creation of the 9x9 base matrix to store
		the given value in order not to alter them
	*/
	base := make([][]rune, MAX)
	for i := range sudoku {
		sudoku[i] = make([]rune, MAX)
		copy[i] = make([]rune, MAX)
		base[i] = make([]rune, MAX)
	}

	// Filling the matrices
	for i := range sudoku {
		arg := []rune(os.Args[i+1])
		for j := range sudoku[i] {
			sudoku[i][j] = arg[j]
			copy[i][j] = arg[j]
			if arg[j] >= '1' && arg[j] <= '9' {
				base[i][j] = '1'
			} else {
				base[i][j] = '0'
			}
		}
	}

	// Checking if the given values to start the Sudoku are valid
	for i := range sudoku {
		for j := range sudoku[i] {
			if sudoku[i][j] != '.' && !piscine.CanInsert(sudoku[i][j], i, j, sudoku) {
				fmt.Println("Error")
				return
			}
		}
	}

	// Solution
	counter1 := 0
	counter2 := 0
	insert := 1
	for counter1 >= 0 && counter1 < MAX {
		for counter2 >= 0 && counter2 < MAX {
			// for i := range sudoku {
			// 	rowEnd := len(sudoku[i]) - 1
			// 	for j := range sudoku[i] {
			// 		z01.PrintRune(sudoku[i][j])
			// 		if j < rowEnd {
			// 			z01.PrintRune(' ')
			// 		}
			// 	}
			// 	z01.PrintRune('\n')
			// }
			// z01.PrintRune('\n')
			if base[counter1][counter2] == '0' {
				for insert <= MAX && !piscine.CanInsert(rune('0'+insert), counter1, counter2, sudoku) {
					insert++
				}
				if insert > MAX {
					sudoku[counter1][counter2] = '.'
					counter2--
					for counter2 >= 0 && base[counter1][counter2] != '0' {
						counter2--
					}
					if counter2 >= 0 {
						insert = int((sudoku[counter1][counter2] + 1) % '0')
					}
				} else {
					sudoku[counter1][counter2] = rune('0' + insert)
					insert = 1
					counter2++
				}
			} else {
				insert = 1
				counter2++
			}
		}
		if counter2 < 0 {
			counter1--
			counter2 = MAX - 1
			for counter1 >= 0 && counter2 >= 0 && (base[counter1][counter2] != '0' || sudoku[counter1][counter2] == '.') {
				counter2--
			}
		} else {
			counter1++
			counter2 = 0
		}
	}

	if counter1 < 0 {
		fmt.Println("Error")
		return
	}

	for i := range sudoku {
		rowEnd := len(sudoku[i]) - 1
		for j := range sudoku[i] {
			z01.PrintRune(sudoku[i][j])
			if j < rowEnd {
				z01.PrintRune(' ')
			}
		}
		z01.PrintRune('\n')
	}
}
