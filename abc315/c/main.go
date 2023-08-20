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

	// 一番目に満足度の高いアイスの味と満足度を算出
	var MaxF, MaxS uint64

	// keyが味(F), valueが満足度(S)のmapを作成する
	iceMap := make(map[uint64][]uint64)
	for i := uint64(0); i < n; i++ {
		var f, s uint64
		fmt.Fscan(r, &f, &s)
		iceMap[f] = append(iceMap[f], s)

		if s > MaxS {
			MaxS = s
			MaxF = f
		}
	}

	// 一番満足度の高いアイスと二番目に満足度の高いアイスの味を算出
	var sForMultipleF uint64
	// 違う味で、二番目に満足度の高いアイスの満足度を算出
	var SecondMaxS uint64
	for i, v := range iceMap {
		if i != MaxF {
			for _, l := range v {
				if SecondMaxS < l {
					SecondMaxS = l
				}
			}
		}
	}
	sForMultipleF = MaxS + SecondMaxS

	// 同じ味で一番満足度の高い組み合わせ
	var sForSameF uint64
	for _, v := range iceMap {
		if len(v) >= 2 {
			sort.Slice(v, func(i, j int) bool { return v[i] > v[j] })
			res := v[0] + v[1]/2
			if res > sForSameF {
				sForSameF = res
			}
		}
	}

	// 比較して満足度の高い方を出力
	if sForMultipleF >= sForSameF {
		fmt.Println(sForMultipleF)
	} else {
		fmt.Println(sForSameF)
	}
}
