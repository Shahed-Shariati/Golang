package main

import (
	"fmt"
	"strconv"
)

var sum int //Global variable

func main() {
	var str string
	str = "shahed" // local variable
	num1 := 5
	num2 := 6
	sum = num1 + num2
	fmt.Println(str)
	sumPrint()

	fmt.Println("-----------------------------------------------------")

	strNum := "3"
	number := 5

	sum2 := strNum + strconv.Itoa(number) // convert num to string
	fmt.Println(sum2)

	numStr, _ := strconv.Atoi(strNum) // convert string to number
	sum3 := number + numStr
	fmt.Println(sum3)

}

func sumPrint() {
	fmt.Println(sum)
}



