package main

import (
	"fmt"
	"os/exec"
	"strconv"
)

func main() {
	// init structs for pos and solver
	position := newPosition()
	solver := newSolver()
	InitializeBoard(position, "2323445532")

	var userInput int
	for position.moves <= 49 {

		// five steps:
		// 1, Get user input
		fmt.Print("User Input: ") // for debug purposes :D
		fmt.Scanln(&userInput)
		userInput -= 1

		// check if user Input is winning move:
		if IsWinningMove(position, userInput) {
			fmt.Printf("%b", position.mask)
			fmt.Println("Player Wins")
			break
		}
		// 2, play user move with position bitboard
		Play(position, userInput)
		// 3, run negamax and play and get column
		fmt.Println("Running Negamax")
		col := bestMove(position, solver)

		// 4, run python function to move to column
		pythonLine := "import serial; print(serial.to_serial(" + strconv.Itoa(col) + "))"

		// 5, repeat
		fmt.Println("Running Python")
		cmd := exec.Command("python3", "-c", pythonLine)
		// fmt.Println(cmd.Args)

		_, err := cmd.CombinedOutput()
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println("Column: ", col+1)

		if IsWinningMove(position, col) {
			fmt.Println("AI Wins")
			break
		}
		Play(position, col)
	}
}
