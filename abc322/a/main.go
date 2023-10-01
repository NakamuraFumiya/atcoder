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

	var n int
	fmt.Fscan(r, &n)

	var s string
	fmt.Fscan(r, &s)

	ss := strings.Split(s, "")

	count := 0
	for i := range ss {
		if count == 0 {
			if ss[i] == "A" {
				count++
			}
		} else if count == 1 {
			if ss[i] == "B" {
				count++
			} else {
				count = 0
			}
		} else if count == 2 {
			if ss[i] == "C" {
				fmt.Println((i + 1) - 2)
				count++
				break
			} else {
				count = 0
			}
		}

		if count == 0 && ss[i] == "A" {
			count++
		}
	}

	if count < 3 {
		fmt.Println(-1)
	}
}
