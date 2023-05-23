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

	var n int
	fmt.Fscan(r, &n)

	tmp := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(r, &tmp[i])
	}
	sort.Slice(tmp, func(i, j int) bool { return tmp[i] < tmp[j] })

	ans := 0
	for i := 0; i < len(tmp); i++ {
		for j := i + 1; j < len(tmp); j++ {
			for k := j + 1; k < len(tmp); k++ {
				if tmp[k] < tmp[i]+tmp[j] && tmp[i] != tmp[j] && tmp[i] != tmp[k] && tmp[j] != tmp[k] {
					ans++
				}
			}
		}
	}

	fmt.Println(ans)
}
