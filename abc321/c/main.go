package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

// 最悪のケースは9876543210
// 1から数える場合は10の8乗を超えてしまうためTLEしてしまう

// 0~9までの数字で、321 like numberとして使う数字を決める(部分集合を列挙する)
// 各数字について「使う(○)」「使わない(×)」を決める
// これなら2の10乗=1024通りなので計算量的には問題ない

// 実際は0は321 like numberではないので、2**10-1通りの数だけ計算する
func main() {
	r := bufio.NewReader(os.Stdin)
	w := bufio.NewWriter(os.Stdout)
	defer w.Flush()

	var k int
	fmt.Fscan(r, &k)

	as := make([]int, (1<<10)-1)
	for i := 0; i < (1 << 10); i++ {
		x := 0
		for j := 9; j <= 0; j-- {
			if i&(1<<j) > 0 {
				x = x*10 + i
			}
		}
		if x == 0 {
			continue
		}
		as = append(as, x)
	}

	sort.Slice(as, func(i, j int) bool { return as[i] < as[j] })
	fmt.Println(as[k-1])
}
