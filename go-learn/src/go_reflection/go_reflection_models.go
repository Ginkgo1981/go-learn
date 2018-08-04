package go_reflection
//https://golangbot.com/reflection/
import (
	"fmt"
	"reflect"
)

type Order struct {
	id int
	customerId int
}


type Employee struct {
	id int
	name string
	address string
	salary int
	country string
}


func ReflectModel()  {

	order := Order{
		id: 1,
		customerId: 1,
	}

	createQuery(order)


	employee := Employee{
		id: 1,
		name: "chenjian",
		address: "address-xxx",
		salary: 10000,
		country:"ZH",
	}
	createQuery(employee)
}
func createQuery(q interface{}) {

	if reflect.ValueOf(q).Kind() == reflect.Struct {

		t := reflect.TypeOf(q).Name()

		query := fmt.Sprintf("insert into %s values(",t)

		v := reflect.ValueOf(q)

		for i := 0; i < v.NumField(); i++ {
			switch v.Field(i).Kind() {
			case reflect.Int:
				if i == 0 {
					query = fmt.Sprintf("%s%d", query, v.Field(i).Int())

				}else {
					query = fmt.Sprintf("%s, %d", query, v.Field(i).Int())

				}

			case reflect.String:
				if i == 0 {
					query = fmt.Sprintf("%s\"%s\"", query, v.Field(i).String())
				} else {
					query = fmt.Sprintf("%s, \"%s\"", query, v.Field(i).String())
				}
			default:
				fmt.Println("Unsupported type")
				return
			}
		}
		query = fmt.Sprintf("%s)",  query)
		fmt.Println(query)
	}else {
		fmt.Println("No A Struct")
	}
}