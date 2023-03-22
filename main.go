package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	input := bufio.NewReader(os.Stdin)
	s, err := input.ReadString('\n')
	if err != nil {
		panic(err)
	}
	var result string
	mapCount := make(map[string]int64)
	for _, c := range s {
		mapCount[string(c)]++
	}
	var max int64
	var maxStr []string
	for _, i := range mapCount {
		if i > max {
			max = i
		}
	}
	for c, i := range mapCount {
		if i == max {
			maxStr = append(maxStr, c)
		}
	}
	for i := 0; len(result) <= 2 && i < len(s); i++ {
		c := s[i : i+1]
		for _, m := range maxStr {
			if c == m {
				result += c
				break
			}
		}
	}
	fmt.Println(result)
}
