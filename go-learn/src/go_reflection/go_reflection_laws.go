package go_reflection

import (
			"fmt"
	"os"
	"io"
)

//https://blog.golang.org/laws-of-reflection

func ReflectionLaws() (interface{}, error)  {


	var r io.Reader

	tty, err := os.OpenFile("/Users/chenjian/current/go-projects/src/github.com/ginkgo1981/go-learn/src/go_reflection/go_reflection_laws.go", os.O_RDWR, 0)

	if err != nil {
		return nil, err
	}

	fmt.Println(tty)

	r = tty

	fmt.Println(r)
	w := r.(io.Writer)
	fmt.Println(w)


	var empty interface{}

	empty = w
	println(empty)

	var x io.Reader
	x = empty.(io.Reader)
	fmt.Println(x)

	return nil, err


}

func typeof(v interface{}) string {
	return fmt.Sprintf("%T",v)
}


