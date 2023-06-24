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

	var n uint64
	fmt.Fscan(r, &n)

	as := make([]uint64, 0)
	for i := uint64(0); i < 7*n; i++ {
		var num uint64
		fmt.Fscan(r, &num)
		as = append(as, num)
	}

	count := 0
	var total uint64
	ans := []uint64{}
	for _, v := range as {
		count++
		total += v
		if count%7 == 0 {
			ans = append(ans, total)
			total = 0
			count = 0
		}
		// if i == len(as)-1 {
		// 	ans = append(ans, total)
		// 	break
		// }
	}

	for i, v := range ans {
		if i == len(ans)-1 {
			fmt.Printf("%d\n", v)
		} else {
			fmt.Printf("%d ", v)
		}
	}

}
