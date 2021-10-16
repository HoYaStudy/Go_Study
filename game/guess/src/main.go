package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"time"
)

var stdin = bufio.NewReader(os.Stdin)

func InputIntValue() (int, error) {
	var n int
	_, err := fmt.Scanln(&n)
	if err != nil {
		stdin.ReadString('\n')
	}
	return n, err
}

func main() {
	rand.Seed(time.Now().UnixNano())

	r := rand.Intn(100)
	cnt := 1

	for {
		fmt.Printf("Input number (0 ~ 99): ")
		n, err := InputIntValue()
		if err != nil {
			fmt.Println("Error: Not a number!")
		} else {
			if n > r {
				fmt.Println("Input number is bigger.")
			} else if n < r {
				fmt.Println("Input number is smaller.")
			} else {
				fmt.Printf("Correct! You tried %d times.\n", cnt)
				break
			}
			cnt++
		}
	}
}
