package go_defers

import "fmt"

type person struct {
	firstName string
	lastName string
}

func (p person)fullName()  {
	fmt.Printf("%s %s", p.firstName, p.lastName)
}

func printA(a int)  {
	fmt.Println("value of a in deferred function", a)
}

func DeferDemo()  {
	p := person{
		firstName: "chen",
		lastName: "jian",
	}
	defer p.fullName()
	fmt.Println("welcome")

	a := 5
	defer printA(5)
	a =10
	fmt.Println("value of a before deferred function call", a)
}

