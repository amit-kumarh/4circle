package main

import (
	"fmt"
	"os/exec"
	"strconv"
)

func main() {

	for i := 0; i <= 7; i++ {
		var col int
		fmt.Print("Player Played in: ")
		fmt.Scanln(&col)
		pythonLine := "import pythonfile; print(pythonfile.to_serial(" + strconv.Itoa(col) + "))"
		cmd := exec.Command("python3", "-c", pythonLine)
		// fmt.Println(cmd.Args)
		out, err := cmd.CombinedOutput()
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println("Column: ", string(out))
	}

}
