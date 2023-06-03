package main

import (
	"bufio"
	"fmt"
	"os"
)

type Coordinate struct {
	x int
	y int
}

func main() {
	r := bufio.NewReader(os.Stdin)
	w := bufio.NewWriter(os.Stdout)
	defer w.Flush()

	var n, m, h, k int
	fmt.Fscan(r, &n, &m, &h, &k)

	var s string
	fmt.Fscan(r, &s)

	items := make(map[Coordinate]int, 0)
	for i := 0; i < m; i++ {
		var x, y int
		fmt.Fscan(r, &x, &y)
		item := Coordinate{x, y}
		items[item]++
	}

	ans := true
	current := Coordinate{0, 0}
	for i := range s {
		h--

		w := string(s[i])
		switch w {
		case "R":
			current.x++
		case "L":
			current.x--
		case "U":
			current.y++
		case "D":
			current.y--
		}

		if h < 0 {
			ans = false
			break
		}

		// アイテムの消費
		if items[current] > 0 {
			h++
			items[current]--
		}
	}

	if ans {
		fmt.Println("Yes")
	} else {
		fmt.Println("No")
	}
}
