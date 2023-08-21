package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"sort"
)

func main() {
	r := bufio.NewReader(os.Stdin)
	w := bufio.NewWriter(os.Stdout)
	defer w.Flush()

	var n int64
	fmt.Fscan(r, &n)

	var as = make([]int64, n)
	for i := int64(0); i < n; i++ {
		fmt.Fscan(r, &as[i])
	}

	var sum int64
	for i := range as {
		sum += as[i]
	}

	sort.Slice(as, func(i, j int) bool { return as[i] < as[j] })

	p := sum / n
	l := sum % n

	// 操作後のスライスを作成する
	var bs = make([]int64, n)
	for i := range bs {
		bs[i] = p
	}
	for i := int64(0); i < l; i++ {
		bs[i]++
	}

	sort.Slice(bs, func(i, j int) bool { return bs[i] < bs[j] })

	var diff uint64
	for i := range as {
		diff += uint64(math.Abs(float64(as[i] - bs[i])))
	}

	fmt.Println(diff / 2)
}
