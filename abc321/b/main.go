package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func main() {
	r := bufio.NewReader(os.Stdin)
	w := bufio.NewWriter(os.Stdout)
	defer w.Flush()

	var n, x int
	fmt.Fscan(r, &n, &x)

	as := make([]int, n-1)
	for i := range as {
		fmt.Fscan(r, &as[i])
	}

	ok := false
	for i := 0; i < 101; i++ {
		bs := make([]int, n-1)
		for j := range bs {
			bs[j] = as[j]
		}
		bs = append(bs, i)
		sort.Slice(bs, func(i, j int) bool { return bs[i] < bs[j] })

		sum := 0
		for k := 1; k < len(bs)-1; k++ {
			sum += bs[k]
		}

		if sum >= x {
			fmt.Println(i)
			ok = true
			break
		}
	}

	if !ok {
		fmt.Println(-1)
	}
}
