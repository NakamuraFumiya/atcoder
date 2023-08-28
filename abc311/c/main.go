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

	var n int
	fmt.Fscan(r, &n)

	// 設問と添字を一致させるためn+1で初期化
	as := make([]int, n+1)
	// 0番目の要素は使わないので適当に-1を入れておく
	as[0] = -1
	for i := 1; i < n+1; i++ {
		fmt.Fscan(r, &as[i])
	}

	tmp := make([]int, n+1)
	// 探索を始める頂点をv(vertex)とする
	v := 1
	// 探索回数をkとする
	k := 1

	// tmpに値を詰めていき、同じ頂点に戻ってきた時はtmpには値が入っている
	// つまり有向閉路になる
	for tmp[v] == 0 {
		tmp[v] = k
		k++
		// 探索でみている頂点vを移動させる
		v = as[v]
	}

	// 有向閉路の頂点の数を出力
	l := k - tmp[v]
	fmt.Println(l)

	for i := 0; i < l; i++ {
		fmt.Printf("%d ", v)
		v = as[v]
	}
}
