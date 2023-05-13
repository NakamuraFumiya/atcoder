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

	var N int
	fmt.Fscan(r, &N)

	s := make([]int, N)
	for i := range s {
		fmt.Fscan(r, &s[i])
	}

	count := 0
	flg := true
	for flg {
		for i := range s {
			if s[i]%2 == 0 {
				s[i] = s[i] / 2
			} else {
				flg = false
				break
			}
		}
		if flg {
			count++
		}
	}

	fmt.Println(count)
}
