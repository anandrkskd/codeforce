package main

import (
	"fmt"
	"strings"
)

func main() {
	var s, newstr string
	fmt.Scanf("%s", &s)

	for _, i := range s {
		cmp := strings.ToLower(string(i))
		if cmp == "a" || cmp == "e" || cmp == "i" || cmp == "o" || cmp == "u" || cmp == "y" {
			continue
		} else {
			newstr = strings.Join([]string{newstr, ".", cmp}, "")
		}
	}
	fmt.Println(newstr)
}
