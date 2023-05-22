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

	for si := 0; si < H; si++ {
		for sj := 0; sj < W; sj++ {

		}
	}
}
