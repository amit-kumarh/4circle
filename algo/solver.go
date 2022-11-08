package main

// solver function takes the position and alpha-beta values and evaluates the score

// what is the solver going to do:
// takes a position and returns the score of that position
// then algo will pick position with greatest score.

type Position struct {
	position   uint64
	mask       uint64
	moveNumber int
}

func possibleNonLosingMoves() Position {

}

func Negamax(position Position, alpha int, beta int) int {

	if beta < alpha {
		panic("Alpha should be less than beta!")
	}

	position.moveNumber++ // increment number of nodes explored

	return -1
}
