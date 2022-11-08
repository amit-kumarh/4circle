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

func Negamax(position Position) int {
	// check for draw
	// to do

	// checking if opponent can win next move
	for i := 0; i < 7; i++ {
		if canPlay(position.position, i) && IsWinningMove(position.position, position.mask, i) {
			return 0 // 0 indicates that they can
		}
	}

	bestScore := 0 // var to store best possible score

	// look for best possible score, save that score in var
	for i := 0; i < 7; i++ {
		if canPlay(position.mask, i) {
			opponentPlay := position
			play(opponentPlay.position, opponentPlay.mask, i)

			opponentScore := -Negamax(opponentPlay)

			if opponentScore > bestScore {
				bestScore = opponentScore
			}
		}
	}

	return bestScore
}
