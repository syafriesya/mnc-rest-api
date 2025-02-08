package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func isValid(s string) bool {
	if len(s) < 1 || len(s) > 4096 {
		return false
	}

	pairs := map[rune]rune{
		'}': '{',
		']': '[',
		'>': '<',
	}

	stack := []rune{}

	for _, char := range s {
		switch char {
		case '{', '[', '<':
			stack = append(stack, char)
		case '}', ']', '>':
			if len(stack) == 0 || stack[len(stack)-1] != pairs[char] {
				return false
			}
			stack = stack[:len(stack)-1]
		default:
			return false
		}
	}

	return len(stack) == 0
}

func main() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Println("input:")

	input, _ := reader.ReadString('\n')
	input = strings.TrimSpace(input)

	if isValid(input) {
		fmt.Println("Valid: TRUE")
	} else {
		fmt.Println("Valid: FALSE")
	}
}
