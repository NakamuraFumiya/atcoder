package main

import (
	"fmt"
)

func ain() {
	var s, t string
	fmt.Scan(&s)
	fmt.Scan(&t)

	ms, mt := map[string]int{}, map[string]int{}
	atS, atT := 0, 0

	for i := 0; i < len(s); i++ {
		if string(s[i]) == "@" {
			atS++
		} else {
			ms[string(s[i])]++
		}

		if string(t[i]) == "@" {
			atT++
		} else {
			mt[string(t[i])]++
		}
	}

	atcoder := []string{"a", "t", "c", "o", "d", "e", "r"}
	ans := true

	for _, v := range atcoder {
		if ms[v] > mt[v] {
			tmp := ms[v] - mt[v]
			if tmp > atT {
				break
			}
			mt[v] += tmp
			atT -= tmp
		} else if ms[v] < mt[v] {
			tmp := mt[v] - ms[v]
			if tmp > atS {
				break
			}
			ms[v] += tmp
			atS -= tmp
		}
	}

	for j := range ms {
		ans = ans && ms[j] == mt[j]
	}
	for j := range mt {
		ans = ans && ms[j] == mt[j]
	}

	if ans {
		fmt.Println("Yes")
	} else {
		fmt.Println("No")
	}

}
