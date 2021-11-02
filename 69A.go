package main

import "fmt"

func main() {
	var n, a, b, c int

	fmt.Scanf("%d\n", &n)

	var twoDSlice = make([][]int, n)
	for i := 0; i < n; i++ {
		fmt.Scanf("%d %d %d\n", &a, &b, &c)
		twoDSlice[i] = []int{a, b, c}
	}
	a = 0
	b = 0
	c = 0
	for _, i := range twoDSlice {
		a += i[0]
		b += i[1]
		c += i[2]
	}

	if a == 0 && b == 0 && c == 0 {
		fmt.Println("YES")
	} else {
		fmt.Println("NO")
	}
}
