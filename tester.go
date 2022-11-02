// ideally takes a data set
// plays the positions
// evaluates time to sovle

package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func readPos(scanner bufio.Scanner) string {
	/*
	 * Function to read intended position from benchmark data file
	 */

	// first scan is position
	scanner.Scan()
	position := scanner.Text()

	scanner.Scan()
	scanner.Text()
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return position

}

func main() {
	file, err := os.Open("Test_L1_R1.txt")

	if err != nil {
		log.Fatal(err)
	}

	// defer file.Close()
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanWords)

	fmt.Println("Position 1: ", readPos(*scanner))
	// scanner.Scan()
	fmt.Println("Position 2: ", readPos(*scanner))
	// scanner.Scan()
}
