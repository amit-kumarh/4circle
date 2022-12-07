package main

import lru "github.com/hashicorp/golang-lru/v2"

// import "C"

// solver function takes the position and alpha-beta values and evaluates the score

// what is the solver going to do:
// takes a position and returns the score of that position
// then algo will pick position with greatest score.

const NUM_SPACES int = 42

type Solver struct {
	nodesExplored int
	transpo       *lru.Cache[uint64, int]
}

func newSolver() *Solver {
	transpo, _ := lru.New[uint64, int](10000000) // store 10 million entries
	return &Solver{0, transpo}
}

//export Negamax
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
	val, present := sol.transpo.Get(Key(position))
	if present {
		max = val
	}

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
	sol.transpo.Add(Key(position), alpha)
	return alpha
}

func Solve(pos *Position, sol *Solver) int {
	min := -(42 - pos.moves) / 2
	max := (42 - pos.moves) / 2

	for min < max {
		med := min + (max-min)/2
		if (med <= 0) && ((min / 2) < med) {
			med = min / 2
		} else if (med >= 0) && ((max / 2) > 2) {
			med = max / 2
		}

		r := Negamax(pos, sol, med, med+1) // use a null depth window to know if score is greater than or less than med
		if r <= med {
			max = r
		} else {
			min = r
		}
	}
	return min
}

func Score(pos *Position) [7]int {
	for col := 0; col < 7; col++ {
		test := pos
		test.position |= (test.mask + bottomMask(col)) & columnMask(col)
		score = ComputeWinningPosition(test)

	}
}
