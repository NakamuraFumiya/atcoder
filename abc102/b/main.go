package main

import (
	"bufio"
	"fmt"
	"math"
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

	var res int
	for i := 0; i+1 < N; i++ {
		for j := i + 1; j < N; j++ {
			abs := int(math.Abs(float64(s[j] - s[i])))
			if abs > res {
				res = abs
			}
		}
	}
	fmt.Println(res)
}
