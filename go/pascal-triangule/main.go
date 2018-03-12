package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("[main] - Problems of args.")
		return
	}
	val, err := strconv.Atoi(os.Args[1])
	if err != nil {
		fmt.Println("[main] - Convert Error: ", err.Error())
	}
	for index := 1; index <= val; index++ {
		for jndex := 1; jndex <= index; jndex++ {
			fmt.Printf("%5d", pascal(index, jndex))
		}
		fmt.Println()
	}
}

func pascal(line int, column int) int {
	if column == 0 {
		return 0
	} else if line == 1 && column == 1 {
		return 1
	} else if column > line {
		return 0
	}
	return pascal(line-1, column-1) + pascal(line-1, column)
}
