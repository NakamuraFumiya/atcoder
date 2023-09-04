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

	ss := make([]string, n)
	for i := range ss {
		fmt.Fscan(r, &ss[i])
	}

	if rec(ss, 0) {
		fmt.Println("Yes")
	} else {
		fmt.Println("No")
	}
}

func diff(s1 string, s2 string) int {
	diff := 0
	for i, v := range s1 {
		if string(v) != string(s2[i]) {
			diff++
		}
	}
	return diff
}

func rec(ss []string, i int) bool {
	// 終了条件: 文字列を使い切ったら true
	if i == len(ss) {
		return true
	}

	// まだ使ってない文字列を選ぶ
	for j := i; j < len(ss); j++ {
		if i == 0 || diff(ss[i-1], ss[j]) == 1 {
			// 要素を入れ替える
			tmp := ss[i]
			ss[i] = ss[j]
			ss[j] = tmp
			if rec(ss, i+1) {
				return true
			}
		}
	}

	return false
}
