package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func findStrings(n int, stringsList []string) (bool, []int) {
	seen := make(map[string][]int)

	for i := 0; i < n; i++ {
		lowerStr := strings.ToLower(stringsList[i])
		seen[lowerStr] = append(seen[lowerStr], i+1)
	}

	for _, indices := range seen {
		if len(indices) > 1 {
			return true, indices
		}
	}

	return false, nil
}

func main() {
	var n int
	fmt.Print("input : ")
	fmt.Scan(&n)

	stringsList := make([]string, n)
	scanner := bufio.NewScanner(os.Stdin)
	for i := 0; i < n; i++ {
		fmt.Printf("input string ke-%d: ", i+1)
		scanner.Scan()
		stringsList[i] = scanner.Text()
	}

	found, result := findStrings(n, stringsList)
	if found {
		fmt.Print("output: ")
		for _, index := range result {
			fmt.Printf("%d ", index)
		}
		fmt.Println()
	} else {
		fmt.Println("output: false")
	}
}
