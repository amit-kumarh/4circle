package main

import "fmt"

func canPlay(mask uint64, col int) bool {
	return (mask)&((1<<5)<<(col*7)) == 0
}

func play(pos uint64, mask uint64, col int) (uint64, uint64) {
	pos = pos ^ mask
	mask |= mask + 1<<(col*7)
	return pos, mask
}

func key(pos uint64, mask uint64) uint64 {
	return pos + mask
}

func aligned(pos uint64) bool {
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

func main() {
	fmt.Println("vim-go")
}
