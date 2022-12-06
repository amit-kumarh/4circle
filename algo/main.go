package main

// build using go build -buildmode=c-shared -o main.so

import "C"

//export getCol
func getCol(col int) C.int {
	return C.int(col)

}
func main() {
	Tester()
	// init structs for pos and solver
	position := newPosition()
	solver := newSolver()

	for position.moves <= 49 {
		// five steps:
		// 1, Get user input
		// 2, play user move with position bitboard

		// 3, run negamax and get column
		col := Negamax(position, solver, -42, 42)
		// 4, run python function to move to column
		print(col)
		// 5, repeat
	}

}
