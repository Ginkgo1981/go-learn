package go_reflection
//https://gist.github.com/drewolson/4771479
import (
	"reflect"
	"fmt"
)

type Person struct {
	FirstName string `tag_name:"your first_name"`
	LastName string `tag_name:"your last_name"`
	Age int `tag_name:"your age"`
}


func (f *Person) reflect() {

	val := reflect.ValueOf(f).Elem()

	for i := 0; i < val.NumField(); i++ {

		valueField := val.Field(i)
		typeField := val.Type().Field(i)

		tag := typeField.Tag

		fmt.Printf("Field Name: %s \t Field Value: %v \t Tag Value: %s\n", typeField.Name, valueField.Interface(), tag.Get("tag_name"))
	}

}

func ReflectNow() {
	f := &Person{
		FirstName: "Chen",
		LastName: "Jian",
		Age: 35,
	}

	f.reflect()
}