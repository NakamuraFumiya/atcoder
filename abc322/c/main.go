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

	var n, m int
	fmt.Fscan(r, &n, &m)

	ns := make([]int, m)
	for i := 0; i < m; i++ {
		fmt.Fscan(r, &ns[i])
	}

	p := 0
	for i := 1; i <= n; i++ {
		if i == ns[p] {
			fmt.Println(0)
			p++
		} else {
			fmt.Println(ns[p] - i)
		}
	}
}
