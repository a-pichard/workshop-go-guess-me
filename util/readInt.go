package util

import (
	"bufio"
	"fmt"
	"os"
)

func ReadInt(label string, cond func(int) bool) int {
	stdin := bufio.NewReader(os.Stdin)
	var number int
	for {
		fmt.Print(label)
		_, err := fmt.Fscan(stdin, &number)
		if err == nil {
			if cond(number) {
				break
			} else {
				fmt.Println("Error: bad value")
			}
		} else {
			stdin.ReadString('\n')
			fmt.Println("Error: not a number.")
		}
	}
	return number
}
