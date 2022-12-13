package main

import (
	"fmt"
	"sort"

	lru "github.com/hashicorp/golang-lru/v2"
)

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

	nextMoves := NonLosingMoves(position)
	if nextMoves == 0 {
		return -(42 - position.moves) / 2
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

	columnOrder := ColumnOrder(position)

	// columnOrder := []int{3, 4, 2, 5, 1, 6, 0}

	// look for best possible score, save that score in var
	for i := 0; i < 7; i++ {
		// fmt.Println("In for loop")
		if (nextMoves & columnMask(columnOrder[i])) > 0 {
			// if CanPlay(position, columnOrder[i]) {

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

func ColumnOrder(pos *Position) []int {
	scores := []int{0, 0, 0, 0, 0, 0, 0}
	keys := []int{3, 4, 2, 5, 1, 6, 0}

	for i := 0; i < len(keys); i++ {
		test := *pos
		test.position |= (test.mask + bottomMask(keys[i])) & columnMask(keys[i])
		scores[keys[i]] = NumWinningMoves(&test)
	}

	sort.SliceStable(keys, func(i, j int) bool {
		return scores[keys[i]] > scores[keys[j]]
	})
	return keys
}

func NonLosingMoves(pos *Position) uint64 {
	possibleMask := possible(pos)
	opponentWin := OpponentWinningMoves(pos)
	forcedMoves := possibleMask & opponentWin
	if forcedMoves > 0 {
		if (forcedMoves & (forcedMoves - 1)) > 0 {
			return 0
		} else {
			return forcedMoves
		}
	}
	return possibleMask & ^(opponentWin >> 1)
}

func bestMove(pos *Position, sol *Solver) int {
	bestCol := 0
	bestScore := -42
	// look for best possible score, save that score in var
	for i := 0; i < 7; i++ {
		score := 0
		if CanPlay(pos, i) && IsWinningMove(pos, i) {
			// fmt.Println("Can win next move")
			score = (43 - pos.moves) / 2
		} else if CanPlay(pos, i) {
			to_check := *pos
			Play(&to_check, i)
			score = -Negamax(&to_check, sol, -22, 22)
		}
		fmt.Println(i+1, score)
		if score > bestScore {
			bestScore = score
			bestCol = i
		}
	}
	return bestCol
}
