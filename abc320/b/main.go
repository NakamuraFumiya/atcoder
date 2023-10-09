package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	r := bufio.NewReader(os.Stdin)
	w := bufio.NewWriter(os.Stdout)
	defer w.Flush()

	var s string
	fmt.Fscan(r, &s)

	ss := strings.Split(s, "")

	// 1文字も回文として扱うため1で初期化する
	maxCount := 1
	for i := 0; i < len(ss); i++ {
		for j := i + 1; j < len(ss); j++ {
			subString := s[i : j+1]
			reverseSubString := reverse(subString)
			if subString == reverseSubString {
				if len(subString) > maxCount {
					maxCount = len(subString)
				}
			}
		}
	}

	fmt.Println(maxCount)
}

func reverse(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}

	return string(runes)
}
