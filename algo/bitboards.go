package main

import "fmt"

type Position struct {
	position uint64
	mask     uint64
	moves    int
}

func CanPlay(pos *Position, col int) bool {
	return (pos.mask)&((1<<5)<<(col*7)) == 0
}

func Play(pos *Position, col int) {
	pos.position ^= pos.mask
	pos.mask |= pos.mask + (1 << (col * 7))
	pos.moves++
}

func Key(pos uint64, mask uint64) uint64 {
	return pos + mask
}

func IsWinningMove(pos *Position, col int) bool {
	test_pos := pos.position
	test_pos |= (pos.mask + (1 << col * 7)) & (63 << col * 7)
	fmt.Println("Test pos: ", test_pos)
	return Aligned(test_pos)
}

func Aligned(pos uint64) bool {
	// Horizontal
	inter := pos & (pos >> 7)
	if (inter & (inter >> 14)) != 0 {
		fmt.Println("Horizontal win")
		return true
	}

	// Vertical
	inter = pos & (pos >> 1)
	if (inter & (inter >> 2)) != 0 {
		fmt.Println("Vertical win")
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
