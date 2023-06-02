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

	var s1, s2 string
	fmt.Fscan(r, &s1)
	fmt.Fscan(r, &s2)

	ans := true
	for i := 0; i < n; i++ {
		if string(s1[i]) == string(s2[i]) {
			continue
		} else if (string(s1[i]) == "1" && string(s2[i]) == "l") || (string(s1[i]) == "l" && string(s2[i]) == "1") {
			continue
		} else if (string(s1[i]) == "0" && string(s2[i]) == "o") || (string(s1[i]) == "o" && string(s2[i]) == "0") {
			continue
		}
		ans = false
		break
	}

	if ans {
		fmt.Println("Yes")
	} else {
		fmt.Println("No")
	}
}
