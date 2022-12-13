package main

type Position struct {
	position uint64
	mask     uint64
	moves    int
}

const BOARD_MASK uint64 = 0b0000000000000000111111011111101111110111111011111101111110111111
const BOTTOM_MASK uint64 = 0b0000000000000000000001000000100000010000001000000100000010000001

//export newPosition
func newPosition() *Position {
	pos := &Position{0, 0, 0}
	return pos
}

func CanPlay(pos *Position, col int) bool {
	return (pos.mask & topMask(col)) == 0
}

func possible(pos *Position) uint64 {
	return (pos.mask + BOTTOM_MASK) & BOARD_MASK
}

func InitializeBoard(pos *Position, seq string) {
	for i := 0; i < len(seq); i++ {
		colByte := seq[i] - '1'
		// fmt.Println(colByte)
		col := int(colByte)
		if col < 0 || col > 7 || !CanPlay(pos, col) || IsWinningMove(pos, col) {
			return
		}
		Play(pos, col)
	}
}

func Play(pos *Position, col int) {
	pos.position ^= pos.mask
	pos.mask |= pos.mask + bottomMask(col)
	pos.moves++
}

func Key(pos *Position) uint64 {
	return pos.position + pos.mask
}

func IsWinningMove(pos *Position, col int) bool {
	test_pos := pos.position

	test_pos |= (pos.mask + bottomMask(col)) & columnMask(col)
	// fmt.Println("Test pos: ", strconv.FormatInt(int64(test_pos), 2))
	return Aligned(test_pos)
}

// return a bitmask containing a single 1 corresponding to the top cell of a given column
func topMask(col int) uint64 {
	return (1 << 5) << (col * 7)
}

// return a bitmask containing a single 1 corresponding to the bottom cell of a given column
func bottomMask(col int) uint64 {
	return 1 << (col * 7)
}

// return a bitmask 1 on all the cells of a given column
func columnMask(col int) uint64 {
	return ((1 << 6) - 1) << (col * 7)
}

func Aligned(pos uint64) bool {

	// Horizontal
	inter := pos & (pos >> 7)
	if (inter & (inter >> 14)) != 0 {
		return true
	}

	// Vertical
	inter = pos & (pos >> 1)
	if (inter & (inter >> 2)) != 0 {
		return true
	}

	// Diagonal Down
	inter = pos & (pos >> 6)
	if (inter & (inter >> 12)) != 0 {
		return true
	}

	// Diagonal Up
	inter = pos & (pos >> 8)
	if (inter & (inter >> 16)) != 0 {

		return true
	}

	return false
}

// returns the popcount of compute winning position
func ComputeWinningPosition(pos *Position) uint64 {

	// vertical wins
	winningMovesBitmask := (pos.position << 1) & (pos.position << 2) & (pos.position << 3)

	// horizontal win
	// check for a win to the directly to the left
	shiftedPosition := (pos.position << 7) & (pos.position << 14)
	winningMovesBitmask |= shiftedPosition & (pos.position << 21)
	// check for a win in the middle (1 1 0 1) where 0 is our winning space
	winningMovesBitmask |= shiftedPosition & (pos.position >> 7)

	// check for right wins
	shiftedPosition = (pos.position >> 7) & (pos.position >> 14)
	winningMovesBitmask |= shiftedPosition & (pos.position << 7)
	winningMovesBitmask |= shiftedPosition & (pos.position >> 21)

	// check for diagonal wins:

	// downwards diagonal
	shiftedPosition = (pos.position << 6) & (pos.position << 12)
	winningMovesBitmask |= shiftedPosition & (pos.position << 18)
	winningMovesBitmask |= shiftedPosition & (pos.position >> 6)

	// upwards diagonal
	shiftedPosition = (pos.position << 8) & (pos.position << 16)
	winningMovesBitmask |= shiftedPosition & (pos.position << 24)
	winningMovesBitmask |= shiftedPosition & (pos.position >> 8)

	winningMovesBitmask &= (BOARD_MASK ^ pos.mask)
	return winningMovesBitmask
}

func OpponentWinningMoves(pos *Position) uint64 {
	return ComputeWinningPosition(&Position{pos.position ^ pos.mask, pos.mask, 1})
}

func NumOpponentWinningMoves(pos *Position) int {
	winningMovesBitmask := OpponentWinningMoves(pos)
	var counter uint64
	counter = 0
	for counter < winningMovesBitmask {
		winningMovesBitmask &= (winningMovesBitmask - 1)
		counter++
	}
	// return popcount
	return int(counter)
}

func NumWinningMoves(pos *Position) int {
	winningMovesBitmask := ComputeWinningPosition(pos)
	var counter uint64
	counter = 0
	for counter < winningMovesBitmask {
		winningMovesBitmask &= (winningMovesBitmask - 1)
		counter++
	}
	// return popcount
	return int(counter)
}
