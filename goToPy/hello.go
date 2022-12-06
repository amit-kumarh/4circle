package main

import (
	"fmt"
	"os/exec"
	"strconv"
)

func main() {
	for col := 1; col <= 7; col++ {
		pythonLine := "import pythonfile; print(pythonfile.to_serial(" + strconv.Itoa(col) + "))"
		cmd := exec.Command("python3", "-c", pythonLine)
		// fmt.Println(cmd.Args)
		out, err := cmd.CombinedOutput()
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println(string(out))
	}

}
