package main

import (
	"fmt"
)

func main() {
	println("-----------------Arrays--------------------------")
	nums := []int{102, 45, 846, 423}
	for _, i := range nums {
		fmt.Println(i)
	}

	//slice
	println("-----------------slice--------------------------")
	mySlice := nums[1:3]
	fmt.Println(len(mySlice))
	fmt.Println(mySlice[0])
	println("-----------------map--------------------------")
	kvs := map[string]string{"a": "apple", "b": "banana", "c": "banana"}
	for key, value := range kvs {

		fmt.Printf("%s -> %s\n", key, value)

	}
	fmt.Println()
	var employee = make(map[string]int)
	employee["Mark"] = 10
	employee["Sandy"] = 20
	fmt.Println(employee)
	//delete(kvs, "b")

}
