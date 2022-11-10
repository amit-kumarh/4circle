// ideally takes a data set
// plays the positions
// evaluates time to sovle

package main

// package main

import (
	// direct path
	// direct path
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"time"
)

func tester() [][]int {
	// files := []string{"Test_L1_R1.txt", "Test_L1_R2.txt", "Test_L1_R3.txt", "Test_L2_R1.txt", "Test_L2_R2.txt", "Test_L3_R1.txt"}
	files := []string{"testPosition.txt"}
	testData := make([][]int, len(files))

	for i := 0; i < len(files); i++ {
		file, err := os.Open(files[i])

		if err != nil {
			log.Fatal(err)
		}

		// defer file.Close()
		scanner := bufio.NewScanner(file)
		scanner.Split(bufio.ScanWords)

		for scanner.Scan() {
			position := scanner.Text()
			var positionStruct Position

			scanner.Scan()

			expectedScore, error := strconv.Atoi(scanner.Text())

			if error != nil {
				log.Fatal(error)
			}

			posArray := strings.Split(position, "")
			for i := 0; i < len(posArray); i++ {
				posInt, _ := strconv.Atoi(posArray[i])
				Play(&positionStruct, posInt)
			}

			fmt.Println("Position Bitstring: ", positionStruct.position)
			fmt.Println("Mask Bitstring: ", positionStruct.mask)
			fmt.Println("Num of Moves: ", positionStruct.moves)
			// start timer
			timerStart := time.Now()

			score := Negamax(positionStruct, -21, 21)

			duration := time.Since(timerStart)
			duration = duration / time.Microsecond

			fmt.Println("score: ", score)
			fmt.Println("expected score: ", expectedScore)
			if score != expectedScore {
				panic("Actual Score does not equal expected score!")
			}
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
