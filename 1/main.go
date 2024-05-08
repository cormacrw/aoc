package main

import (
	"bufio"
	"embed"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var content embed.FS

func main() {
	sum := 0
	file, err := os.Open("/Users/cormacwilliamson/cormacdev/go/aoc-1/input.txt")

	if err != nil {
		fmt.Println(err)
		fmt.Println("file open err")
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	// optionally, resize scanner's capacity for lines over 64K, see next example
	for scanner.Scan() {
		fmt.Println("LINE", scanner.Text())
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

		fmt.Println("FOUND VALUE", leftVal, rightVal)

		val, err := strconv.Atoi(leftVal + rightVal)
		if err == nil {
			fmt.Println("VALUE", val)
			sum = sum + val
		}
	}

	fmt.Println(sum)

	if err := scanner.Err(); err != nil {
		fmt.Println("scanner err")
	}


}

