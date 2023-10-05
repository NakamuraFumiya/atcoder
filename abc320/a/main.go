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

	var a, b float64
	fmt.Fscan(r, &a, &b)

	fmt.Println(int(math.Pow(a, b) + math.Pow(b, a)))
}
