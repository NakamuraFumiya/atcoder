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

	var n, m int
	fmt.Fscan(r, &n)
	fmt.Fscan(r, &m)

	cs := make([]string, n)
	for i := 0; i < n; i++ {
		var c string
		fmt.Fscan(r, &c)
		cs[i] = c
	}

	ds := make([]string, m)
	for i := 0; i < m; i++ {
		var d string
		fmt.Fscan(r, &d)
		ds[i] = d
	}

	ps := make([]int, m+1)
	for i := 0; i < m+1; i++ {
		var p int
		fmt.Fscan(r, &p)
		ps[i] = p
	}

	// fmt.Println(n, m, cs, ds, ps)

	priceMap := make(map[string]int)
	for i, v := range ds {
		priceMap[v] = ps[i+1]
	}

	// fmt.Println(priceMap)

	total := 0
	for _, c := range cs {
		if p, ok := priceMap[c]; ok {
			total += p
		} else {
			total += ps[0]
		}
	}
	fmt.Println(total)

}
