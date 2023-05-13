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

	var N int
	fmt.Fscan(r, &N)

	s := make([]int, N)
	for i := range s {
		fmt.Fscan(r, &s[i])
	}

	sort.Slice(s, func(i, j int) bool { return s[i] < s[j] })
	fmt.Println(s[len(s)-1] - s[0])
}
