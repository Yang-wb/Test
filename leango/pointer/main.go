package main

import "fmt"

type Student struct {
	name    string
	age     int
	married bool
	sex     int8
}

func main() {
	s1 := Student{"Steven", 35, true, 1}
	s2 := Student{"Sunny", 20, false, 0}

	var a *Student = &s1
	var b *Student = &s2

	fmt.Println("\n---------------------")
	fmt.Printf("s1类型为%T， 值为%v \n", s1, s1)
	fmt.Printf("s2类型为%T， 值为%v \n", s2, s2)

	fmt.Println("\n---------------------")
	fmt.Printf("a类型为%T， 值为%v \n", a, a)
	fmt.Printf("b类型为%T， 值为%v \n", b, b)

	fmt.Println("\n---------------------")
	fmt.Printf("*a类型为%T， 值为%v \n", *a, *a)
	fmt.Printf("*b类型为%T， 值为%v \n", *b, *b)

	fmt.Println("\n---------------------")
	fmt.Println(s1.name, s1.age, s1.married, s1.sex)
	fmt.Println(a.name, a.age, a.married, a.sex)

	fmt.Println("\n---------------------")
	fmt.Println((*a).name, (*a).age, (*a).married, (*a).sex)
	fmt.Println(&a.name, &a.age, &a.married, &a.sex)
}
