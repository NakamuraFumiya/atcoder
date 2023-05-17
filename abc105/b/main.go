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

	var n int
	fmt.Fscan(r, &n)

	ans := false
	for i := 0; i <= (n / 4); i++ {
		for j := 0; j <= (n / 7); j++ {
			if 4*i+7*j == n {
				ans = true
			}
		}
	}

	if ans {
		fmt.Println("Yes")
	} else {
		fmt.Println("No")
	}
}
