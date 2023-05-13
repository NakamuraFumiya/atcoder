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

	var N int
	fmt.Fscan(r, &N)

	s := make([]int, N)
	for i := range s {
		fmt.Fscan(r, &s[i])
	}

	res := 0
	for true {
		existOdd := false
		// 奇数が存在するか確認
		for i := range s {
			if s[i]%2 != 0 {
				existOdd = true
				break
			}
		}

		// 奇数が存在するならループ処理を抜ける
		if existOdd {
			break
		}

		// 全て偶数なら操作処理を行う
		for i := range s {
			s[i] /= 2
		}
		res++
	}

	fmt.Println(res)
}
