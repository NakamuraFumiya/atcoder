package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	r := bufio.NewReader(os.Stdin)
	w := bufio.NewWriter(os.Stdout)
	defer w.Flush()

	var n string
	fmt.Fscan(r, &n)
	ns := strings.Split(n, "")
	ans := true

	f, _ := strconv.Atoi(string(ns[0]))
	if len(ns) > 1 {
		for i := range ns {
			if i == 0 {
				continue
			}
			a, _ := strconv.Atoi(ns[i])
			if f <= a {
				ans = false
				break
			}
			f = a
		}
	}

	if ans {
		fmt.Println("Yes")
	} else {
		fmt.Println("No")
	}
}
