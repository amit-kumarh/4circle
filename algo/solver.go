package main

import "fmt"

// solver function takes the position and alpha-beta values and evaluates the score

// what is the solver going to do:
// takes a position and returns the score of that position
// then algo will pick position with greatest score.

const NUM_SPACES int = 42

func Negamax(position Position, alpha int, beta int) int {
	// check for draw
	if position.moves == NUM_SPACES {
		return 0
	}

	// checking if we can win next move
	for i := 0; i < 7; i++ {
		if CanPlay(&position, i) && IsWinningMove(&position, i) {
			return 21 - position.moves/2
		}
	}

	max := 20 - (position.moves / 2)
	if beta > max {
		beta = max
		if alpha >= beta {
			return beta
		}
	}

	// look for best possible score, save that score in var
	for i := 0; i < 7; i++ {
		if CanPlay(&position, i) {
			to_check := position
			Play(&to_check, i)

			score := -Negamax(to_check, -alpha, -beta)

			fmt.Println("score: ", score)

			if score >= beta {
				return beta
			}
			if score > alpha {
				alpha = score
			}
		}
	}

	return alpha
}
