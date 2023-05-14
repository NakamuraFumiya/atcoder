package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
)

func main() {
	r := bufio.NewReader(os.Stdin)
	w := bufio.NewWriter((os.Stdout))
	defer w.Flush()

	var n, t, a int
	fmt.Fscan(r, &n)
	fmt.Fscan(r, &t)
	fmt.Fscan(r, &a)

	// s := []int{}
	s := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(r, &s[i])
	}

	calc := make([]float64, n)
	for i, v := range s {
		calc[i] = float64(t) - float64(v)*0.006
	}

	for i := range calc {
		calc[i] = math.Abs(calc[i] - float64(a))
	}

	index := 0
	min := calc[0]
	for i, v := range calc {
		if v < min {
			index = i
			min = v
		}
	}

	fmt.Println(index + 1)
}
