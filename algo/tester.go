package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"time"
)

/*
	Tester runs through the test files and computes the average time to run through
	each position in the file and the average nodes explored per position
*/
func Tester() {
	// fmt.Println("Test")
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
			positionString := scanner.Text()
			pos := newPosition()
			sol := newSolver()

			scanner.Scan()
			expectedScore, error := strconv.Atoi(scanner.Text())

			if error != nil {
				log.Fatal(error)
			}

			InitializeBoard(pos, positionString)

			// start timer
			timerStart := time.Now()
			score := Negamax(pos, sol, -22, 22)

			duration := time.Since(timerStart)

			averageNodesExplored += sol.nodesExplored
			averageTime += duration

			if score != expectedScore {
				panic("Actual Score does not equal expected score!")
			}

			fmt.Println("----Final Results----")
			fmt.Println("Score: ", score)
			fmt.Println("Computation Time: ", duration)
			fmt.Println("Amount of Nodes Explored: ", sol.nodesExplored)
		}

		averageTime /= 1000
		averageNodesExplored /= 1000

		fmt.Println("")
		fmt.Println("----Final Results----")
		fmt.Println("Average Computation Time: ", averageTime)
		fmt.Println("Average Amount of Nodes Explored: ", averageNodesExplored)
		file.Close()
	}
}
