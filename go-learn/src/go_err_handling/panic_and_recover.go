package go_err_handling
//https://golangbot.com/panic-and-recover/
import "fmt"



func fullName(firstName *string, lastname *string) {
	defer fmt.Println("deferred call in fullName")

	defer recoverName()

	if firstName == nil {
		panic("runtime error: first name cannot be nil")
	}

	if lastname == nil {
		panic("runtime error: last name cannot be nil")
	}
	fmt.Printf("%s %s", *firstName, *lastname)
	fmt.Println("returned normally from fullName")
}
func recoverName() {
	if r := recover(); r != nil {
		fmt.Println("recovered from ", r)
	}
}

func PanicRecoverDemo()  {
	defer fmt.Println("deffer call PanicRecoverDemo")
	firstName := "Elon"
	fullName(&firstName, nil)
	fmt.Println("returned normally from main")
}