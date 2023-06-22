package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
)

type Person struct {
	x        int
	y        int
	hasVirus bool
}

func main() {
	r := bufio.NewReader(os.Stdin)
	w := bufio.NewWriter(os.Stdout)
	defer w.Flush()

	var n, d int
	fmt.Fscan(r, &n, &d)

	people := []Person{}
	for i := 0; i < n; i++ {
		var x, y int
		fmt.Fscan(r, &x, &y)
		people = append(people, Person{x, y, false})
	}

	people[0].hasVirus = true

	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			distance := math.Sqrt(math.Pow((float64(people[j].x-people[i].x)), 2) + math.Pow((float64(people[j].y-people[i].y)), 2))
			if people[i].hasVirus || people[j].hasVirus {
				// fmt.Println(i, j, people[i], people[j], d, distance, int(distance))
				if distance <= float64(d) {
					// fmt.Println(i, j, people[i], people[j], d, distance, int(distance))
					people[i].hasVirus = true
					people[j].hasVirus = true
				}
			}
		}
	}
	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			distance := math.Sqrt(math.Pow((float64(people[j].x-people[i].x)), 2) + math.Pow((float64(people[j].y-people[i].y)), 2))
			if people[i].hasVirus || people[j].hasVirus {
				if distance <= float64(d) {
					people[i].hasVirus = true
					people[j].hasVirus = true
				}
			}
		}
	}
	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			distance := math.Sqrt(math.Pow((float64(people[j].x-people[i].x)), 2) + math.Pow((float64(people[j].y-people[i].y)), 2))
			if people[i].hasVirus || people[j].hasVirus {
				if distance <= float64(d) {
					people[i].hasVirus = true
					people[j].hasVirus = true
				}
			}
		}
	}

	for _, v := range people {
		if v.hasVirus {
			fmt.Println("Yes")
		} else {
			fmt.Println("No")
		}
	}
}
