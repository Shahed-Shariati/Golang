package main

import (
	"fmt"
	"regexp"
	"strings"
)

func main() {
	firstName := "Shahed"
	lastName := "SHARIATI"
	fmt.Println(strings.ToUpper(firstName))
	fmt.Println(strings.ToLower(lastName))
	fmt.Println(strings.Compare(firstName, lastName))
	fmt.Println(strings.Contains(lastName, "ATI"))
	fmt.Println("Length: ", len(lastName))
	fmt.Println("------------------------------Example 1------------------------------------------")

	// Creating and initializing the strings
	str1 := "Welcome, to the online portal, of Scaler Academy"
	str2 := "My name is Akshay"
	str3 := "I like to play Cricket"

	// Displaying strings
	fmt.Println("String 1: ", str1)
	fmt.Println("String 2: ", str2)
	fmt.Println("String 3: ", str3)

	// Splitting the given strings by using SplitAfter() function
	res1 := strings.SplitAfter(str1, ",")
	res2 := strings.SplitAfter(str2, "")
	res3 := strings.SplitAfter(str3, "!")
	res4 := strings.SplitAfter("", "Scaler,Academy")

	// Displaying the result
	fmt.Println("\nResult 1: ", res1)
	fmt.Println("Result 2: ", res2)
	fmt.Println("Result 3: ", res3)
	fmt.Println("Result 4: ", res4)
	fmt.Println("------------------Split by a Comma or other Delimiter------------------------")
	fruitsString := "jainey.rohan.amisha.disha"
	fmt.Println("String: ", fruitsString)
	fruits := strings.SplitN(fruitsString, ".", 3)
	fmt.Println(fruits)
	fmt.Println("---------------------Split on Regular Expression--------------------------")
	s := regexp.MustCompile("[0-9]").Split("jainey1rohan2amisha3jain", -1)
	fmt.Println(s)

	fmt.Println("--------------------Split by Whitespace and Newline--------------------------")
	fmt.Println("Fields are: %q", strings.Fields(`jainey rohan amisha pear`))
}
