// ideally takes a data set
// plays the positions
// evaluates time to sovle

package main

import (
	"bufio"
	"fmt"
	"log"
	"math/rand"
	"os"
	"strconv"
	"time"
)

func tester() [][]int {
	files := []string{"Test_L1_R1.txt", "Test_L1_R2.txt", "Test_L1_R3.txt", "Test_L2_R1.txt", "Test_L2_R2.txt", "Test_L3_R1.txt"}
	testData := make([][]int, len(files))

	for i := 0; i < len(files); i++ {
		file, err := os.Open(files[i])

		if err != nil {
			log.Fatal(err)
		}

		// defer file.Close()
		scanner := bufio.NewScanner(file)
		scanner.Split(bufio.ScanWords)

		averageScore := 0
		averageTime := 0
		averageNodes := 0

		for scanner.Scan() {
			position := scanner.Text()

			scanner.Scan()

			expectedScore, error := strconv.Atoi(scanner.Text())

			if error != nil {
				log.Fatal(error)
			}

			fmt.Println("Position: ", position)
			// fmt.Println("Expected Score: ", expectedScore)

			// start timer
			timerStart := time.Now()

			// run solver function -- returns should be:
			// 		- actual score
			// 		- number of nodes explored
			// will prob need to split up position to be integers and fed in to the solver 1 at a time
			// end timer

			// TO BE CHANGED:
			actualScore := expectedScore

			rand.Seed(time.Now().UnixNano())
			nodesExplored := rand.Intn(10000000)

			//

			duration := time.Since(timerStart)
			duration = duration / time.Microsecond

			if actualScore != expectedScore {
				panic("Actual Score does not equal expected score!")
			}

			// calc averages
			averageTime += int(duration)
			averageNodes += nodesExplored
			averageScore += actualScore

		}

		averageTime = averageTime / 1000
		averageNodes = averageNodes / 1000
		averageScore = averageScore / 1000

		testData[i] = []int{averageTime, averageNodes, averageScore}
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
