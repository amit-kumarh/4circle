package main

import (
	"fmt"
	"os/exec"
	"strconv"
)

func main() {
	// Tester()
	// init structs for pos and solver
	position := newPosition()
	solver := newSolver()
	InitializeBoard(position, "2323445532")
	var userInput int
	pythonLine := "import serial; print(serial.to_serial("
	for position.moves <= 49 {

		// five steps:
		// 1, Get user input
		fmt.Print("User Input: ") // for debug purposes :D
		fmt.Scanln(&userInput)

		// check if user Input is winning move:
		if IsWinningMove(position, userInput) {
			fmt.Println("Opponent Wins")
			break
		}
		// 2, play user move with position bitboard
		Play(position, userInput)
		// 3, run negamax and get column
		fmt.Println("Running Negamax")
		col := Negamax(position, solver, -42, 42) // NEED TO EDIT HERE
		// 4, run python function to move to column
		pythonLine += strconv.Itoa(col) + "))"

		// 5, repeat
		fmt.Println("Running Python")
		cmd := exec.Command("python3", "-c", pythonLine)
		// fmt.Println(cmd.Args)

		out, err := cmd.CombinedOutput()
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println("Column: ", string(out))
	}

}
