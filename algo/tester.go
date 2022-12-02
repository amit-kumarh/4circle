// // ideally takes a data set
// // plays the positions
// // evaluates time to sovle

package main

// // package main

import (
	// direct path
	// direct path
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"time"
)

func tester() [][]int {
	masterTimer := time.Now()
	var max time.Duration
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

		for scanner.Scan() {
			position := scanner.Text()
			pos := createPosition()
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
			score := Negamax(pos, sol, -1, 1)

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
		file.Close()
	}
	elapsed := time.Since(masterTimer)
	fmt.Println(elapsed / 1000)
	fmt.Println(max)
	return testData
}
