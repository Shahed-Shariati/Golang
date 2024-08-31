package main

import (
	"fmt"
	"strings"
)

func main() {
	num := 2
	str := "Shahed"
	if num < 2 {
		fmt.Println("ok")
	} else {
		fmt.Println("no")
	}

	if str == "Shahed" {
		fmt.Println("ok")
	} else {
		fmt.Println("not equals....")
	}

	switch str {
	case "sam":
		fmt.Println("sam")
	case "Shahed":
		fmt.Println(strings.ToUpper(str))
	default:
		fmt.Println("not equals")

	}

}