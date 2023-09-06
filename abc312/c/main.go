package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
)

func main() {
	r := bufio.NewReader(os.Stdin)
	w := bufio.NewWriter(os.Stdout)
	defer w.Flush()

	var n, m int
	fmt.Fscan(r, &n, &m)

	as := make([]int, n)
	for i := range as {
		fmt.Fscan(r, &as[i])
	}
	ms := make([]int, m)
	for i := range ms {
		fmt.Fscan(r, &ms[i])
	}

	left := 0
	right := int(math.Pow(10, 9))
	for left+1 < right {
		checkNum := (left + right) / 2
		seller := 0
		for i := range as {
			if as[i] <= checkNum {
				seller++
			}
		}
		buyer := 0
		for i := range ms {
			if ms[i] >= checkNum {
				buyer++
			}
		}
		if seller >= buyer {
			right = checkNum
		} else {
			left = checkNum
		}
	}
	fmt.Println(right)
}
