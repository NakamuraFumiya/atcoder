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

	var n, m int
	fmt.Fscan(r, &n, &m)

	var s, t string
	fmt.Fscan(r, &s, &t)

	ans := map[string]bool{"prefix": false, "suffix": false}

	// 接頭辞であるかチェック
	if s == t[0:n] {
		ans["prefix"] = true
	}
	// 接尾辞であるかチェック
	if s == t[m-n:] {
		ans["suffix"] = true
	}

	if ans["prefix"] && ans["suffix"] {
		fmt.Println(0)
	} else if ans["prefix"] {
		fmt.Println(1)
	} else if ans["suffix"] {
		fmt.Println(2)
	} else {
		fmt.Println(3)
	}
}
