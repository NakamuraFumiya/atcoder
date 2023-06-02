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

	var n, m, h, k int
	fmt.Fscan(r, &n)
	fmt.Fscan(r, &m)
	fmt.Fscan(r, &h)
	fmt.Fscan(r, &k)

	var s string
	fmt.Fscan(r, &s)

	xItem := make([]int, m)
	yItem := make([]int, m)
	for i := 0; i < m; i++ {
		var x, y int
		fmt.Fscan(r, &x)
		fmt.Fscan(r, &y)
		xItem = append(xItem, x)
		yItem = append(yItem, y)
	}

	ans := true
	x, y := 0, 0
	for i := range s {
		h--

		w := string(s[i])
		switch w {
		case "R":
			x++
		case "L":
			x--
		case "U":
			y++
		case "D":
			y--
		}

		for j := 0; j < m; j++ {
			if (x == xItem[j] && y == yItem[j]) && h < k {
				h++
				// アイテムの消費
				xItem = append(xItem[:j], xItem[j+1:]...)
				yItem = append(yItem[:j], yItem[j+1:]...)
			}
		}

		if h < 0 {
			ans = false
			break
		}

		// fmt.Printf("h: %d, w: %s\n", h, w)

	}

	if ans {
		fmt.Println("Yes")
	} else {
		fmt.Println("No")
	}
}
