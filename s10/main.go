package main

type User struct {
	name   string
	family string
	age    int
}

func main() {
	var obj User
	obj.name = "Shahed"
	obj.family = "Shariati"
	obj.age = 37
	println(obj.name)

	user := User{name: "Samrand", family: "Shariati", age: 5}
	println(user.name)

}
