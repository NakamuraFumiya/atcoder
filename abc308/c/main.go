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

	var n uint64
	fmt.Fscan(r, &n)

	as := make([]uint64, n)
	bs := make([]uint64, n)

	for i := uint64(0); i < n; i++ {
		var a, b uint64
		fmt.Fscan(r, &a, &b)
		as[i] = a
		bs[i] = b
	}
	// fmt.Println(as, bs)

	// 確率がキー、人(複数人)がバリューのMapを作る
	// バリューに複数いるときは、人が早い順でスライスに詰める
	// ans := make([]uint64, n)
	tmp := make(map[float64][]uint64)
	for i := uint64(0); i < n; i++ {
		// fmt.Printf("as[i]+bs[i]: %d\n", as[i]+bs[i])
		// fmt.Printf("as[i]/(as[i]+bs[i]): %d\n", as[i]/(as[i]+bs[i]))
		// fmt.Printf("test////as[i]/(as[i]+bs[i]): %v\n", float64(as[i])/float64(as[i]+bs[i]))
		chance := float64(as[i]) / float64(as[i]+bs[i])
		tmp[float64(chance)] = append(tmp[float64(chance)], i+1)
	}

	// tmpのバリューを昇順でsortしておく
	for _, v := range tmp {
		sort.Slice(v, func(i, j int) bool { return v[i] < v[j] })
	}

	// chanceの降順のスライスを作っておく
	var orderChances []float64
	for i := range tmp {
		orderChances = append(orderChances, i)
	}
	sort.Slice(orderChances, func(i, j int) bool { return orderChances[i] > orderChances[j] })

	ans := make([]uint64, 0)
	for _, v := range orderChances {
		ans = append(ans, tmp[v]...)
	}

	for i := 0; i < len(ans)-1; i++ {
		fmt.Printf("%d ", ans[i])
	}
	fmt.Println(ans[len(ans)-1])

}
