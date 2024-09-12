package main

import (
	"fmt"
	"math"
)

type Sms interface {
	sendSms() bool
}
type SmsInfo struct {
	mobile string
	msg    string
}

func (obj SmsInfo) sendSms() bool {
	println("send ", obj.msg, "To mobile ", obj.mobile)
	return true
}

type myData interface {
	int | string
}

// Generic
func sum[T int | string](num1, num2 T) T {
	return num1 + num2
}

func sum2[T myData](num1, num2 T) T {
	return num1 + num2
}

// ----------------------------------------------------------------------------------------------------
// ----------------------------------------------------------------------------------------------------
type geometry interface {
	area() float64
	perim() float64
}

type rect struct {
	width, height float64
}
type circle struct {
	radius float64
}

func (r rect) area() float64 {
	return r.width * r.height
}
func (r rect) perim() float64 {
	return 2*r.width + 2*r.height
}
func (c circle) area() float64 {
	return math.Pi * c.radius * c.radius
}
func (c circle) perim() float64 {
	return 2 * math.Pi * c.radius
}
func measure(geom geometry) {
	fmt.Println(geom)
	fmt.Println("perim : ", geom.perim())
	fmt.Println("Area : ", geom.area())
}

// ----------------------------------------------------------------------------------------------------
// ----------------------------------------------------------------------------------------------------
func main() {
	/*	var sms Sms
		sms = SmsInfo{mobile: "09188717152", msg: "Hello"}
		sms.sendSms()

		println(sum(4, 5))
		println(sum("s", "h"))

		println(sum2(4, 5))
		println(sum2("s", "h"))*/

	r := rect{3, 4}
	c := circle{radius: 5}
	measure(r)
	println("------------------Circle---------------------------")
	measure(c)
}
