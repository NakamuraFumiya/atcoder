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

	var a, b int
	fmt.Fscan(r, &a)
	fmt.Fscan(r, &b)

	ans := a / b
	if a%b != 0 {
		ans++
	}

	fmt.Println(ans)
}
