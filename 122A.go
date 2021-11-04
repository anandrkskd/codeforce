package main

import "fmt"

func main() {
	var a int
	var result bool = true
	fmt.Scanf("%d", &a)
	stringA := fmt.Sprintf("%d", a)

	for _, i := range stringA {
		if i == '4' || i == '7' {
			continue
		} else {
			result = false
			break
		}
	}
	if int(a)%4 == 0 || int(a)%7 == 0 || luckychecker(a) && result == false {
		result = true
	}

	check(result)
}
func luckychecker(a int) bool {
	result := true
	for i := 8; i <= a/2; i++ {
		result = true
		stringA := fmt.Sprintf("%d", i)

		for _, i := range stringA {
			if i == '4' || i == '7' {
				continue
			} else {
				result = false
				break
			}
		}

		if result && a%i == 0 {
			return true
		}
	}
	return false
}

func check(result bool) {
	if result {
		fmt.Println("YES")
	} else {
		fmt.Println("NO")
	}
}
