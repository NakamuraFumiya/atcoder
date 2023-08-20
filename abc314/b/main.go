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

	cs := make([]int, n)
	// keyが人の添字、valueが実際に賭けた出目
	betNumbers := make(map[int][]int)

	for i := 0; i < n; i++ {
		fmt.Fscan(r, &cs[i])
		as := make([]int, cs[i])
		for i := range as {
			fmt.Fscan(r, &as[i])
		}
		betNumbers[i+1] = as
	}

	// 実際の出目
	var x int
	fmt.Fscan(r, &x)

	// keyがルーレットに当たった人の賭けた出目数、valueが人の添字
	var ans map[int][]int

	// 賭けた出目数の最大値で初期化する
	min := 0
	for _, v := range cs {
		if v > min {
			min = v
		}
	}
	for i, v := range betNumbers {
		for j := range v {
			if v[j] == x {
				// mapの初期化
				if ans == nil {
					ans = make(map[int][]int)
				}
				ans[len(v)] = append(ans[len(v)], i)
				// ルーレットに当たった人の賭けた最小の出目数を更新
				if len(v) < min {
					min = len(v)
				}
			}
		}
	}

	// 答えを昇順に並び替える
	sort.Slice(ans[min], func(i, j int) bool { return ans[min][i] < ans[min][j] })

	if ans == nil {
		fmt.Println(0)
		fmt.Println("")
	} else {
		// 人数
		fmt.Println(len(ans[min]))
		// 人を出力
		for i := 0; i < len(ans[min])-1; i++ {
			fmt.Printf("%d ", ans[min][i])
		}
		fmt.Println(ans[min][len(ans[min])-1])
	}
}
