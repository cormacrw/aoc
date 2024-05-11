package main

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

const MAX_RED = 12
const MAX_GREEN = 13
const MAX_BLUE = 14

func main() {
	sum := 0

	filePath, err := filepath.Abs("2/input.txt")
	file, err := os.Open(filePath)

	if err != nil {
		fmt.Println(err)
		fmt.Println("file open err")
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	// optionally, resize scanner's capacity for lines over 64K, see next example
	for scanner.Scan() {
		str := scanner.Text()

		sections := strings.Split(str, ": ")

		fmt.Println(sections[0])
		gameId, err := strconv.Atoi(strings.Split(sections[0], " ")[1])

		if (err != nil) {
			break
		}

		fmt.Println("gameId", gameId)

		rounds := strings.Split(sections[1], "; ")

		valid := true

		for i := 0; i < len(rounds); i++ {
			fmt.Println("rounds", rounds[i])
			picks := strings.Split(rounds[i], ", ")

			for j := 0; j < len(picks); j++ {
				p := strings.Split(picks[j], " ")

				val, _ := strconv.Atoi(p[0])

				if p[1][:1] == "r" && val > MAX_RED {
					valid = false
				}

				if p[1][:1] == "g" && val > MAX_GREEN {
					valid = false
				}

				if p[1][:1] == "b" && val > MAX_BLUE {
					valid = false
				}
			}
		}

		if valid {
			sum = sum + gameId
		}
  }

	fmt.Println("ANSWER: ", sum)
}
