//go:build ignore

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

	tmp := make(map[int]struct{})
	for i := 0; i < n; i++ {
		var x int
		fmt.Fscan(r, &x)
		tmp[x] = struct{}{}
	}

	bar := make([]int, 0)
	for i := range tmp {
		bar = append(bar, i)
	}
	fmt.Printf("tmp: %v\n", tmp)
	fmt.Printf("bar: %v\n", bar)

	ans := 0
	for i := 0; i < len(bar); i++ {
		for j := i + 1; j < len(bar); j++ {
			for k := j + 1; k < len(bar); k++ {
				ans++
			}
		}
	}

	fmt.Println(ans)
}
