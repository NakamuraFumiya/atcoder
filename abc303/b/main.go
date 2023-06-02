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

	var n, m int
	fmt.Fscan(r, &n)
	fmt.Fscan(r, &m)

	a := make([][]int, m)
	for i := 0; i < m; i++ {
		b := make([]int, 0)
		for j := 0; j < n; j++ {
			var tmp int
			fmt.Fscan(r, &tmp)
			b = append(b, tmp)
		}
		a[i] = b
	}

	base := make([]int, 0)
	base = append(base, a[0]...)
	sort.Slice(base, func(i, j int) bool { return base[i] < base[j] })

	ans := 0
	for i := 0; i < len(base); i++ {
		for j := i + 1; j < len(base); j++ {
			exist := false
			// 与えられたスライスに対して全探索
		Loop:
			for _, v := range a {
				for l := 0; l+1 < n; l++ {
					o := l + 1
					if (v[l] == base[i] && v[o] == base[j]) || (v[l] == base[j] && v[o] == base[i]) {
						// fmt.Printf("i: %d, j: %d, base[i]: %v, base[j]: %v, l: %d, o: %d, v[l]: %v, v[o]: %v\n", i, j, base[i], base[j], l, o, v[l], v[o])
						exist = true
						break Loop
					}
				}
			}
			if !exist {
				ans++
			}
		}
	}

	fmt.Println(ans)
}
