package main

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

func main() {
	sum := 0
	filePath, err := filepath.Abs("1/input.txt")
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
		chars := strings.Split(str, "")
		i := 0
		j := len(chars) - 1
		leftVal := ""
		rightVal := ""

		for i <= j {
			if (len(leftVal) > 0 && len(rightVal) > 0) {
				break
			}

			if (len(leftVal) == 0) {
				_, err1 := strconv.Atoi(chars[i])

				if err1 == nil {
					leftVal = chars[i]
				} else {
					i++
				}
			}

			if (len(rightVal) == 0) {
				_, err2 := strconv.Atoi(chars[j])

				if err2 == nil {
					rightVal = chars[j]
				} else {
					j--
				}
			}

		}

		val, err := strconv.Atoi(leftVal + rightVal)
		if err == nil {
			sum = sum + val
		}
	}

	fmt.Println("ANSWER: ", sum)

	if err := scanner.Err(); err != nil {
		fmt.Println("scanner err")
	}

}

