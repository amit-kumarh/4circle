// ideally takes a data set
// plays the positions
// evaluates time to sovle

package fourcircle

// package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"time"
)

func tester() [][]int {
	// files := []string{"Test_L1_R1.txt", "Test_L1_R2.txt", "Test_L1_R3.txt", "Test_L2_R1.txt", "Test_L2_R2.txt", "Test_L3_R1.txt"}
	files := []string{"Test_L1_R1.txt"}
	testData := make([][]int, len(files))

	for i := 0; i < len(files); i++ {
		file, err := os.Open(files[i])

		if err != nil {
			log.Fatal(err)
		}

		// defer file.Close()
		scanner := bufio.NewScanner(file)
		scanner.Split(bufio.ScanWords)

		// averageScore := 0
		// averageTime := 0
		// averageNodes := 0

		for scanner.Scan() {
			position := scanner.Text()

			scanner.Scan()

			expectedScore, error := strconv.Atoi(scanner.Text())

			if error != nil {
				log.Fatal(error)
			}

			fmt.Println("Position: ", position)
			// fmt.Println("Expected Score: ", expectedScore)

			pos, err := strconv.ParseUint(position, 10, 64)
			if err != nil {
				log.Fatal(err)
			}

			// start timer
			timerStart := time.Now()

			var positionStruct Position
			positionStruct.position = pos
			score := Negamax(positionStruct, -22, 22)

			// run solver function -- returns should be:
			// 		- actual score
			// 		- number of nodes explored
			// will prob need to split up position to be integers and fed in to the solver 1 at a time
			// end timer

			duration := time.Since(timerStart)
			duration = duration / time.Microsecond

			if score != expectedScore {
				panic("Actual Score does not equal expected score!")
			}

			fmt.Println("Score: ", score)
			fmt.Println("Time: ", duration)
		}
		file.Close()
	}

	// fmt.Println("--------Averages------")

	// for i := 0; i < len(testData); i++ {
	// 	data := testData[i]
	// 	fmt.Println("Test ", i)
	// 	fmt.Println("Average Time: ", data[0])
	// 	fmt.Println("Average Nodes Explored: ", data[1])
	// 	fmt.Println("Average Score: ", data[2])
	// 	fmt.Println("Time per Node Explored: ", int(data[0]/data[1]))
	// }
	return testData
}
