package main

import (
	"fmt"
	"sort"
	"strings"
)

func main() {
	var s, t string
	fmt.Scan(&s)
	fmt.Scan(&t)
	tmp1 := strings.Split(s, "")
	tmp2 := strings.Split(t, "")
	// 昇順
	sort.Slice(tmp1, func(i, j int) bool { return tmp1[i] < tmp1[j] })
	// 降順
	sort.Slice(tmp2, func(i, j int) bool { return tmp2[i] > tmp2[j] })
	var ss, tt string
	for i := range tmp1 {
		ss = ss + tmp1[i]
	}
	for i := range tmp2 {
		tt = tt + tmp2[i]
	}
	if ss < tt {
		fmt.Println("Yes")
	} else {
		fmt.Println("No")
	}
}
