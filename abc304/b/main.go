package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
)

func main() {
	r := bufio.NewReader(os.Stdin)
	w := bufio.NewWriter(os.Stdout)
	defer w.Flush()

	var n uint64
	fmt.Fscan(r, &n)

	if n <= (1000 - 1) {
		fmt.Println(n)
	} else if n <= (10000 - 1) {
		fmt.Println(uint64(math.Floor(float64(n/10))) * 10)
	} else if n <= (100000 - 1) {
		fmt.Println(uint64(math.Floor(float64(n/100))) * 100)
	} else if n <= (1000000 - 1) {
		fmt.Println(uint64(math.Floor(float64(n/1000))) * 1000)
	} else if n <= (10000000 - 1) {
		fmt.Println(uint64(math.Floor(float64(n/10000))) * 10000)
	} else if n <= (100000000 - 1) {
		fmt.Println(uint64(math.Floor(float64(n/100000))) * 100000)
	} else if n <= (1000000000 - 1) {
		fmt.Println(uint64(math.Floor(float64(n/1000000))) * 1000000)
	}

}
