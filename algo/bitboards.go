package main

import (
	"fmt"
	"strconv"
)

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

	test_pos |= (pos.mask + bottomMask(col)) & columnMask(col)
	fmt.Println("Test pos: ", strconv.FormatInt(int64(test_pos), 2))
	return Aligned(test_pos)
}

// return a bitmask containg a single 1 corresponding to the top cell of a given column
func topMask(col int) uint64 {
	// fmt.Println("Top Mask: ", strconv.FormatInt(((1<<5)<<col*7), 2))
	return (1 << 5) << col * 7
}

// return a bitmask containg a single 1 corresponding to the bottom cell of a given column
func bottomMask(col int) uint64 {
	// fmt.Println("Bottom Mask: ", strconv.FormatInt((1<<col*7), 2))
	return 1 << col * 7
}

// return a bitmask 1 on all the cells of a given column
func columnMask(col int) uint64 {
	// fmt.Println("Column Mask: ", strconv.FormatInt(((1<<5)<<col*7), 2))
	return (1 << 5) << col * 7
}
func Aligned(pos uint64) bool {
	// Horizontal
	inter := pos & (pos >> 7)
	if (inter & (inter >> 14)) != 0 {
		fmt.Println("Horizontal win: ", strconv.FormatInt(int64(inter&(inter>>14)), 2))

		return true
	}

	// Vertical
	inter = pos & (pos >> 1)
	if (inter & (inter >> 2)) != 0 {
		fmt.Println("Vertical win: ", strconv.FormatInt(int64(inter&(inter>>2)), 2))
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
