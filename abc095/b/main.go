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

	var n, x int
	fmt.Fscan(r, &n)
	fmt.Fscan(r, &x)

	tmp := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(r, &tmp[i])
	}

	count := 0
	for i := range tmp {
		x -= tmp[i]
		count++
	}

	min := tmp[0]
	for i := range tmp {
		if tmp[i] < min {
			min = tmp[i]
		}
	}

	count += x / min
	fmt.Println(count)
}
