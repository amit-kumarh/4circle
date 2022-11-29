package main

const NUM_SPACES int = 42

type Solver struct {
	nodesExplored int
}

func newSolver() *Solver {
	return &Solver{0}
}

func Negamax(position *Position, sol *Solver, alpha int, beta int) int {
	if alpha >= beta {
		panic("Alpha must be less than Beta!")
	}
	columnOrder := []int{3, 4, 2, 5, 1, 6, 0}
	sol.nodesExplored++
	// check for draw
	if position.moves == NUM_SPACES {
		return 0
	}

	// checking if we can win next move
	for i := 0; i <= 6; i++ {
		if CanPlay(position, i) && IsWinningMove(position, i) {
			// fmt.Println("Can win next move")
			return (43 - position.moves) / 2
		}
	}

	max := (41 - position.moves) / 2
	// fmt.Println("Max score: ", max)
	if beta > max {
		beta = max
		if alpha >= beta {
			return beta
		}
	}

	// look for best possible score, save that score in var
	for i := 0; i < 7; i++ {
		// fmt.Println("In for loop")
		if CanPlay(position, columnOrder[i]) {

			to_check := *position
			Play(&to_check, columnOrder[i])

			score := -Negamax(&to_check, sol, -beta, -alpha)
			// fmt.Println("Score: ", score)
			// fmt.Println("Nodes explored: ", position.moves)

			if score >= beta {
				return score
			}
			if score > alpha {
				alpha = score
			}
		}
	}
	return alpha
}

func Solve(position *Position, sol *Solver) int {
	// check if you can win in one move
	return 0
}
