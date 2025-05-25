package main

import "fmt"

type user struct {
	Name string
	Age  int
	Mail string
}

func User() {

	user1 := user{
		Name: "Alice",
		Age:  24,
		Mail: "Alice92@gmail.com",
	}

	user2 := user{
		Name: "Boris",
		Age:  31,
		Mail: "BorisHe@web.de",
	}

	fmt.Printf("User Data Print: %v , %v , %v \n", user1.Name, user1.Age, user1.Mail)
	fmt.Printf("User Data Print: %v , %v , %v \n", user2.Name, user2.Age, user2.Mail)
	PrintUser(user1)
	PrintUser(user2)

}

func PrintUser(user user) {
	fmt.Printf("Name: %v, Age: %v, Mail: %v\n", user.Name, user.Age, user.Mail)
}

func main() {

	User()

}
