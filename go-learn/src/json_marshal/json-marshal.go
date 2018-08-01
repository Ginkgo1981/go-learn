package json_marshal

import (
	"encoding/json"
	"fmt"
)

type response1 struct {
	Page int
	Fruits []string
}

type response2 struct {
	Page   int      `json:"page"`
	Fruits []string `json:"fruits"`
}

type App struct {
	Id string
	Title string
}

func MarshalNow() {
	var app App
	data := []byte(`
		{
			"id": "1234",
			"title": "Welcome to App"
		}
`)
	err := json.Unmarshal(data, &app)
	if err != nil {
		println(err)
	}
	fmt.Println(app)
}

