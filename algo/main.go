package main

func main() {
	// Tester()
	pos := newPosition()
	// sol := newSolver()

	InitializeBoard(pos, "32164625")
	ComputeWinningPosition(pos)
}
