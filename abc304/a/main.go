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

	var n int
	fmt.Fscan(r, &n)

	ss := make([]string, n)
	as := make([]int, n)

	minIndex := -1
	minAge := int(math.Pow(10, 9)) + 1
	for i := 0; i < n; i++ {
		fmt.Fscan(r, &ss[i])
		fmt.Fscan(r, &as[i])
		if as[i] < minAge {
			minAge = as[i]
			minIndex = i
		}
	}

	for i := minIndex; i < len(ss); i++ {
		fmt.Println(ss[i])
	}
	for i := 0; i < minIndex; i++ {
		fmt.Println(ss[i])
	}
}
