package main

import "fmt"

func main() {
	nums := []int{1, 2, 3, 4, 5}
	println(sum(5, 6))
	println(multiResult(5))
	println(sums(nums...))

	//anonymous
	sum := func(a, b int) int {
		return a + b
	}
	println(sum(4, 8))

	//closure
	func() {
		println("closure function...")
	}()
	fmt.Println("------------closure----------------------------")
	pos, neg := adder(), adder()
	for i := 0; i < 10; i++ {
		fmt.Println(
			pos(i),
			neg(-2*i),
		)

	}
}

// normal function
func sum(a, b int) int {
	return a + b
}

// multiple result
func multiResult(num int) (a, b int) {
	a = num * 2
	b = num * 4
	return
}

// variadic function
func sums(nums ...int) int {
	total := 0
	for _, num := range nums {
		total += num
	}
	return total
}

// For example, the adder function returns a closure. Each closure is bound to its own sum variable
func adder() func(int) int {
	sum := 0
	return func(x int) int {
		sum += x
		return sum
	}
}
