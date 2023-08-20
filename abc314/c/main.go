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

	var n, m uint64
	fmt.Fscan(r, &n, &m)

	var s string
	fmt.Fscan(r, &s)

	cs := make([]int, n)
	for i := uint64(0); i < n; i++ {
		fmt.Fscan(r, &cs[i])
	}

	// keyがcolor, valueがS(文字列)のindexのMapを作成
	colorMap := make(map[int][]uint64, m)
	for i, v := range cs {
		colorMap[v] = append(colorMap[v], uint64(i))
	}

	// 文字列の操作ができるようにruneでキャストする
	tRune := []rune(s)
	ans := []rune(s)

	// map[1:[0_3_6]_2:[1_4_5_7]_3:[2]]

	for _, v := range colorMap {
		if len(v) < 2 {
			continue
		}
		// 文字列の操作を行う
		for i := 0; i < len(v); i++ {
			if i == len(v)-1 {
				ans[v[0]] = tRune[v[len(v)-1]]
			} else {
				ans[v[i+1]] = tRune[v[i]]
			}
		}
	}

	fmt.Println(string(ans))

}
