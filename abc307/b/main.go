package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	r := bufio.NewReader(os.Stdin)
	w := bufio.NewWriter(os.Stdout)
	defer w.Flush()

	var n int
	fmt.Fscan(r, &n)

	chars := make([]string, n)
	for i := 0; i < n; i++ {
		var char string
		fmt.Fscan(r, &char)
		chars[i] = char
	}

	ans := false
Loop:
	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			// 1つ目の組み合わせ
			checkStr := chars[i] + chars[j]

			runes := []rune(checkStr)
			for k, l := 0, len(runes)-1; k < l; k, l = k+1, l-1 {
				runes[k], runes[l] = runes[l], runes[k]
			}

			reservedStr := string(runes)
			if checkStr == reservedStr {
				ans = true
				break Loop
			}

			// 2つ目の組み合わせ
			checkStr = chars[j] + chars[i]

			runes = []rune(checkStr)
			for k, l := 0, len(runes)-1; k < l; k, l = k+1, l-1 {
				runes[k], runes[l] = runes[l], runes[k]
			}

			reservedStr = string(runes)
			if checkStr == reservedStr {
				ans = true
				break Loop
			}
		}
	}

	if ans {
		fmt.Println("Yes")
	} else {
		fmt.Println("No")
	}
}
