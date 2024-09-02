package main

import "fmt"

func main() {
	/*	for num := 1; num < 20; num++ {
		fmt.Println("Loop: ", num)
	}*/

	/* num := 1
	   for {
	    num += 1
	    if num > 10 {
	     break
	    }
	    fmt.Println("Loop: ", num)
	   }
	*/

	num := 1
	for num <= 10 {
		fmt.Println("Loop: ", num)
		num++
	}
}
