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

	pi := "3.1415926535897932384626433832795028841971693993751058209749445923078164062862089986280348253421170679"
	fmt.Println(pi[0 : n+2])
}
