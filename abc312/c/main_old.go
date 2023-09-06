// go:build ignore
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

	var n, m uint64
	fmt.Fscan(r, &n, &m)

	var as = make([]uint64, n)
	for i := uint64(0); i < n; i++ {
		fmt.Fscan(r, &as[i])
	}

	var bs = make([]uint64, m)
	for i := uint64(0); i < m; i++ {
		fmt.Fscan(r, &bs[i])
	}

	wa, ac := uint64(0), uint64(1001001001)
	for wa+1 < ac {
		wj := uint64((wa + ac) / 2)
		na, nb := 0, 0
		for i := uint64(0); i < n; i++ {
			if as[i] <= wj {
				na++
			}
		}
		for j := uint64(0); j < m; j++ {
			if bs[j] >= wj {
				nb++
			}
		}
		if na >= nb {
			ac = wj
		} else {
			wa = wj
		}
	}
	fmt.Println(ac)
}
