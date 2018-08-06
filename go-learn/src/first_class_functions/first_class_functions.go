package first_class_functions

import "fmt"
//https://golangbot.com/first-class-functions/

func appendStr() func(string) string{

	t := "Hello"
	c := func(b string) string{
		t = t + " " + b
		return t
	}
	return c
}

type student struct {
	firstName string
	lastName string
	grade string
	country string
}

func filter(s []student, f func(student) bool)[]student  {
	var r []student
	for _, v := range s {

		if f(v) == true {
			r = append(r, v)
		}
	}
	return r
}

func iMap(s []student, f func(student) student)[]student {
	var r []student
	for _, v := range s {
		r = append(r, f(v))
	}
	return r
}

func ClosuresDemo()  {
	a := appendStr()
	b := appendStr()
	fmt.Println(a("world"))
	fmt.Println(b("Everyone"))
	fmt.Println(a("Gopher"))
	fmt.Println(b("!"))
}

func HighorderFuncionDemo() {

	s1 := student{
		firstName: "Naveen",
		lastName:  "Ramanathan",
		grade:     "A",
		country:   "India",
	}
	s2 := student{
		firstName: "Samuel",
		lastName:  "Johnson",
		grade:     "B",
		country:   "USA",
	}

	s := []student{s1, s2}

	m := iMap(s, func(s student) student {
		s.firstName = "chen"
		return s
		})

	f := filter(m, func(s student) bool {
		if s.grade == "B"{
			return true
		}
		return false
	})
	fmt.Println(f)
}