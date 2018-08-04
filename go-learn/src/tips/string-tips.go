package tips

import (
	"fmt"
	"bytes"
	"strings"
	"math/rand"
	"time"
)

func TipNow()  {
	//https://www.calhoun.io/6-tips-for-using-strings-in-go/
	MutilineString()
	ConcatString()
	JoinString()
	Conv2String()
}

func Conv2String() {
	i := 123
	//t := strconv.Itoa(i)
	t := fmt.Sprintf("%d", i)
	fmt.Println(t)
}

func ConcatString()  {

	var b bytes.Buffer

	for i := 0; i < 1000; i++ {
		b.WriteString(RandString(10))
	}
	fmt.Println(b.String())
}

func JoinString(){
	var strs []string
	for i := 0; i < 1000; i++ {
		strs = append(strs, RandString(10))
	}
	fmt.Println(strings.Join(strs, ""))
}

var source = rand.NewSource(time.Now().UnixNano())
const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

func RandString(lengh int) string {
	b := make([]byte, lengh)
	for i := range b {
		b[i] = charset[source.Int63()%int64(len(charset))]
	}
	return string(b)
}

func MutilineString()  {
	str := `this is a
			mutiline
			string
			`
	fmt.Println(str)
}