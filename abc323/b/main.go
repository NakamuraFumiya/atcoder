package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strings"
)

func main() {
	r := bufio.NewReader(os.Stdin)
	w := bufio.NewWriter(os.Stdout)
	defer w.Flush()

	var n int
	fmt.Fscan(r, &n)

	// keyが勝った数、valueがプレイヤーの番号,
	ansMap := make(map[int][]int)
	maxCount := -1
	counts := make([]int, 0)
	for i := 0; i < n; i++ {
		var s string
		fmt.Fscan(r, &s)

		ss := strings.Split(s, "")

		count := 0
		for _, v := range ss {
			if v == "o" {
				count++
			}
		}

		ansMap[count] = append(ansMap[count], i+1)
		counts = append(counts, count)

		if count > maxCount {
			maxCount = count
		}
	}

	// mapの配列を昇順にする
	for _, vs := range ansMap {
		sort.Slice(vs, func(i, j int) bool { return vs[i] < vs[j] })
	}

	// 重複削除して、降順にする
	tmp := make(map[int]struct{})
	for _, v := range counts {
		tmp[v] = struct{}{}
	}
	orderCounts := make([]int, 0)
	for i := range tmp {
		orderCounts = append(orderCounts, i)
	}
	sort.Slice(orderCounts, func(i, j int) bool { return orderCounts[i] > orderCounts[j] })

	ansStr := ""
	for _, c := range orderCounts {
		for _, v := range ansMap[c] {
			ansStr += fmt.Sprintf("%d ", v)
		}
	}

	fmt.Printf("%s\n", strings.Trim(ansStr, " "))
}
