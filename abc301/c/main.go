package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	r := bufio.NewReader(os.Stdin)
	w := bufio.NewWriter(os.Stdout)
	defer w.Flush()

	var s, t string
	fmt.Fscan(r, &s)
	fmt.Fscan(r, &t)

	sMap, tMap := map[string]int{}, map[string]int{}
	atcoderStr := "atcoder"
	// atcoderStr := []string{"a", "t", "c", "o", "d", "e", "r"}

	for i := range s {
		sMap[string(s[i])]++
	}
	for i := range t {
		tMap[string(t[i])]++
	}

	ans := true
	for i := range sMap {
		if i == "@" {
			continue
		}
		if sMap[i] < tMap[i] {
			tmp := tMap[i] - sMap[i]
			if sMap["@"] >= tmp && strings.Contains(atcoderStr, i) {
				sMap["@"] -= tmp
				sMap[i] += tmp
			} else {
				ans = false
			}
		}

		if sMap[i] > tMap[i] && strings.Contains(atcoderStr, i) {
			tmp := sMap[i] - tMap[i]
			if tMap["@"] >= tmp {
				tMap["@"] -= tmp
				tMap[i] += tmp
			} else {
				ans = false
			}
		}
	}

	for i := range tMap {
		if i == "@" {
			continue
		}
		if sMap[i] < tMap[i] && strings.Contains(atcoderStr, i) {
			tmp := tMap[i] - sMap[i]
			if sMap["@"] >= tmp {
				sMap["@"] -= tmp
				sMap[i] += tmp
			} else {
				ans = false
			}
		}

		if sMap[i] > tMap[i] && strings.Contains(atcoderStr, i) {
			tmp := sMap[i] - tMap[i]
			if tMap["@"] >= tmp {
				tMap["@"] -= tmp
				tMap[i] += tmp
			} else {
				ans = false
			}
		}
	}

	for i := range sMap {
		ans = ans && sMap[i] == tMap[i]
	}
	for i := range tMap {
		ans = ans && tMap[i] == sMap[i]
	}

	if ans {
		fmt.Println("Yes")
	} else {
		fmt.Println("No")
	}
}
