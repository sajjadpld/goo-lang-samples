package main

import "fmt"

func main() {
	fmt.Println("hello world")

	var name = "sajjad"
	family := "PLD"

	fmt.Println(name, " ", family)

	const age = 24
	fmt.Println(age)

	fmt.Printf("Hi, my name is %v and my last name is %v, im %v years old \n", name, family, age)

	var fatherName string

	fatherName = "Mahmoud"

	fmt.Println(fatherName)

	fmt.Println(&fatherName)

	fmt.Printf("the type of fatherName is %T\n", fatherName)


	var gfName string
	fmt.Scan(&gfName)

	fmt.Printf("your gir friend name is %v", gfName)

}
