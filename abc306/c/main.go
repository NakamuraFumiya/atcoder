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

	a := make([]int, 3*n)
	for i := 0; i < len(a); i++ {
		var num int
		fmt.Fscan(r, &num)
		a[i] = num
	}

	ans := []int{}
	tmp := make(map[int]int)

	for _, v := range a {
		tmp[v]++
		if tmp[v] == 2 {
			ans = append(ans, v)
		}
	}

	for i := 0; i < len(ans)-1; i++ {
		fmt.Printf("%d ", ans[i])
	}
	fmt.Printf("%d\n", ans[len(ans)-1])
}
