package main

import (
	"fmt"
	"regexp"
)

func main() {
	var s string
	fmt.Scanf("%s", &s)
	re2, _ := regexp.Compile(".*h.*e.*l.*l.*o")
	if re2.Match([]byte(s)) {
		fmt.Println("YES")
	} else {
		fmt.Println("NO")
	}
}
