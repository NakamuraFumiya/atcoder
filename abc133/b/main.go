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

	var n, d int
	fmt.Fscan(r, &n)
	fmt.Fscan(r, &d)

	ans := 0
	var tmp = make([][]int, n)
	for i := 0; i < n; i++ {
		for j := 0; j < d; j++ {
			var x int
			fmt.Fscan(r, &x)
			tmp[i] = append(tmp[i], x)
		}
	}

	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			var x float64
			for k := 0; k < d; k++ {
				x += math.Pow(float64(tmp[i][k]-tmp[j][k]), 2)
			}
			x = math.Sqrt(x)
			if math.Floor(x) == x {
				ans++
			}
		}
	}

	fmt.Println(ans)
}
