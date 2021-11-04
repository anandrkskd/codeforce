package main

import (
	"fmt"
	"sort"
)

func main() {
	var a, b, c int
	fmt.Scanf("%d\n", &a)
	fmt.Scanf("%d\n", &b)
	fmt.Scanf("%d\n", &c)
	var d []int = []int{a, b, c}

	sort.Slice(d, func(i, j int) bool {
		return d[i] > d[j]
	})

	if a == 1 && b == 1 && c == 1 {
		fmt.Println(3)
	} else if a == 1 || b == 1 || c == 1 {
		fmt.Println(d[0] * (d[1] + d[2]))
	} else {
		fmt.Println(a * b * c)
	}
}
