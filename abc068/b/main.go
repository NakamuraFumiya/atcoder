package main

import (
	"fmt"
)

func main() {
	var N int
	fmt.Scan(&N)

	s := make([]int, N)
	for i := 0; i < N; i++ {
		s[i] = i + 1
	}

	// 2で割れた最大の数
	maxCount := 0
	// 出力用の変数
	res := 0
	for i := range s {
		count := 0
		ss := s[i]
		for {
			if ss%2 == 0 {
				ss /= 2
				count++
			} else {
				if count >= maxCount {
					maxCount = count
					res = s[i]
				}
				break
			}
		}
	}
	fmt.Println(res)
}
