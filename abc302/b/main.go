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

	var H, W int
	fmt.Fscan(r, &H)
	fmt.Fscan(r, &W)

	strSlice := make([]string, H)
	for i := 0; i < H; i++ {
		fmt.Fscan(r, &strSlice[i])
	}

	T := "snuke"
	// i, jがどのくらい変化するのかのテーブルを作る
	di := []int{-1, -1, -1, 0, 0, 1, 1, 1}
	dj := []int{-1, 0, 1, -1, 1, -1, 0, 1}
	for si := 0; si < H; si++ {
		for sj := 0; sj < W; sj++ {
			// 8方向で探索
			for k := 0; k < 8; k++ {
				i, j := si, sj
				for l := 0; l < 5; l++ {
					if i < 0 || j < 0 || i >= H || j >= W {
						break
					}
					if string(strSlice[i][j]) != string(T[l]) {
						break
					}
					if l == 4 {
						i, j = si, sj
						for m := 0; m < 5; m++ {
							fmt.Printf("%d %d\n", i+1, j+1)
							i += di[k]
							j += dj[k]
						}
					}
					i += di[k]
					j += dj[k]
				}
			}
		}
	}
}
