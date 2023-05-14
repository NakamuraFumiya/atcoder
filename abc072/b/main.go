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

	var s string
	fmt.Fscan(r, &s)

	res := ""
	for i, v := range s {
		i += 1
		if i%2 != 0 {
			res += string(v)
		}
	}

	fmt.Println(res)
}
