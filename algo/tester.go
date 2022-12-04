// // ideally takes a data set
// // plays the positions
// // evaluates time to sovle

package main

// // package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"time"
)

func Tester() [][]int {
	masterTimer := time.Now()
	var max time.Duration
	// files := []string{"Test_L1_R1.txt", "Test_L1_R2.txt", "Test_L1_R3.txt", "Test_L2_R1.txt", "Test_L2_R2.txt", "Test_L3_R1.txt"}
	files := []string{"Test_L2_R2.txt", "Test_L3_R1.txt"}
	var averageTime time.Duration
	averageNodesExplored := 0

	for i := 0; i < len(files); i++ {
		file, err := os.Open(files[i])

		if err != nil {
			log.Fatal(err)
		}

		scanner := bufio.NewScanner(file)
		scanner.Split(bufio.ScanWords)

		for scanner.Scan() {
			position := scanner.Text()
			pos := newPosition()
			sol := newSolver()

			scanner.Scan()
			// fmt.Println("Position: ", position)
			_, error := strconv.Atoi(scanner.Text())

			if error != nil {
				log.Fatal(error)
			}

			InitializeBoard(pos, position)

			// fmt.Println("Position Bitstring: ", pos.position)
			// fmt.Println("Mask Bitstring: ", pos.mask)
			// start timer
			timerStart := time.Now()
			score := Solve(pos, sol)

			duration := time.Since(timerStart)
			if duration > max {
				max = duration
			}

			// fmt.Println("")
			// fmt.Println("----Final Results----")
			fmt.Println("score: ", score)
			// fmt.Println("expected score: ", expectedScore)
			// if score != expectedScore {
			// 	panic("Actual Score does not equal expected score!")
			// }
			fmt.Println("Time: ", duration)
		}

		averageTime /= 1000
		averageNodesExplored /= 1000

		fmt.Println("")
		fmt.Println("----Final Results----")
		fmt.Println("Average Computation Time: ", averageTime)
		fmt.Println("Average Amount of Nodes Explored: ", averageNodesExplored)
		file.Close()
	}
	elapsed := time.Since(masterTimer)
	fmt.Println(elapsed / 1000)
	fmt.Println(max)
	return testData
}
