package main

import "fmt"

type User struct {
	name string
	age  int
}

func main() {
	// a := User{} //stack 할당

	a := new(User) //heap 할당

	a.name = "jake"
	a.age = 27

	/*

	 */
	fmt.Printf("%v\n", a)  //output : {name, 10}
	fmt.Printf("%#v\n", a) //output : main.User{Name:"name", Age:10}
	fmt.Printf("%+v\n", a) //output : {Name:name Age:10}
	fmt.Printf("%T\n", a)  //output : main.User
	fmt.Println(a)         //output : &{jake 27}
	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}
	fmt.Println("hello world")
}
