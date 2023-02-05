package piscine

func CanInsert(insert rune, xAxis, yAxis int, sudoku [][]rune) bool {
	// Not possible if the given coordinates are beyond our sudoku matrix size
	if xAxis >= len(sudoku) || yAxis >= len(sudoku[xAxis]) {
		return false
	}

	// Not possible if the value to insert is already in the **row**
	for column := range sudoku[xAxis] {
		if column != yAxis && sudoku[xAxis][column] == insert {
			return false
		}
	}

	// Not possible if the value to insert is already in the **column**
	for row := 0; row < len(sudoku[xAxis]); row++ {
		if row != xAxis && sudoku[row][yAxis] == insert {
			return false
		}
	}

	// Checking the box coordinates from the given coordinates
	boxPosX := 0
	boxPosY := 0
	if xAxis < 3 {
		boxPosX = 0
	} else if xAxis < 6 {
		boxPosX = 1
	} else if xAxis < 9 {
		boxPosX = 2
	}

	if yAxis < 3 {
		boxPosY = 0
	} else if yAxis < 6 {
		boxPosY = 1
	} else if yAxis < 9 {
		boxPosY = 2
	}

	// Not possible if the value to insert is already in the **box**
	for i := 0 + boxPosX*3; i < 3+boxPosX*3; i++ {
		for j := 0 + boxPosY*3; j < 3+boxPosY*3; j++ {
			if xAxis != i && yAxis != j && sudoku[i][j] == insert {
				return false
			}
		}
	}

	return true
}
