package main

import "fmt"

// solver function takes the position and alpha-beta values and evaluates the score

// what is the solver going to do:
// takes a position and returns the score of that position
// then algo will pick position with greatest score.

const NUM_SPACES int = 42

func Negamax(position Position, alpha int, beta int) int {

	// fmt.Println("Running Negamax")
	columnOrder := []int{3, 4, 2, 5, 1, 6, 0}
	position.moves++ // increment num of nodes explored
	// check for draw
	if position.moves == NUM_SPACES {
		return 0
	}

	// checking if we can win next move
	for i := 0; i <= 6; i++ {
		if CanPlay(&position, i) && IsWinningMove(&position, i) {
			// fmt.Println("Can win next move")
			// fmt.Println("Moves: ", position.moves)
			fmt.Println("Col: ", i)
			return 22 - ((position.moves + 1) / 2)
		}
	}

	max := 21 - ((position.moves + 1) / 2)
	if beta > max {
		beta = max
		if alpha >= beta {
			return beta
		}
	}

	// look for best possible score, save that score in var
	for i := 0; i < 7; i++ {
		if CanPlay(&position, columnOrder[i]) {
			to_check := position
			Play(&to_check, columnOrder[i])

			score := -Negamax(to_check, -alpha, -beta)
			// fmt.Println("Score: ", score)

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
