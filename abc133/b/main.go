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

	arrParent := make([][]int, n)
	for i := 0; i < n; i++ {
		arrChild := make([]int, d)
		for j := 0; j < d; j++ {
			fmt.Fscan(r, &arrChild[j])
		}
		arrParent[i] = arrChild
	}

	ans := 0
	for i := 0; i+1 < n; i++ {
		tmp := float64(0)
		for j := 0; j < d; j++ {
			tmp += math.Pow(float64(arrParent[i][j]-arrParent[i+1][j]), 2)
		}
		if math.Floor(math.Sqrt(tmp)) == math.Sqrt(tmp) {
			ans++
		}
	}

	// fmt.Printf("%v", arrParent)
	fmt.Println(ans)
}
