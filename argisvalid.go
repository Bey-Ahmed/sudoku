package piscine

/* 
* Check if a given argument fill the criterias for a valid row in Sudoku
*/
func ArgIsValid(argument string) bool {
	arg := []rune(argument)

	// Not valid if less than 9 characters have been given
	if len(arg) != 9 {
		return false
	}

	for _, digit := range arg {
		// Not valid if the characters are not numerical nor a dot
		if (digit < '1' || digit > '9') && digit != '.' {
			return false
		}
	}

	return true
}
